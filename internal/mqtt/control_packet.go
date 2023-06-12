package mqtt

type FixedHeader struct {
	controlPacketType byte
}

type ControlPacket struct {
	Header         string
	VariableHeader string
	Payload        string
}
