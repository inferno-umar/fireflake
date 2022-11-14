package global

// Machine ID
var MachineID int64

// A sequence
var sequence int64

var IP, Port string

func configure() {
	if IP == "" {
		IP = "127.0.0.1"
	}

	if Port == "" {
		Port = ":80"
	}
}
