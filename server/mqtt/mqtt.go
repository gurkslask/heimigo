package mqtt

import (
	"fmt"
	"heimigo/server/helpers"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type TT struct {
	Value string
	Topic string
}

func (T TT) Print() string {
	return fmt.Sprintf(`
	Value: %v
	Topic: %s `, T.Value, T.Topic)
}
func (T TT) GetFloat() float64 {
	s := string(T.Value)
	f, err := strconv.ParseFloat(s, 32)
	helpers.CheckErr(err)
	return f
}

func ConnectMQTT(tt chan TT) {
	server := "tcp://192.168.20.157:1883"
	topic := "/mosquitto/data"
	topic2 := "/mosquitto/SUN_GT3"
	topic3 := "/mosquitto/SUN_GT2"
	clientid := "freja" + strconv.Itoa(time.Now().Second())

	connOpts := MQTT.NewClientOptions().AddBroker(server).SetClientID(clientid).SetCleanSession(true)

	subscribe := func(client MQTT.Client, itopic string) {
		token := client.Subscribe(itopic, 1, func(client MQTT.Client, message MQTT.Message) {
			//fmt.Printf("Message: %s, Topic: %s", message.Payload(), itopic)
			t := TT{string(message.Payload()), itopic}
			t.GetFloat()
			tt <- t
		})

		token.Wait()
		//fmt.Printf("Subscribed to topic %s\n", itopic)
	}

	connOpts.OnConnect = func(c MQTT.Client) {
		subscribe(c, topic)
		subscribe(c, topic2)
		subscribe(c, topic3)
	}

	client := MQTT.NewClient(connOpts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", server)
	}

}
