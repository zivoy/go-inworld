package goinworld

import (
	"github.com/google/uuid"
	"github.com/zivoy/go-inworld/internal/protoBuf/packets"
	"time"
)

func (i *InworldPacket) IsText() bool           { return i.packetType == InworldPacketText }
func (i *InworldPacket) IsAudio() bool          { return i.packetType == InworldPacketAudio }
func (i *InworldPacket) IsControl() bool        { return i.packetType == InworldPacketControl }
func (i *InworldPacket) IsTrigger() bool        { return i.packetType == InworldPacketTrigger }
func (i *InworldPacket) IsEmotion() bool        { return i.packetType == InworldPacketEmotion }
func (i *InworldPacket) IsSilence() bool        { return i.packetType == InworldPacketSilence }
func (i *InworldPacket) IsCancelResponse() bool { return i.packetType == InworldPacketCancelResponse }
func (i *InworldPacket) IsInteractionEnd() bool {
	return i.IsControl() && i.Control.Type == InworldControlInteractionEnd
}

func getTypeFromProto(packet *packets.InworldPacket) InworldPacketType {
	p := packet.GetPacket()
	switch p.(type) {
	case *packets.InworldPacket_Text:
		return InworldPacketText
	case *packets.InworldPacket_Custom:
		return InworldPacketTrigger
	case *packets.InworldPacket_Control:
		return InworldPacketControl
	case *packets.InworldPacket_Emotion:
		return InworldPacketEmotion
	case *packets.InworldPacket_Mutation:
		m := p.(*packets.InworldPacket_Mutation).Mutation
		if _, ok := m.Mutation.(*packets.MutationEvent_CancelResponses); ok {
			return InworldPacketCancelResponse
		}
	case *packets.InworldPacket_DataChunk:
		d := p.(*packets.InworldPacket_DataChunk).DataChunk
		if d.GetType() == packets.DataChunk_AUDIO {
			return InworldPacketAudio
		} else if d.GetType() == packets.DataChunk_SILENCE {
			return InworldPacketSilence
		}
	}
	return InworldPacketUnknown
}

func getMessageFromProto(packet *packets.InworldPacket) *InworldPacket {
	p := &InworldPacket{
		PacketId: packetIdFromProto(packet.GetPacketId()),
		Routing:  routingFromProto(packet.GetRouting()),
		Date:     packet.GetTimestamp().AsTime(),

		packetType: getTypeFromProto(packet),
	}

	switch p.packetType {
	case InworldPacketTrigger:
		p.Trigger = getTriggerEvent(packet.GetCustom())
	case InworldPacketText:
		p.Text = getTextEvent(packet.GetText())
	case InworldPacketAudio:
		p.Audio = getAudioEvent(packet.GetDataChunk())
	case InworldPacketControl:
		p.Control = getControlEvent(packet.GetControl())
	case InworldPacketSilence:
		p.Silence = getSilenceEvent(packet.GetDataChunk())
	case InworldPacketEmotion:
		p.Emotions = getEmotionEvent(packet.GetEmotion())
	case InworldPacketCancelResponse:
		p.CancelResponses = getCancelEvent(packet.GetCancelresponses())
	}

	return p
}

func packetIdFromProto(id *packets.PacketId) *PacketId {
	return &PacketId{
		PacketId:      uuid.MustParse(id.PacketId),
		UtteranceId:   uuid.MustParse(id.UtteranceId),
		InteractionId: uuid.MustParse(id.InteractionId),
	}
}

func routingFromProto(route *packets.Routing) *Routing {
	return &Routing{
		Target: actorFromProto(route.GetTarget()),
		Source: actorFromProto(route.GetSource()),
	}
}

func actorFromProto(actor *packets.Actor) *Actor {
	t := actor.GetType()
	return &Actor{
		Name:        actor.GetName(),
		IsPlayer:    t == packets.Actor_PLAYER,
		IsCharacter: t == packets.Actor_AGENT,
	}
}

func getTriggerEvent(event *packets.CustomEvent) *TriggerEvent {
	params := event.GetParametersList()
	e := &TriggerEvent{
		Name:       event.GetName(),
		Parameters: make([]*TriggerParameter, len(params)),
	}

	for i, param := range params {
		e.Parameters[i] = &TriggerParameter{
			Name:  param.GetName(),
			Value: param.GetValue(),
		}
	}

	return e
}

func getTextEvent(event *packets.TextEvent) *TextEvent {
	return &TextEvent{
		Text:  event.GetText(),
		Final: event.GetFinal(),
	}
}

func getAudioEvent(event *packets.DataChunk) *AudioEvent {
	p := event.GetAdditionalPhonemeInfoList()
	e := &AudioEvent{
		Chunk:                 event.GetChunk(),
		AdditionalPhonemeInfo: make([]*AdditionalPhonemeInfo, len(p)),
	}

	for i, info := range p {
		e.AdditionalPhonemeInfo[i] = &AdditionalPhonemeInfo{
			Phoneme:     info.GetPhoneme(),
			StartOffset: info.GetStartOffset().AsDuration(),
		}
	}

	return e
}

func getControlEvent(event *packets.ControlEvent) *ControlEvent {
	e := InworldControlUnknown
	// maybe switch to switch if there are more events
	if event.GetAction() == packets.ControlEvent_INTERACTION_END {
		e = InworldControlInteractionEnd
	}

	return &ControlEvent{
		Type: e,
	}
}

func getSilenceEvent(event *packets.DataChunk) *SilenceEvent {
	return &SilenceEvent{
		Duration: time.Duration(event.GetDurationMs()) * time.Millisecond,
	}
}

func getEmotionEvent(event *packets.EmotionEvent) *EmotionEvent {
	return &EmotionEvent{
		Behavior: EmotionalBehavior(event.GetBehavior()),
		Strength: EmotionalStrength(event.GetStrength()),
	}
}

func getCancelEvent(event *packets.CancelResponsesEvent) *CancelResponsesEvent {
	utterances := event.GetUtteranceIdList()
	e := &CancelResponsesEvent{
		InteractionId: uuid.MustParse(event.GetInteractionId()), // maybe make a check for if it's empty?
		UtteranceId:   make([]uuid.UUID, len(utterances)),
	}

	for i, utterance := range utterances {
		e.UtteranceId[i] = uuid.MustParse(utterance)
	}

	return e
}
