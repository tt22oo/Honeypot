package logger

import "log"

const (
	TypeError string = "\033[31m[ERROR]\033[0m"
)

// print error log
func Error(data string) {
	log.Printf("%s %s\r\n", TypeError, data)
}
