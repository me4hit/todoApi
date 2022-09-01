package main

import "todoApi/Infra"

func main() {
	r := Infra.InitServer()
	r.Run(":9090")
}
