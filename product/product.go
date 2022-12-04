package product

import (
	"fmt"
	"time"
)

type Product struct {
	ExpirationDate time.Time
	Name           string
	Price          float32
	HandleTime     time.Duration
}

func NewProduct() (p Product) {
	p.DefinePrice()
	p.GenerateRandomExpirationDate()
	p.GenerateRandomName()
	p.GenerateRandomTime()
	return
}

func (p *Product) HandleOnConvenyor() string {
	fmt.Printf("ESTEIRA: recebendo um(a) %s. Esse produto leverá %v para ser processado.\n", p.Name, p.HandleTime)
	msg := fmt.Sprintf("ESTEIRA: %s processado(a)\n", p.Name)
	return handle(p.HandleTime, msg)
}

func (p *Product) HandleOnAnalyser(idAnalyser string) string {
	handleTime := p.HandleTime / 2
	fmt.Printf("ANALISADOR %s: retirando um(a) %s da esteira. Esse produto leverá %v para ser processado.\n", idAnalyser, p.Name, handleTime)
	msg := fmt.Sprintf("ANALISADOR: %s processado(a)\n", p.Name)
	return handle(handleTime, msg)
}

func handle(handleTime time.Duration, msg string) string {
	for timeout := time.After(handleTime); ; {
		select {
		case <-timeout:
			return msg
		default:
		}
	}
}
