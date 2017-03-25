package main

import "github.com/s-gheldd/mashine-learning-labs/reflex-agent/vacuum"
import "log"
import "os"

func main() {
	world := vacuum.NewWorld()
	agent := &vacuum.Agent{}

	for {
		action := agent.AgentAction(world.Perception())
		if err := world.Update(action); err != nil {
			if world.IsClean() {
				log.Println(err)
				os.Exit(0)
			}
			log.Fatal(err)
		}
	}
}
