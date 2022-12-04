package conveyor

import (
	"fmt"
	"sync"
	"unisinos/so/tga/product"
)

type Conveyor struct {
	HandledItems int
	Items        chan product.Product
	LenHandle    int
	mu           sync.Mutex
}

func NewConveyor(lenHandle int) (c Conveyor) {
	c.LenHandle = lenHandle
	c.Items = make(chan product.Product)
	return
}

func (c *Conveyor) AddProduct(p product.Product) {
	msg := p.HandleOnConvenyor()
	fmt.Print(msg)

	c.mu.Lock()
	defer c.mu.Unlock()

	c.Items <- p
	c.HandledItems++

	if c.HandledItems == c.LenHandle {
		close(c.Items)
	}
}
