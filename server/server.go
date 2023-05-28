package main

import (
	"fmt"
	// "heimigo/server/modbus"
	"heimigo/server/mqtt"
	"time"
)

func main() {
	fmt.Println("Hello")

	ch := make(chan mqtt.TT)
	var mqtt_test mqtt.TT
	values := make(map[string]mqtt.TT)

	go mqtt.ConnectMQTT(ch)
	for {
		go func() {
			mqtt_test = <-ch
			fmt.Printf("\n------\n%s\n-----\n", mqtt_test.Value)
			values[mqtt_test.Topic] = mqtt_test
		}()
		// modbus.ModbusConn(4)
		// weatherAPI.ReadWeather()
		for _, val := range values {
			fmt.Println(val.Print())

		}
		time.Sleep(5 * time.Second)
	}

}
