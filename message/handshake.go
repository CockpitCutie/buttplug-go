package message

type RequestServerInfo struct {
	message
	ClientName     string
	MessageVersion uint
}

type ServerInfo struct {
	message
	ServerName     string
	MessageVersion uint
	MaxPingTime    uint
}
