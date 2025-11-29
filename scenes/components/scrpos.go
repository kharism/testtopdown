package components

import "github.com/yohamta/donburi"

type LogicalPositionData struct {
	X float64
	Y float64
}

var LogicalPos = donburi.NewComponentType[LogicalPositionData]()
