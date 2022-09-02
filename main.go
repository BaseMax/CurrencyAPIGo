package main

import (
	"fmt"
	"log"
)

func main() {
	code := "azadi1"
	c, err := GetCoin(code)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("%#v\n", c)
}
