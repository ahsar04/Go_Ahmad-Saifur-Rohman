package main

import (
	"fmt"
)

type Car struct {
	Type string
	fuelin float32

}

func (c Car)count() float32{
	var jarak float32
	switch c.Type {
	case "premium":
		jarak = 1.5 * c.fuelin
	case "pertalite":
		jarak = 2 * c.fuelin
	case "pertamax":
		jarak = 3 * c.fuelin
	}
	return jarak
}

func main()  {
	var hasil = Car{"premium",5}
	fmt.Print("jarak tempuh adalah ",hasil.count()," mil")
}