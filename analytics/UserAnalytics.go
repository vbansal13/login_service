package analytics

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/vbansal/login_service/config"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//Analytics data type
type Analytics struct {
	loginKafkaTopic string
}

//NewAnalytics factory method for creating Analytics object
func NewAnalytics() *Analytics {
	appConfig := config.GetInstance()
	return &Analytics{
		loginKafkaTopic: appConfig.KafkaLoginTopic,
	}
}

//UserLoginAttempt Method for writing login activity to a Kafka Topic
func (atics *Analytics) UserLoginAttempt(successful bool, username string, message string) {

	producer, err := atics.newKafkaProducer()

	if err != nil {
		log.Warning("Error creating kafka producer.")
		return
	}

	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for events := range producer.Events() {
			switch ev := events.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Warning("Delivery failed: ", ev.TopicPartition)
				} else {
					log.Warning("Delivered message to topic: ", ev.TopicPartition)
				}
			}
		}
	}()

	mMessage, mErr := atics.createAnalyticsMessage(successful, username, message)
	if mErr != nil {
		return
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &atics.loginKafkaTopic, Partition: kafka.PartitionAny},
		Value:          mMessage,
	}, nil)

	// Wait for message deliveries before shutting down
	producer.Flush(15 * 1000)

}

func (atics *Analytics) createAnalyticsMessage(successful bool, username string, message string) ([]byte, error) {
	kMessage := UserAnalyticsModel{
		Username:        username,
		LoginSuccessful: successful,
		Message:         message,
	}
	mMessage, mErr := json.Marshal(kMessage)
	if mErr != nil {
		log.Warning("Error while creating login attempt message.")
		return nil, mErr
	}
	return mMessage, nil
}

func (atics *Analytics) newKafkaProducer() (*kafka.Producer, error) {
	appConfig := config.GetInstance()
	return kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": appConfig.KafkaBootStrapServer})
}
