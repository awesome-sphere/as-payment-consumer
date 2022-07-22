package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/awesome-sphere/as-payment-consumer/db"
	"github.com/awesome-sphere/as-payment-consumer/db/models"
	"github.com/awesome-sphere/as-payment-consumer/internal"
	"github.com/awesome-sphere/as-payment-consumer/kafka/interfaces"
	"github.com/segmentio/kafka-go"
)

func readerInit(topic_name string, groupBalancers []kafka.GroupBalancer) *kafka.Reader {
	var groupID string
	switch topic_name {
	case CREATE_ORDER_TOPIC:
		groupID = "create-order-consumer"
	case UPDATE_ORDER_TOPIC:
		groupID = "update-order-consumer"
	}

	config := kafka.ReaderConfig{
		Brokers:        []string{KAFKA_ADDR},
		Topic:          topic_name,
		GroupID:        groupID,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		GroupBalancers: groupBalancers,
	}
	r := kafka.NewReader(config)
	return r
}

func createOrderRead(r *kafka.Reader, topic_name string) {
	defer r.Close()
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("could not read message: " + err.Error())
			continue
		}

		fmt.Printf("Reading message at topic [%v] partition [%v] offset [%v]: %s", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		var val interfaces.CreateOrderMessageInterface

		err = json.Unmarshal(msg.Value, &val)

		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err.Error())
			continue
		}
		db.CreateUserHistory(val.UserID, val.TimeSlotId, val.TheaterId, val.SeatNumber, val.Price)
	}
}

func updateOrderRead(r *kafka.Reader, topic_name string) {
	defer r.Close()
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("could not read message: " + err.Error())
			continue
		}

		fmt.Printf("Reading message at topic [%v] partition [%v] offset [%v]: %s", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		var val interfaces.UpdateOrderMessageInterface

		err = json.Unmarshal(msg.Value, &val)

		if err != nil {
			log.Printf("Failed to unmarshal message: %v]n", err.Error())
			continue
		}
		if len(val.SeatNumber) == 0 {
			log.Printf("This order does not have any seat number\n")
			continue
		}
		res := db.UpdateUserHistory(val.UserID, val.TimeSlotId, val.TheaterId, val.SeatNumber[0], models.OrderStatus(val.Status))
		if res {
			internal.NotifyOtherServices(val)
		}
	}
}

func messageRead(topic_name string) {
	groupBalancers := make([]kafka.GroupBalancer, 0)
	groupBalancers = append(groupBalancers, kafka.RangeGroupBalancer{})

	readers := make([]*kafka.Reader, 0)
	for i := 0; i < PARTITION; i++ {
		readers = append(readers, readerInit(topic_name, groupBalancers))
	}
	for _, reader := range readers {
		switch topic_name {
		case CREATE_ORDER_TOPIC:
			go createOrderRead(reader, topic_name)
		case UPDATE_ORDER_TOPIC:
			timer := time.NewTimer(1*time.Minute + 30*time.Second)
			go func() {
				<-timer.C
				updateOrderRead(reader, topic_name)
			}()
		}
	}
}
