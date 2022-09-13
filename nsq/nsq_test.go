package nsq

import "testing"

func TestProducer(t *testing.T) {
	manager := NewNsqManager("127.0.0.1:59201")
	manager.PublishFromStdin("from_stdin")
}
