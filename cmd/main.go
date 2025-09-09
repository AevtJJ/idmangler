package main

import (
	"fmt"

	"github.com/AevtJJ/idmangler"
)

func main() {
	item, error := idmangler.DecodeItemObject("󰀀󰄀󰉁󶙴󶕲󷍨󶽣󶬀󰌅󰀸󵴉󶱈󵌊󵌃󷰄󰐄󳆌󶀅󰋿")
	if error != nil {
		panic(error)
	}

	fmt.Println(item)
}