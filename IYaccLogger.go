package yaccidl

import "log"

type IYaccLogger interface {
	LogWithLevel(level int, cb func(logger *log.Logger))
	Log(info string, params ...interface{})
}
