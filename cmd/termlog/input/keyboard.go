package input

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

type Key int

const (
	KeyCtrlA      Key = 0x01
	KeyCtrlB      Key = 0x02
	KeyCtrlC      Key = 0x03
	KeyCtrlD      Key = 0x04
	KeyCtrlE      Key = 0x05
	KeyCtrlF      Key = 0x06
	KeyCtrlG      Key = 0x07
	KeyBackspace  Key = 0x08
	KeyCtrlH      Key = 0x08
	KeyTab        Key = 0x09 // same as Ctrl+I
	KeyCtrlI      Key = 0x09
	KeyCtrlJ      Key = 0x0A
	KeyCtrlK      Key = 0x0B
	KeyCtrlL      Key = 0x0C
	KeyEnter      Key = 0x0D
	KeyCtrlM      Key = 0x0D
	KeyCtrlN      Key = 0x0E
	KeyCtrlO      Key = 0x0F
	KeyCtrlP      Key = 0x10
	KeyCtrlQ      Key = 0x11
	KeyCtrlR      Key = 0x12
	KeyCtrlS      Key = 0x13
	KeyCtrlT      Key = 0x14
	KeyCtrlU      Key = 0x15
	KeyCtrlV      Key = 0x16
	KeyCtrlW      Key = 0x17
	KeyCtrlX      Key = 0x18
	KeyCtrlY      Key = 0x19
	KeyCtrlZ      Key = 0x1A
	KeyEscape     Key = 0x1B
	KeyBackspace2 Key = 0x7f
	KeyUnknown    Key = 256 + iota
	KeyShiftTab
	KeyDelete
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
)

func ReadKeyEvent() Key {
	d := [10]byte{}
	ev := termbox.PollRawEvent(d[:])
	switch ev.N {
	case 1:
		return Key(d[0])
	case 3:
		return parse3(d[0:3])
	case 4:
		return parse4(d[0:4])
	}
	return KeyUnknown
}

func parse3(d []byte) Key {
	if len(d) != 3 {
		log.Println("expected a three byte event, got", d)
		return KeyUnknown
	}
	switch d[0] {
	case 0x1b: // escape
		switch d[1] {
		case 0x4f:
			switch d[2] {
			case 0x41:
				return KeyArrowUp
			case 0x42:
				return KeyArrowDown
			case 0x43:
				return KeyArrowRight
			case 0x44:
				return KeyArrowLeft
			}
		case 0x5b:
			switch d[2] {
			case 0x5A:
				return KeyShiftTab
			}
		}
	}
	return KeyUnknown
}

func parse4(d []byte) Key {
	if len(d) != 4 {
		log.Println("expected a four byte event, got", d)
		return KeyUnknown
	}
	switch d[0] {
	case 0x1b: // escape
		switch d[1] {
		case 0x5b:
			switch d[2] {
			case 0x33:
				switch d[3] {
				case 0x7E:
					return KeyDelete
				}
			}
		}
	}
	return KeyUnknown
}