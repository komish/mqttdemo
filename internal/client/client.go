package client

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Setting = func(*mqtt.ClientOptions)

func TCPConnectionString(host, port string) string {
	return fmt.Sprintf("tcp://%s:%s", host, port)
}

func New(
	id string,
	connectionString string,
	settings ...Setting) mqtt.Client {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectionString)
	opts.SetClientID(id)

	for _, s := range settings {
		s(opts)
	}

	return mqtt.NewClient(opts)
}

func WithUserPass(user, pass string) Setting {
	return func(co *mqtt.ClientOptions) {
		co.SetUsername(user)
		co.SetPassword(pass)
	}
}

func WithDefaultPublishHandler(h mqtt.MessageHandler) Setting {
	return func(co *mqtt.ClientOptions) {
		co.SetDefaultPublishHandler(h)
	}
}

func WithOnConnectHandler(h mqtt.OnConnectHandler) Setting {
	return func(co *mqtt.ClientOptions) {
		co.OnConnect = h
	}
}

func WithConnectionLostHandler(h mqtt.ConnectionLostHandler) Setting {
	return func(co *mqtt.ClientOptions) {
		co.OnConnectionLost = h
	}
}
