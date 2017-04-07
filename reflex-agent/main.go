package main

import (
	"flag"
	"log"

	"github.com/s-gheldd/mashine-learning-labs/reflex-agent/vacuum"
)

var testCases int

func init() {
	flag.IntVar(&testCases, "c", 1, "Number of test cases.")
	flag.Parse()
}

func main() {

	for i := 0; i < testCases; i++ {

		world := vacuum.NewWorld()
		agent := &vacuum.Agent{}

	inner:
		for {
			action := agent.AgentAction(world.Perception())
			if err := world.Update(action); err != nil {
				if world.IsClean() {
					log.Println(err)
					break inner
				}
				log.Fatal(err)
			}
		}
	}
}
