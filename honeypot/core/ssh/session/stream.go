package session

import (
	"fmt"

	"github.com/gliderlabs/ssh"
)

func WriteAndRead(s ssh.Session, data string) (string, error) {
	// write
	_, err := s.Write([]byte(data))
	if err != nil {
		return "", fmt.Errorf("Write Error: %s", err.Error())
	}

	// read
	var result string
	buf := make([]byte, 1)
	for {
		_, err := s.Read(buf)
		if err != nil {
			return "", fmt.Errorf("Read Error: %s", err.Error())
		}

		_, err = s.Write(buf)
		if err != nil {
			return "", fmt.Errorf("Write Error: %s", err.Error())
		}

		if buf[0] == '\n' || buf[0] == '\r' {
			_, err = s.Write([]byte("\r\n"))
			if err != nil {
				return "", fmt.Errorf("Write Error: %s", err.Error())
			}

			break
		}

		result += string(buf[0])
	}

	return result, nil
}
