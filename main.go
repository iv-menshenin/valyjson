package main

import (
	"fmt"
	"time"

	"github.com/iv-menshenin/valyjson/generator"
)

func main() {
	started := time.Now()
	g := generator.New("./benchmark/data.go")
	if err := g.Parse(); err != nil {
		panic(err)
	}
	fmt.Printf("PARSED [%v]\n", time.Since(started))

	started = time.Now()
	g.BuildFillers()
	g.BuildJsoners()
	g.Print("./benchmark/data_json.go")

	fmt.Printf("DONE [%v]\n", time.Since(started))
}
