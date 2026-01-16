package message

type RequestServerInfo struct {
	message
	ClientName           string
	ProtocolVersionMajor uint
	ProtocolMinorVersion uint
}

type ServerInfo struct {
	message
	ServerName           string
	MaxPingTime          uint
	ProtocolVersionMajor uint
	ProtocolMinorVersion uint
}
