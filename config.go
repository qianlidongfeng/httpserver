package httpserver

import (
	"time"
)

type Config struct{
	Address string
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration
}