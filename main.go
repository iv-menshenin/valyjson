package main

import "valyjson/generator"

func main() {
	g := generator.New("./test/struct.go")
	if err := g.Parse(); err != nil {
		panic(err)
	}
	g.BuildFillers()
	g.Print("./test/struct_json.go")
}
