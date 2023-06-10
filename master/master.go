package master

import (
	conxt "bluetooth/const"
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

type engine struct {
	c               chan byte
	isMouse         bool
	stopMouseAction chan struct{}
}

func NewEngine() *engine {
	return &engine{
		c:               make(chan byte, 1000),
		stopMouseAction: make(chan struct{}),
	}
}

func (e *engine) Serve() {
	for true {
		select {
		case btn := <-e.c:
			e.handleFunc()(btn)
		}
	}
}
func (e *engine) handleFunc() func(btn byte) {
	if e.isMouse {
		return e.handleMouse
	} else {
		return e.handleKeyboard
	}
}

func (e *engine) handleKeyboard(btn byte) {
	if btn == conxt.TurnMouse && e.isMouse == false {
		e.isMouse = true
		e.stopMouseAction = make(chan struct{})
		fmt.Println("turn to mouse")
		return
	}
	for i := 0; i < 8; i++ {
		key := btn & (1 << i)
		s, ok := conxt.KeyBoardMap[key]
		if ok {
			robotgo.KeyTap(s)
		}
	}
}

func (e *engine) handleMouse(btn byte) {
	if btn == conxt.TurnKeyboard && e.isMouse == true {
		e.isMouse = false
		if !isChanClose(e.stopMouseAction) {
			close(e.stopMouseAction)
		}
		return
	}
	// 归位
	if btn == conxt.BtnNone {
		fmt.Println("close 0")
		close(e.stopMouseAction)
		return
	}
	if isChanClose(e.stopMouseAction) {
		e.stopMouseAction = make(chan struct{})
	}
	switch conxt.KeyMouseMap[btn] {
	case "left":
		fmt.Println("in left")
		go func() {
			for true {
				select {
				case <-e.stopMouseAction:
					fmt.Println(" return")
					return

				default:
					x, y := robotgo.GetMousePos()
					fmt.Println("左移")
					robotgo.Move(x-conxt.MoveAtom, y)
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()
	case "right":
		go func() {
			for true {
				select {
				case <-e.stopMouseAction:
					return

				default:
					x, y := robotgo.GetMousePos()
					robotgo.Move(x+conxt.MoveAtom, y)
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()
	case "up":
		go func() {
			for true {
				select {
				case <-e.stopMouseAction:
					return

				default:
					x, y := robotgo.GetMousePos()
					robotgo.Move(x, y-conxt.MoveAtom)
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()
	case "down":
		go func() {
			for true {
				select {
				case <-e.stopMouseAction:
					return

				default:
					x, y := robotgo.GetMousePos()
					robotgo.Move(x, y+conxt.MoveAtom)
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()
	case "click":
		robotgo.Click()
	case "clickRight":
		robotgo.Click("right")
	case "scrollUp":
		go func() {
			for true {
				select {
				case <-e.stopMouseAction:
					return

				default:
					robotgo.ScrollMouse(conxt.ScrollAtom, "up")
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()
	case "scrollDown":
		go func() {
			for true {
				select {
				case <-e.stopMouseAction:
					return

				default:
					robotgo.ScrollMouse(conxt.ScrollAtom, "down")
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()
	default:
		fmt.Println("按键冲突")
	}

}

func (e *engine) Push(buf []byte) {
	for i := 0; i < len(buf); i++ {
		e.c <- buf[i]
	}
}

func isChanClose(ch chan struct{}) bool {
	select {
	case _, received := <-ch:
		return !received
	default:
	}
	return false
}
