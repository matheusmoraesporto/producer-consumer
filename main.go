package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"unisinos/so/tga/analyser"
	"unisinos/so/tga/conveyor"
	"unisinos/so/tga/product"
	"unisinos/so/tga/utils"
)

const (
	endMsg           = "\n\nTodos os produtos foram processados"
	inputParamPrefix = "--in="
	lenArgs          = 2
	maxFlow          = 10
	minFlow          = 3
	startMsg         = "Os produtos começarão a passar pela esteira."
)

func main() {
	n, err := inputValidations()
	if err != nil {
		fmt.Print(err)
		return
	}
	runConveyor(n)
}

func inputValidations() (int, error) {
	if len(os.Args) != lenArgs {
		msg := fmt.Sprintf("A quantidade de parâmetos informada está incorreta.\nEx de entrada:\n\tgo run . --in=3\n")
		return -1, errors.New(msg)
	}

	inputArg := strings.Split(os.Args[1], inputParamPrefix)[1:][0]
	n, err := strconv.Atoi(inputArg)
	if err != nil {
		return -1, errors.New("Parâmetro \"--in\" inválido.\n")
	}

	if n < minFlow || n > maxFlow {
		return -1, errors.New("Parâmetro \"--in\" deve ser maior que 3 e menor que 10.\n")
	}

	return n, nil
}

func runConveyor(n int) {
	analysers := []analyser.Analyser{
		analyser.NewAnalyser("1"),
		analyser.NewAnalyser("2"),
	}
	c := conveyor.NewConveyor(n)

	var wg sync.WaitGroup
	wg.Add(n + len(analysers))

	fmt.Println(startMsg)
	for i := 0; i < n; i++ {
		go func() {
			p := product.NewProduct()
			c.AddProduct(p)
			wg.Done()
		}()
	}

	for i := 0; i < len(analysers); i++ {
		fmt.Printf("O analisador %s começará a processar os produtos da esteira\n", analysers[i].Id)
		go analysers[i].ConsomeConvenyor(&c, &wg)
	}

	utils.PrintSeparator()
	wg.Wait()
	fmt.Println(endMsg)
	utils.PrintSeparator()

	for _, a := range analysers {
		a.AnalyseProducts()
	}
}
