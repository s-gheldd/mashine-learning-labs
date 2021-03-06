package main

import (
	"log"

	"flag"

	"github.com/s-gheldd/mashine-learning-labs/chinese-room-agent/vacuum"
)

var testCases int

func init() {
	flag.IntVar(&testCases, "c", 1, "Number of test cases.")
	flag.Parse()
}

func main() {

	for i := 0; i < testCases; i++ {
		world := vacuum.NewWorld()
		agent := vacuum.NewAgent()

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
