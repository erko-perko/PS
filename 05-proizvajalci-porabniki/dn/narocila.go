package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var promet float64 = float64(0.0)
var stNarocil int = 0
var lock sync.Mutex
var wgProducer, wgConsumer sync.WaitGroup

var narocila = []narocilo{
	izdelek{"Prenosnik", 1200.00, 2.5},
	eknjiga{"Učenje programskega jezika Go", 29.90},
	spletniTecaj{"Uvod v Go", 10, 15.00},
	izdelek{"Pametni telefon", 800.00, 0.3},
	eknjiga{"Napredni Go", 39.90},
	spletniTecaj{"Napredni Go", 15, 20.00},
	izdelek{"Brezžične slušalke", 150.00, 0.2},
	eknjiga{"Go za začetnike", 19.90},
	spletniTecaj{"Go concurrency", 8, 25.00},
	izdelek{"Zunanji disk", 99.99, 0.5},
	izdelek{"Grafična kartica", 450.00, 1.1},
	eknjiga{"Go Patterns", 24.90},
	spletniTecaj{"Go Testing", 5, 30.00},
	izdelek{"Mehanska tipkovnica", 89.99, 0.8},
	eknjiga{"Concurrent Programming in Go", 34.90},
}

type producer struct {
	id       int
	narocilo narocilo
}

type consumer struct {
	id int
}

type narocilo interface {
	obdelaj()
}

type izdelek struct {
	imeIzdelka string
	cena       float64
	teza       float64
}

type eknjiga struct {
	naslovKnjige string
	cena         float64
}

type spletniTecaj struct {
	imeTecaja   string
	trajanjeUre int
	cenaUre     float64
}

func (p *producer) start(id int, interval time.Duration, data chan<- narocilo, quit <-chan struct{}) {
	defer wgProducer.Done()
	p.id = id
	p.narocilo = narocila[rand.Intn(len(narocila))]
	for {
		select {
		case <-quit:
			return
		default:
			data <- p.narocilo
		}
		time.Sleep(interval)
	}
}

func (c *consumer) start(id int, interval time.Duration, data <-chan narocilo) {
	defer wgConsumer.Done()
	c.id = id
	for {
		if narocilo, more := <-data; more {
			narocilo.obdelaj()
		} else {
			return
		}
		time.Sleep(interval)
	}
}

func (i izdelek) obdelaj() {
	lock.Lock()
	promet += i.cena
	stNarocil++
	fmt.Printf("Številka naročila: %d\nIme izdelka: %s\nCena: %.2f €\nTeža: %.2f kg\n\n", stNarocil, i.imeIzdelka, i.cena, i.teza)
	lock.Unlock()
}

func (ek eknjiga) obdelaj() {
	lock.Lock()
	promet += ek.cena
	stNarocil++
	fmt.Printf("Številka naročila: %d\nNaslov knjige: %s\nCena: %.2f €\n\n", stNarocil, ek.naslovKnjige, ek.cena)
	lock.Unlock()
}

func (st spletniTecaj) obdelaj() {
	lock.Lock()
	promet += st.cenaUre * float64(st.trajanjeUre)
	stNarocil++
	fmt.Printf("Številka naročila: %d\nIme tečaja: %s\nTrajanje: %d ur\nCena na uro: %.2f €\n\n", stNarocil, st.imeTecaja, st.trajanjeUre, st.cenaUre)
	lock.Unlock()
}

func readKey(input chan bool) {
	fmt.Scanln()
	input <- true
}

func main() {
	data := make(chan narocilo, 1000)
	quit := make(chan struct{})
	input := make(chan bool)
	nProducers := 5
	nConsumers := 2
	wgProducer.Add(nProducers)
	wgConsumer.Add(nConsumers)

	producers := make([]producer, nProducers)
	consumers := make([]consumer, nConsumers)
	intervalProducers := 500 * time.Millisecond
	intervalConsumers := 500 * time.Millisecond

	for i := range producers {
		go producers[i].start(i, intervalProducers, data, quit)
	}

	for i := range consumers {
		go consumers[i].start(i, intervalConsumers, data)
	}

	go readKey(input)

	func() {
		for range input {
			close(quit)
			wgProducer.Wait()
			close(data)
			wgConsumer.Wait()
			fmt.Printf("Število naročil: %d\n", stNarocil)
			fmt.Printf("Skupni znesek naročil: %.2f €\n", promet)
			return
		}
	}()
}
