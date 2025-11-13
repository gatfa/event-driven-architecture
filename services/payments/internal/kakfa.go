package internal

import (
	"context"
	"encoding/json"

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
	notification *kafka.Writer
	inventory    *kafka.Writer
}

func NewKafkaClient() *KafkaClient {
	inventory := &kafka.Writer{
		Addr:  kafka.TCP(BROKER_ADDRESS),
		Topic: INVENTORY_TOPIC,
	}

	notification := &kafka.Writer{
		Addr:  kafka.TCP(BROKER_ADDRESS),
		Topic: NOTIFICATION_TOPIC,
	}

	return &KafkaClient{
		inventory,
		notification,
	}
}

func sendKafkaMessage(writer *kafka.Writer, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	err = writer.WriteMessages(context.TODO(), msg)

	if err != nil {
		return err
	}

	return nil
}

func (k KafkaClient) SendInventory() error {
	return sendKafkaMessage(k.inventory, "payment_inventory", Inventaire{})
}

func (k KafkaClient) SendNotification() error {
	return sendKafkaMessage(k.notification, "payment_notification", Notification{Action: "Payment processed successfully"})
}

func (k KafkaClient) Close() {
	k.Close()
}
