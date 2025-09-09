package main

import (
	"fmt"

	"idmangler/idmangler"
)

func main() {
	item, error := idmangler.DecodeItemObject("󰀀󰄀󰉁󶙴󶕲󷍨󶽣󶬀󰌅󰀸󵴉󶱈󵌊󵌃󷰄󰐄󳆌󶀅󰋿")
	if error != nil {
		panic(error)
	}

	fmt.Println(item)
}