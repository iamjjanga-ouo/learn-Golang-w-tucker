package main

import (
	"fmt"
	"time"
)

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) { // Event 인터페이스 구현
	m.listener = listener
}

func (m *Mail) OnRecv() { // 등록된 listener의 OnFire() 호출
	time.Sleep(3 * time.Second) // 3초 슬립
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() { // EventListener 인터페이스 구현
	// 알람
	fmt.Println("알람! 알람!")
}

func main() {
	var mail = &Mail{}
	var listener EventListener = &Alarm{}

	mail.Register(listener)
	mail.OnRecv() // 알람이 울리게 됩니다.
}
