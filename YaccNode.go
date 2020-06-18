package yaccidl

import (
	"fmt"
	"strings"
)

type YaccNode struct {
	id            string
	idlIdentifier string
	lexemData     *LexemValue
	nextNode      IYaccNode
	array         []int64
}

func (y *YaccNode) AddArray(array []int64) {
	y.array = array
}

func (y *YaccNode) IsArray() bool {
	return len(y.array) > 0
}

func (y *YaccNode) Array() []int64 {
	return y.array
}

func (y YaccNode) GetIdlIdentifier() string {
	return y.idlIdentifier
}

func (y YaccNode) String() string {
	var ss []string
	addValue := func(s string) {
		if s == "" {
			s = "Identifier not assigned"
		}
		ss = append(ss, s)
	}

	var node IYaccNode
	node = &y
	addValue(node.Identifier())
	multiple := false
	for node.GetNextNode() != nil {
		node = node.GetNextNode()
		multiple = true
	}
	if multiple {
		addValue("..")
		addValue(node.Identifier())
	}
	return fmt.Sprintf("%v", strings.Join(ss, ","))
}

func (y YaccNode) LexemData() *LexemValue {
	return y.lexemData
}

func (y YaccNode) GetNextNode() IYaccNode {
	return y.nextNode
}

func (y YaccNode) Identifier() string {
	return y.id
}

func (y *YaccNode) SetNextNode(node IYaccNode) error {
	y.nextNode = node
	return nil
}

func (y *YaccNode) asIYaccNode() IYaccNode {
	return y
}

func NewYaccNode(id string, lexemData *LexemValue) *YaccNode {
	if len(id) >= 2 && id[0] == '@' {
		id = id[1:]
	}

	return &YaccNode{
		id:        id,
		lexemData: lexemData,
		nextNode:  nil,
	}
}
