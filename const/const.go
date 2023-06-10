package conxt

const MoveAtom = 1
const ScrollAtom = 1

const BtnNone = 0
const (
	Btn1 = 1 << iota
	Btn2
	Btn3
	Btn4
	Btn5
	Btn6
	Btn7
	Btn8
)

const (
	TurnMouse = iota + 0xfe
	TurnKeyboard
)

var KeyBoardMap = map[byte]string{
	BtnNone: "",
	Btn1:    "t",
	Btn2:    "y",
	Btn3:    "u",
	Btn4:    "i",
	Btn5:    "o",
	Btn6:    "p",
	Btn7:    "a",
	Btn8:    "s",
}

var KeyMouseMap = map[byte]string{
	BtnNone: "none",
	Btn1:    "left",
	Btn2:    "right",
	Btn3:    "up",
	Btn4:    "down",
	Btn5:    "click",
	Btn6:    "clickRight",
	Btn7:    "scrollUp",
	Btn8:    "scrollDown",
}
