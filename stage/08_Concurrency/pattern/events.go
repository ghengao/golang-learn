package main

import "fmt"

func main() {
	btn := MakeButton()
	handlerOne := make(chan string)
	handlerTwo := make(chan string)
	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)
	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler One: " + msg)
		}
	}()

	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler Two: " + msg)
		}
	}()

	btn.TriggerEvent("click", "Button clicked")
	btn.RemoveEventListener("click", handlerTwo)
	btn.TriggerEvent("click", "Button clicked again!")
	fmt.Scanln()

}

type Button struct {
	eventListeners map[string][]chan string
}

func MakeButton() *Button {
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

func (button *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := button.eventListeners[event]; present {
		button.eventListeners[event] = append(button.eventListeners[event], responseChannel)
	} else {
		button.eventListeners[event] = []chan string{responseChannel}
	}
}

func (button *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := button.eventListeners[event]; present {
		for idx, listener_channel := range button.eventListeners[event] {
			if listener_channel == listenerChannel {
				button.eventListeners[event] = append(button.eventListeners[event][:idx], button.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (button *Button) TriggerEvent(event string, response string) {
	if _, present := button.eventListeners[event]; present {
		for _, handler := range button.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
