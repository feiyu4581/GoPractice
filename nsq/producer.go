package nsq

import (
	"bufio"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"strings"
)

const (
	QuitSignal = "q"
)

type NsqManager struct {
	Producer *nsq.Producer
}

func NewNsqManager(nsqAddr string) *NsqManager {
	producer, err := nsq.NewProducer(nsqAddr, nsq.NewConfig())
	if err != nil {
		log.Fatalf("Init Producer error：%s\n", err.Error())
	}

	return &NsqManager{Producer: producer}
}

func (manager *NsqManager) PublishFromStdin(topic string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Read string from stdin faild, err:%s", err.Error())
		}

		command = strings.TrimSpace(command)
		if strings.ToLower(command) == QuitSignal {
			return
		}

		if err := manager.Producer.Publish(topic, []byte(command)); err != nil {
			log.Fatalf("Nsq publish topic failed, err:%s", err.Error())
		}
	}
}
