package global

import (
	"log"
	"strconv"
)

func ParseArgs(args []string) {
	for i, v := range args {
		switch v {
		case "addr":
			IP = args[i+1]
			break
		case "port":
			Port = ":" + args[i+1]
			break
		case "machine":
			id, err := strconv.Atoi(args[i+1])
			if err != nil {
				log.Panic("Invalid machine ID")
			}
			MachineID = int64(id)
		}
	}

	configure()
}
