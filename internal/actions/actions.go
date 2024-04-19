package actions

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func PublishContinuouslyTo(topic string, client mqtt.Client, logger *log.Logger) {
	logger.Println("publishing continuously")
	for {
		text := fmt.Sprintf("message, %s", time.Now().Format(time.RFC1123))
		logger.Println("publishing: ", text)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(5 * time.Second)
	}
}

func SubscribeTo(topic string, client mqtt.Client, logger *log.Logger) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	logger.Printf("Subscribed to topic: %s", topic)
}
