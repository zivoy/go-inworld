package goinworld

import "github.com/zivoy/go-inworld/internal/protoBuf/packets"

func getType(packet *packets.InworldPacket) InworldPacketType {
	switch packet.Packet.(type) {
	case *packets.InworldPacket_Text:
		return InworldPacketText
	case *packets.InworldPacket_Custom:
		return InworldPacketTrigger
	case *packets.InworldPacket_Control:
		return InworldPacketControl
	case *packets.InworldPacket_Emotion:
		return InworldPacketEmotion
	case *packets.InworldPacket_Mutation:
		m := packet.Packet.(*packets.InworldPacket_Mutation).Mutation
		if _, ok := m.Mutation.(*packets.MutationEvent_CancelResponses); ok {
			return InworldPacketCancelResponse
		}
	case *packets.InworldPacket_DataChunk:
		d := packet.Packet.(*packets.InworldPacket_DataChunk).DataChunk
		if d.GetType() == packets.DataChunk_AUDIO {
			return InworldPacketAudio
		} else if d.GetType() == packets.DataChunk_SILENCE {
			return InworldPacketSilence
		}
	}
	return InworldPacketUnknown
}
