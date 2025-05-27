package main

import "fmt"

type customStr string

func (customText customStr) log() {
	fmt.Println(customText)
}

func main() {
	var name customStr

	name = "Afan"

	name.log()

}
