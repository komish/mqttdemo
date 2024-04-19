# MQTTDEMO

This is an initial exploration of MQTT libraries available for Golang.

This just runs a simple publisher and subscriber to a single topic. Messages are published every 5 seconds, and the subscriber prints it out.

## Running

This demo only has three configuration options, via environment variables, and assumes a TCP connection to the broker.

`MQTTDEMO_HOST` is the hostname of the broker (localhost)

`MQTTDEMO_PORT` is the port of the broker (1883)

`MQTTDEMO_TOPIC` is the topic the application uses (super_cool_topic)