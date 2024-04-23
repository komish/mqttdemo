package main

import (
	"log"
	"os"

	"github.com/komish/mqttdemo/internal/actions"
	"github.com/komish/mqttdemo/internal/client"
	"github.com/komish/mqttdemo/internal/handlers"
)

var (
	confProto = configOrDefault(os.Getenv("MQTTDEMO_PROTO"), "tcp")
	confHost  = configOrDefault(os.Getenv("MQTTDEMO_HOST"), "localhost")
	confPort  = configOrDefault(os.Getenv("MQTTDEMO_PORT"), "1883")
	confTopic = configOrDefault(os.Getenv("MQTTDEMO_TOPIC"), "super_cool_topic")
)

func main() {
	pubLogger := log.New(os.Stdout, "[P-0] ", log.LstdFlags)
	publisher := client.New(
		"demo_publisher",
		client.ConnectionString(confProto, confHost, confPort),
		client.WithConnectionLostHandler(handlers.OnConnectionLostLogTo(pubLogger)),
		client.WithOnConnectHandler(handlers.OnConnectLogTo(pubLogger)),
	)

	subLogger := log.New(os.Stdout, "[S-0] ", log.LstdFlags)
	subscriber := client.New(
		"demo_subscriber",
		client.ConnectionString(confProto, confHost, confPort),
		client.WithDefaultPublishHandler(handlers.LogPublishedMessageTo(subLogger)),
		client.WithConnectionLostHandler(handlers.OnConnectionLostLogTo(subLogger)),
		client.WithOnConnectHandler(handlers.OnConnectLogTo(subLogger)),
	)

	if token := publisher.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("publisher failed to connect with: ", token.Error())
	}

	if token := subscriber.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("subscriber failed to connect with: ", token.Error())
	}

	actions.SubscribeTo(confTopic, subscriber, subLogger)
	actions.PublishContinuouslyTo(confTopic, publisher, pubLogger)
}

func configOrDefault(configValue, fallback string) string {
	if configValue != "" {
		return configValue
	}
	return fallback
}
