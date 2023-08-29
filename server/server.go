package main

import (
	"fmt"
	"heimigo/server/modbus"
	"heimigo/server/mqtt"
	"heimigo/server/weatherAPI"
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
		// weatherAPI.ReadWeather()
		w := weatherAPI.ReadWeather()
		var SUN_GT1 float64 // Givare i pannrum
		var SUN_GT2 float64 // Givare på taket
		var SUN_Hysteres float64 = 3.0
		if SUN_GT2 > SUN_GT1+SUN_Hysteres && w.CheckSunIsUp() {
			modbus.ModbusConn(4)
		} else {
			modbus.ModbusConn(0)
		}
		for _, val := range values {
			fmt.Println(val.Print())

		}
		time.Sleep(5 * time.Second)
	}

}
