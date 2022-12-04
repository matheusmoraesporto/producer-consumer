package analyser

import (
	"fmt"
	"os"
	"sync"
	"time"
	"unisinos/so/tga/conveyor"
	"unisinos/so/tga/product"
	"unisinos/so/tga/utils"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

const (
	badProducts  = "Produtos com má qualidade"
	goodPorducts = "Produtos com boa qualidade"
)

type Analyser struct {
	Products []product.Product
	Id       string
}

func NewAnalyser(id string) (a Analyser) {
	a.Id = id
	return
}

func (a *Analyser) ConsomeConvenyor(convenyor *conveyor.Conveyor, wg *sync.WaitGroup) {
	defer wg.Done()

	for it := range convenyor.Items {
		it.HandleOnAnalyser(a.Id)
		a.Products = append(a.Products, it)
		fmt.Printf("ANALISADOR %s: produto %s consumido\n", a.Id, it.Name)
	}
}

func (a *Analyser) AnalyseProducts() {
	fmt.Printf("Análise realizada pelo analisador %s\n", a.Id)

	var badProducts []product.Product
	var goodProducts []product.Product
	for _, p := range a.Products {
		if IsProductExpired(p) {
			badProducts = append(badProducts, p)
		} else {
			goodProducts = append(goodProducts, p)
		}
	}

	if len(goodProducts) > 0 {
		printGoodProducts(goodProducts)
	}

	if len(badProducts) > 0 {
		printBadProducts(badProducts)
	}

	utils.PrintSeparator()
}

func IsProductExpired(product product.Product) bool {
	return product.ExpirationDate.Before(time.Now())
}

func printBadProducts(products []product.Product) {
	colorOption := table.ColorOptions{
		IndexColumn:  text.Colors{text.FgRed},
		Footer:       text.Colors{text.FgRed},
		Header:       text.Colors{text.FgRed},
		Row:          text.Colors{text.FgRed},
		RowAlternate: text.Colors{text.FgRed},
	}
	printTable(products, badProducts, colorOption)
}

func printGoodProducts(products []product.Product) {
	colorOption := table.ColorOptions{
		IndexColumn:  text.Colors{text.FgHiBlue},
		Footer:       text.Colors{text.FgHiBlue},
		Header:       text.Colors{text.FgHiBlue},
		Row:          text.Colors{text.FgHiBlue},
		RowAlternate: text.Colors{text.FgHiBlue},
	}
	printTable(products, goodPorducts, colorOption)
}

func printTable(products []product.Product, title string, colorOption table.ColorOptions) {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetTitle(title)
	t.AppendHeader(table.Row{"Produto", "Preço", "Data de validade"})

	var rows []table.Row
	for _, p := range products {
		dtString := fmt.Sprintf("%d/%d/%d", p.ExpirationDate.Day(), p.ExpirationDate.Month(), p.ExpirationDate.Year())
		rows = append(rows, table.Row{p.Name, fmt.Sprintf("R$%.2f", p.Price), dtString})
	}
	t.AppendRows(rows)
	t.AppendFooter(table.Row{"", "", "Total", len(products)})

	t.Style().Color = colorOption
	t.Render()
}
