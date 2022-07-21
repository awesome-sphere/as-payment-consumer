package kafka

import (
	"strconv"

	"github.com/awesome-sphere/as-payment-consumer/kafka/interfaces"
	"github.com/awesome-sphere/as-payment-consumer/utils"
	"github.com/segmentio/kafka-go"
)

var CREATE_ORDER_TOPIC string
var UPDATE_ORDER_TOPIC string
var PARTITION int
var KAFKA_ADDR string

func ListTopics(connector *kafka.Conn) map[string]*interfaces.TopicInterface {
	partitions, err := connector.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}
	m := make(map[string]*interfaces.TopicInterface)
	for _, p := range partitions {
		if _, ok := m[p.Topic]; ok {
			m[p.Topic].Partition += 1
		} else {
			m[p.Topic] = &interfaces.TopicInterface{Partition: 1}
		}
	}
	return m
}

func isExisted(connector *kafka.Conn, topic string) bool {
	topic_map := ListTopics(connector)
	_, ok := topic_map[topic]
	return ok
}

func initializeTopic(conn *kafka.Conn, topic string) {
	if !isExisted(conn, topic) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             topic,
				NumPartitions:     PARTITION,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}
}

func InitializeKafka() {
	CREATE_ORDER_TOPIC = utils.GetenvOr("CREATE_ORDER_TOPIC", "create-order")
	UPDATE_ORDER_TOPIC = utils.GetenvOr("UPDATE_ORDER_TOPIC", "update-order")
	var err error
	PARTITION, err = strconv.Atoi(utils.GetenvOr("KAFKA_TOPIC_PARTITION", "5"))
	KAFKA_ADDR = utils.GetenvOr("KAFKA_ADDR", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}

	conn, err := kafka.Dial("tcp", KAFKA_ADDR)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	initializeTopic(conn, CREATE_ORDER_TOPIC)
	initializeTopic(conn, UPDATE_ORDER_TOPIC)

	messageRead(CREATE_ORDER_TOPIC)
	messageRead(UPDATE_ORDER_TOPIC)
}
