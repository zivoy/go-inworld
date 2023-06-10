package internal

import (
	"github.com/google/uuid"
	goinworld "github.com/zivoy/go-inworld"
	"github.com/zivoy/go-inworld/internal/protoBuf/engine"
	"github.com/zivoy/go-inworld/internal/protoBuf/packets"
	"github.com/zivoy/go-inworld/internal/protoBuf/studioTokens"
	"github.com/zivoy/go-inworld/session"
	"time"
)

func GetTypeFromProto(packet *packets.InworldPacket) goinworld.InworldPacketType {
	p := packet.GetPacket()
	switch p.(type) {
	case *packets.InworldPacket_Text:
		return goinworld.InworldPacketText
	case *packets.InworldPacket_Custom:
		return goinworld.InworldPacketTrigger
	case *packets.InworldPacket_Control:
		return goinworld.InworldPacketControl
	case *packets.InworldPacket_Emotion:
		return goinworld.InworldPacketEmotion
	case *packets.InworldPacket_Mutation:
		m := p.(*packets.InworldPacket_Mutation).Mutation
		if _, ok := m.Mutation.(*packets.MutationEvent_CancelResponses); ok {
			return goinworld.InworldPacketCancelResponse
		}
	case *packets.InworldPacket_DataChunk:
		d := p.(*packets.InworldPacket_DataChunk).DataChunk
		if d.GetType() == packets.DataChunk_AUDIO {
			return goinworld.InworldPacketAudio
		} else if d.GetType() == packets.DataChunk_SILENCE {
			return goinworld.InworldPacketSilence
		}
	}
	return goinworld.InworldPacketUnknown
}

func GetMessageFromProto(packet *packets.InworldPacket) *goinworld.InworldPacket {
	p := &goinworld.InworldPacket{
		PacketId: PacketIdFromProto(packet.GetPacketId()),
		Routing:  RoutingFromProto(packet.GetRouting()),
		Date:     packet.GetTimestamp().AsTime(),

		PacketType: GetTypeFromProto(packet),
	}

	switch p.PacketType {
	case goinworld.InworldPacketTrigger:
		p.Trigger = GetTriggerEvent(packet.GetCustom())
	case goinworld.InworldPacketText:
		p.Text = GetTextEvent(packet.GetText())
	case goinworld.InworldPacketAudio:
		p.Audio = GetAudioEvent(packet.GetDataChunk())
	case goinworld.InworldPacketControl:
		p.Control = GetControlEvent(packet.GetControl())
	case goinworld.InworldPacketSilence:
		p.Silence = GetSilenceEvent(packet.GetDataChunk())
	case goinworld.InworldPacketEmotion:
		p.Emotions = GetEmotionEvent(packet.GetEmotion())
	case goinworld.InworldPacketCancelResponse:
		p.CancelResponses = GetCancelEvent(packet.GetCancelresponses())
	}

	return p
}

func PacketIdFromProto(id *packets.PacketId) *goinworld.PacketId {
	return &goinworld.PacketId{
		PacketId:      uuid.MustParse(id.PacketId),
		UtteranceId:   uuid.MustParse(id.UtteranceId),
		InteractionId: uuid.MustParse(id.InteractionId),
	}
}

func RoutingFromProto(route *packets.Routing) *goinworld.Routing {
	return &goinworld.Routing{
		Target: ActorFromProto(route.GetTarget()),
		Source: ActorFromProto(route.GetSource()),
	}
}

func ActorFromProto(actor *packets.Actor) *goinworld.Actor {
	t := actor.GetType()
	return &goinworld.Actor{
		Name:        actor.GetName(),
		IsPlayer:    t == packets.Actor_PLAYER,
		IsCharacter: t == packets.Actor_AGENT,
	}
}

