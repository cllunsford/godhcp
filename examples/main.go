package main

import (
	"github.com/cllunsford/godhcp"
	"fmt"
)

func main() {
	err := godhcp.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}