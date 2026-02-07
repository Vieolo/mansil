package main

import (
	"fmt"
	"time"

	"github.com/vieolo/mansil"
)

func main() {
	fmt.Println("One")
	time.Sleep(2 * time.Second)
	fmt.Println(mansil.CursorUp1 + mansil.ClearLine + "two")
}
