package product

import (
	"math/rand"
	"time"
)

var fruits = map[int]string{
	1:  "maça",
	2:  "banana",
	3:  "morango",
	4:  "pêra",
	5:  "uva",
	6:  "cereja",
	7:  "melancia",
	8:  "limão",
	9:  "manga",
	10: "abacaxi",
}

func (p *Product) GenerateRandomExpirationDate() {
	rand.Seed(time.Now().UnixNano())

	currentYear := time.Now().Year()
	min := time.Date(currentYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(currentYear+1, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	dt := time.Unix(sec, 0)
	p.ExpirationDate = dt
}

func (p *Product) GenerateRandomName() {
	rand.Seed(time.Now().UnixNano())

	x := rand.Intn(10) + 1
	p.Name = fruits[x]
}

func (p *Product) GenerateRandomTime() {
	rand.Seed(time.Now().UnixNano())

	// range de 5 à 30
	x := rand.Intn(25) + 5
	p.HandleTime = time.Duration(x) * time.Second
}

func (p *Product) DefinePrice() {
	rand.Seed(time.Now().UnixNano())

	a := rand.Intn(6) + 1
	b := rand.Float32()
	p.Price = float32(a) + b
}
