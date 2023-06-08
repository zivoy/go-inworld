package grpc

import (
	"github.com/google/uuid"
	"github.com/zivoy/go-inworld/internal/protoBuf/packets"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Generator struct {
	TargetId string
}

func (g *Generator) NewTextPacket(text string) *packets.InworldPacket {
	packet := g.NewBasePacket(false, false)
	packet.Packet = &packets.InworldPacket_Text{
		Text: &packets.TextEvent{
			Text:       text,
			SourceType: packets.TextEvent_TYPED_IN,
			Final:      true,
		},
	}

	return packet
}

func (g *Generator) NewDataChunkPacket(data []byte, dataType packets.DataChunk_DataType) *packets.InworldPacket {
	packet := g.NewBasePacket(false, false)
	packet.Packet = &packets.InworldPacket_DataChunk{
		DataChunk: &packets.DataChunk{
			Data: &packets.DataChunk_Chunk{Chunk: data},
			Type: dataType,
		},
	}

	return packet
}

//todo other packet types

func (g *Generator) NewBasePacket(utteranceId, interactionId bool) *packets.InworldPacket {
	packetId := &packets.PacketId{
		PacketId: uuid.NewString(),
	}
	if utteranceId {
		packetId.UtteranceId = uuid.NewString()
	}
	if interactionId {
		packetId.InteractionId = uuid.NewString()
	}

	route := &packets.Routing{
		Source: &packets.Actor{
			Type: packets.Actor_PLAYER,
		},
		Target: &packets.Actor{
			Type: packets.Actor_AGENT,
			Name: g.TargetId,
		},
	}

	return &packets.InworldPacket{
		Timestamp: timestamppb.Now(),
		Routing:   route,
		PacketId:  packetId,
	}
}
