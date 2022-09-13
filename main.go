package main

import "GoPractice/nsq"

func main() {
	manager := nsq.NewNsqManager("172.28.144.1:4160")
	//manager := nsq.NewNsqManager("0.0.0.0:4160")
	//manager := nsq.NewNsqManager("localhost:4160")
	manager.PublishFromStdin("from_stdin")
}