func GetTriggerEvent(event *packets.CustomEvent) *goinworld.TriggerEvent {
	params := event.GetParametersList()
	e := &goinworld.TriggerEvent{
		Name:       event.GetName(),
		Parameters: make([]*goinworld.TriggerParameter, len(params)),
	}

	for i, param := range params {
		e.Parameters[i] = &goinworld.TriggerParameter{
			Name:  param.GetName(),
			Value: param.GetValue(),
		}
	}

	return e
}

func GetTextEvent(event *packets.TextEvent) *goinworld.TextEvent {
	return &goinworld.TextEvent{
		Text:  event.GetText(),
		Final: event.GetFinal(),
	}
}

func GetAudioEvent(event *packets.DataChunk) *goinworld.AudioEvent {
	p := event.GetAdditionalPhonemeInfoList()
	e := &goinworld.AudioEvent{
		Chunk:                 event.GetChunk(),
		AdditionalPhonemeInfo: make([]*goinworld.AdditionalPhonemeInfo, len(p)),
	}

	for i, info := range p {
		e.AdditionalPhonemeInfo[i] = &goinworld.AdditionalPhonemeInfo{
			Phoneme:     info.GetPhoneme(),
			StartOffset: info.GetStartOffset().AsDuration(),
		}
	}

	return e
}

func GetControlEvent(event *packets.ControlEvent) *goinworld.ControlEvent {
	e := goinworld.InworldControlUnknown
	// maybe switch to switch if there are more events
	if event.GetAction() == packets.ControlEvent_INTERACTION_END {
		e = goinworld.InworldControlInteractionEnd
	}

	return &goinworld.ControlEvent{
		Type: e,
	}
}

func GetSilenceEvent(event *packets.DataChunk) *goinworld.SilenceEvent {
	return &goinworld.SilenceEvent{
		Duration: time.Duration(event.GetDurationMs()) * time.Millisecond,
	}
}

func GetEmotionEvent(event *packets.EmotionEvent) *goinworld.EmotionEvent {
	return &goinworld.EmotionEvent{
		Behavior: goinworld.EmotionalBehavior(event.GetBehavior()),
		Strength: goinworld.EmotionalStrength(event.GetStrength()),
	}
}

func GetCancelEvent(event *packets.CancelResponsesEvent) *goinworld.CancelResponsesEvent {
	utterances := event.GetUtteranceIdList()
	e := &goinworld.CancelResponsesEvent{
		InteractionId: uuid.MustParse(event.GetInteractionId()), // maybe make a check for if it's empty?
		UtteranceId:   make([]uuid.UUID, len(utterances)),
	}

	for i, utterance := range utterances {
		e.UtteranceId[i] = uuid.MustParse(utterance)
	}

	return e
}

///////// session

func GetCharacterFromProto(response *engine.LoadSceneResponse_Agent) *session.Character {
	assets := response.GetCharacterAssets()

	return &session.Character{
		Id:           response.GetAgentId(),
		ResourceName: response.GetBrainName(),
		DisplayName:  response.GetGivenNAme(),
		Assets: &session.Assets{
			AvatarImage:         assets.GetAvatarImg(),
			AvatarImageOriginal: assets.GetAvatarImgOriginal(),
			RpmModelUri:         assets.GetRpmModelUri(),
			RpmImageUriPortrait: assets.GetRpmImageUriPortrait(),
			RpmImageUriPosture:  assets.GetRpmImageUriPosture(),
		},
	}
}

func GetSceneFromProto(response *engine.LoadSceneResponse) *session.Scene {
	characters := response.GetAgentsList()
	chars := make([]*session.Character, len(characters))
	for i, character := range characters {
		chars[i] = GetCharacterFromProto(character)
	}

	return &session.Scene{
		Key:        response.GetKey(),
		Characters: chars,
	}
}

func GetTokenFromProto(proto *studioTokens.SessionAccessToken) *session.Token {
	return &session.Token{
		Token:          proto.GetToken(),
		Type:           proto.GetType(),
		ExpirationTime: proto.GetExpirationTime().AsTime(),
		SessionId:      proto.GetSessionId(),
	}
}
