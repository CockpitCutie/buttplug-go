package message

type RequestServerInfo struct {
	message
	ClientName           string
	ProtocolVersionMajor int
	ProtocolVersionMinor int
}

type ServerInfo struct {
	message
	ServerName           string
	MaxPingTime          int
	ProtocolVersionMajor int
	ProtocolVersionMinor int
}
