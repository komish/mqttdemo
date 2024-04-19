package handlers

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func LogPublishedMessageTo(logger *log.Logger) mqtt.MessageHandler {
	return func(_ mqtt.Client, m mqtt.Message) {
		logger.Printf(
			`message with id "%d" received via topic "%s": %s`,
			m.MessageID(), m.Topic(), m.Payload(),
		)
	}
}

func OnConnectLogTo(logger *log.Logger) mqtt.OnConnectHandler {
	return func(c mqtt.Client) {
		opts := c.OptionsReader()
		logger.Println(opts.ClientID(), "connected")
	}
}

func OnConnectionLostLogTo(logger *log.Logger) mqtt.ConnectionLostHandler {
	return func(c mqtt.Client, err error) {
		opts := c.OptionsReader()
		logger.Println(opts.ClientID(), "lost connection with error: ", err)
	}
}
