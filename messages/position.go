package messages

import "engo.io/engo"

const GetPlayerPositionType string = "GetPlayerPositionMessage"

type GetPlayerPosition struct {
	Position engo.Point
}

func (GetPlayerPosition) Type() string {
	return GetPlayerPositionType
}

const SendPlayerPositionType string = "SendPlayerPositionMessage"

type SendPlayerPosition struct{}

func (SendPlayerPosition) Type() string {
	return SendPlayerPositionType
}
