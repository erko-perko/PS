package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Meritev struct {
	vrsta    string
	vrednost float32
}

func meritevTemperature(meritve chan Meritev) {
	for {
		randomTemp := -20.0 + rand.Float32()*35.0
		meritve <- Meritev{"temperatura", randomTemp}
		time.Sleep(100 * time.Millisecond)
	}
}

func meritevVlage(meritve chan Meritev) {
	for {
		randomVlage := 0.0 + rand.Float32()*100.0
		meritve <- Meritev{"vlaga", randomVlage}
		time.Sleep(100 * time.Millisecond)
	}
}

func meritevTlaka(meritve chan Meritev) {
	for {
		randomTlak := 980.0 + rand.Float32()*1100.0
		meritve <- Meritev{"tlak", randomTlak}
		time.Sleep(100 * time.Millisecond)
	}
}

func readKey(input chan bool) {
	fmt.Scanln()
	input <- true
}

func main() {
	meritve := make(chan Meritev, 3)
	input := make(chan bool)

	go meritevTemperature(meritve)
	go meritevVlage(meritve)
	go meritevTlaka(meritve)
	go readKey(input)

	func() {
		for {
			select {
			case meritev := <-meritve:
				fmt.Printf("Vrsta: %s\nVrednost: %f\n\n", meritev.vrsta, meritev.vrednost)
			case <-input:
				fmt.Println("Prisilna zaustavitev.")
				return
			case <-time.After(5 * time.Second):
				if len(meritve) == 0 {
					fmt.Println("Sistem je neodziven že 5 sekund.")
					return
				}
			}
		}
	}()
}
