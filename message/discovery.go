package message

type StartScanning struct {
	message
}

type StopScanning struct {
	message
}

type ScanningFinished struct {
	message
}

type RequestDeviceList struct {
	message
}
