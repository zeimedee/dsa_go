package main

import "fmt"

type computer interface {
	insertIntoLightingPort()
}

type client struct{}

func (c *client) insertsLightingConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts lighting connector into computer")
	com.insertIntoLightingPort()
}

type mac struct{}

func (m *mac) insertIntoLightingPort() {
	fmt.Println("lightning connector is plugged into mac machine")
}

type windows struct{}

func (w *windows) InsertIntoUsbPort() {
	fmt.Println("USB is connector plugged into windows machine")
}

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertIntoLightingPort() {
	fmt.Println("Adapter converts lighting signal to USB")
	w.windowsMachine.InsertIntoUsbPort()
}

func main() {

	client := &client{}
	mac := &mac{}

	client.insertsLightingConnectorIntoComputer(mac)

	windowsMachine := &windows{}

	windowsMachineAdapter := &windowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.insertsLightingConnectorIntoComputer(windowsMachineAdapter)
}
