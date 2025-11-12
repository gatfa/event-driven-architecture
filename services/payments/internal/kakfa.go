package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	NOTIFICATION_TOPIC = "notifications.central"
	INVENTORY_TOPIC    = "inventories.central"
	BROKER_ADDRESS     = "kafka:29092"
	GROUP_ID           = "payment-group"
)

type Inventaire struct {
}

type Notification struct {
	Action string `json:"action"`
}

type KafkaClient struct {
	reader *kafka.Reader
	writer *kafka.Writer
}

func NewKafkaClient() *KafkaClient {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{BROKER_ADDRESS},
		Topic:   INVENTORY_TOPIC,
		GroupID: GROUP_ID,
	})

	writer := &kafka.Writer{
		Addr:  kafka.TCP(BROKER_ADDRESS),
		Topic: NOTIFICATION_TOPIC,
	}

	return &KafkaClient{
		reader,
		writer,
	}
}

func (k KafkaClient) Read() error {

	for {
		m, err := k.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		var income Inventaire
		if err := json.Unmarshal(m.Value, &income); err != nil {
			log.Printf("Error unmarshaling JSON: %v\n", err)
			continue
		}

		output := fmt.Sprintf("[Payment] Received inventory validation: %s", income)
		log.Println(output)

		// Simulate payment processing delay
		time.Sleep(time.Second * 2)

		payment, err := json.Marshal(Notification{Action: "Payment processed successfully"})
		if err != nil {
			return err
		}

		msg := kafka.Message{
			Key:   []byte("payment_notification"),
			Value: payment,
		}

		err = k.writer.WriteMessages(context.TODO(), msg)

		if err != nil {
			return err
		}
	}

}

func (k KafkaClient) Close() {
	k.Close()
}
