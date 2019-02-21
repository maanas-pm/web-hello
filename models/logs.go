package models

import(
	"time"
)

type Log struct {
    Id        int64	`json:"id"`
    Time      time.Time	`json:"time"`
    Request   string	`json:"request"`
    Response  int	`json:"reswponse"`
    Log_level string	`json:"log_level"`
}
