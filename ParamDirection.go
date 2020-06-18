package yaccidl

type ParamDirection uint8

const ParamDirectionUndefined ParamDirection = 0
const (
	ParamDirectionIn ParamDirection = 1 << iota
	ParamDirectionOut
)
