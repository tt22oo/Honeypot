package session

import "fmt"

func (s *Session) WriteAndRead(data string) (string, error) {
	// write
	_, err := s.Conn.Write([]byte(data))
	if err != nil {
		return "", fmt.Errorf("Write Error: %s", err.Error())
	}

	// read
	var result string
	buf := make([]byte, 1)
	for {
		_, err := s.Conn.Read(buf)
		if err != nil {
			return "", fmt.Errorf("Read Error: %s", err.Error())
		}

		if buf[0] == '\n' || buf[0] == '\r' {
			break
		}

		result += string(buf[0])
	}

	return result, nil
}
