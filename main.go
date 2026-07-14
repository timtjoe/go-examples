package main

import (
	"fmt"
	intermediate "go-examples/intermediate"
)

func main() {
fmt.Println("Hello, World!")
go intermediate.Say("hello")
go intermediate.Say("world")
intermediate.SyncMutex()
intermediate.SyncWait()
}