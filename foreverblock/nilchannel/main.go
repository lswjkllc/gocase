package main

func emptyChan() {
	c := make(chan struct{})
	<-c
}

func nilChan() {
	var c chan struct{}
	<-c
}

func main() {
	// 1、空 channel
	emptyChan()

	// 2、nil channel
	nilChan()
}
