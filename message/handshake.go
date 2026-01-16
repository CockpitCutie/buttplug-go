package message

type RequestServerInfo struct {
	message
	ClientName           string
	ProtocolVersionMajor uint
	ProtocolVersionMinor uint
}

type ServerInfo struct {
	message
	ServerName           string
	MaxPingTime          uint
	ProtocolVersionMajor uint
	ProtocolVersionMinor uint
}
