package goinworld

func (i *InworldPacket) IsText() bool           { return i.PacketType == InworldPacketText }
func (i *InworldPacket) IsAudio() bool          { return i.PacketType == InworldPacketAudio }
func (i *InworldPacket) IsControl() bool        { return i.PacketType == InworldPacketControl }
func (i *InworldPacket) IsTrigger() bool        { return i.PacketType == InworldPacketTrigger }
func (i *InworldPacket) IsEmotion() bool        { return i.PacketType == InworldPacketEmotion }
func (i *InworldPacket) IsSilence() bool        { return i.PacketType == InworldPacketSilence }
func (i *InworldPacket) IsCancelResponse() bool { return i.PacketType == InworldPacketCancelResponse }
func (i *InworldPacket) IsInteractionEnd() bool {
	return i.IsControl() && i.Control.Type == InworldControlInteractionEnd
}
