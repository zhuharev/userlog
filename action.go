package userlog

import (
	"encoding/binary"
)

type ActionType uint16

func NewActionType(in []byte) ActionType {
	return ActionType(binary.BigEndian.Uint16(in))
}

func (at ActionType) ToBytes() []byte {
	var b = make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(at))
	return b
}

const (
	ARegister ActionType = iota + 1
	ALogin
	AVisit
	AClick
	ASearch
)

type ObjectType uint16

func (ot ObjectType) ToBytes() []byte {
	var b = make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(ot))
	return b
}

const (
	OUser ObjectType = iota + 1
	OPost
)
