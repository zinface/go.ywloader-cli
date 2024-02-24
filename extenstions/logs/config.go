package logs

import (
	"fmt"
	"log"
)

type Logs struct {
	Prefix string
}

func (l *Logs) Println(message string) {
	s := fmt.Sprintf("%v: %v", l.Prefix, message)
	log.Println(s)
}

func (l *Logs) Print(prefix, message interface{}) {
	s := fmt.Sprintf("%v: %v", prefix, message)
	l.Println(s)
}

func (l *Logs) Printf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.Println(message)
}
