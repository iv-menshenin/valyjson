package main

import "valyjson/generator"

func main() {
	g := generator.New("./struct.go")
	if err := g.Parse(); err != nil {
		panic(err)
	}
	g.Analyse()
	g.Print("./struct_json.go")
}
