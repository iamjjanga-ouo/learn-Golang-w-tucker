package main

import (
	"main.com/fedex"
	"main.com/koreaPost"
)

// func SendBook(name string, sender *fedex.FedexSender) {
// 	sender.Send(name)
// }

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	// Fedex 전송 객체를 만듭니다.
	// sender := &fedex.FedexSender{}
	// SendBook("어린 왕자", sender)	
	// SendBook("그리스인 조르바", sender)

	// ERROR
	// 타입이 맞지 않습니다.
	// sender := &koreaPost.PostSender{}
	// SendBook("어린 왕자", sender)	
	// SendBook("그리스인 조르바", sender)

	// 인터페이스
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}