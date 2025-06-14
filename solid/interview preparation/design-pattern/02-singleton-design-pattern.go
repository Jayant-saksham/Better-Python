package design_pattern

import (
	"fmt"
	"sync"
)

type Log struct {
	logs []string
}

var (
	logObject Log
	once      sync.Once
)

func GetLogger() *Log {
	once.Do(func() {
		fmt.Println("Creating Logger Instance...")
		logObject = Log{}
	})
	return &logObject
}

func (l *Log) Write(message string) {
	l.logs = append(l.logs, message)
	fmt.Println("Logged:", message)
}

func main() {
	logger1 := GetLogger()
	logger1.Write("First log message")

	logger2 := GetLogger()
	logger2.Write("Second log message")

	fmt.Println("Same logger?", logger1 == logger2)
	fmt.Println(logger1.logs)
}
