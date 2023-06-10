package test

import (
	conxt "bluetooth/const"
	"github.com/go-vgo/robotgo"
	"testing"
)

func TestHandler(t *testing.T) {
	x, y := robotgo.GetMousePos()
	robotgo.MoveSmooth(x-conxt.MoveAtom, y)
}
