package modbus

import (
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

func ModbusConn(value int16) {
	var client *modbus.ModbusClient
	var err error
	URL := "tcp://192.168.1.9:502"

	// for a TCP endpoint
	// (see examples/tls_client.go for TLS usage and options)
	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     URL,
		Timeout: 1 * time.Second,
	})
	defer client.Close()

	if err != nil {
		// error out if client creation failed
		fmt.Printf("Could not connect to %s", URL)
	}

	// now that the client is created and configured, attempt to connect
	err = client.Open()
	if err != nil {
		// error out if we failed to connect/open the device
		// note: multiple Open() attempts can be made on the same client until
		// the connection succeeds (i.e. err == nil), calling the constructor again
		// is unnecessary.
		// likewise, a client can be opened and closed as many times as needed.
	}

	// by default, 16-bit integers are decoded as big-endian and 32/64-bit values as
	// big-endian with the high word first.
	// change the byte/word ordering of subsequent requests to little endian, with
	// the low word first (note that the second argument only affects 32/64-bit values)
	client.SetEncoding(modbus.LITTLE_ENDIAN, modbus.LOW_WORD_FIRST)

	// write -200 to 16-bit (holding) register 100, as a signed integer
	var s int16 = value
	err = client.WriteRegister(2048, uint16(s))

	// close the TCP connection/serial port
}
