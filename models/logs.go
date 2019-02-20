package model

import(
	"fmt"
	"time"
)
type Log_level int

const (
	TRACE = Log_level(iota)
	DEBUG
	INFO
	WARN
	ERROR
)

func SetLogLevel(level Log_level) {
    switch level {
    case TRACE:
        fmt.Println("trace")
        return

    case INFO:
        fmt.Println("info")
        return

    case DEBUG:
	fmt.Println("debug")
        return

    case WARN:
        fmt.Println("warning")
        return
    
    case ERROR:
        fmt.Println("error")
        return

    default:
        fmt.Println("default")
        return

    }
}

type Log struct {
    Id        int64	`json:"id"`
    Time      time.Time	`json:"time"`
    Request   string	`json:"request"`
    Response  int	`json:"reswponse"`
    Log_level Log_level	`json:"log_level"`
}
