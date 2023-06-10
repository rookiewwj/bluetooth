package main

import (
	"bluetooth/master"
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	s := initSerial()
	engine := master.NewEngine()
	go engine.Serve()
	for true {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf[:n])
		engine.Push(buf[:n])
	}

}

func initSerial() *serial.Port {
	c := &serial.Config{Name: "COM6", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}
	return s
}
