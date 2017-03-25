package vacuum

import (
	"log"
)

type Agent struct {
}

func (a Agent) AgentAction(perception *Perception) Action {
	log.Printf("Agent percieved: %+v \n", *perception)
	if perception.Clean[perception.Pos] {
		if lookLeft(perception.Pos, perception.Clean) {
			return logDecision(Left)
		}
		if lookRight(perception.Pos, perception.Clean) {
			return logDecision(Right)
		}

		return logDecision(Stop)
	}
	return logDecision(Clean)
}

func logDecision(action Action) Action {
	log.Printf("Agent decided to %v", action)
	return action
}

func lookLeft(pos int, clean [3]bool) bool {
	if clean[pos] {
		if pos > 0 {
			return lookLeft(pos-1, clean)
		}
		return false
	}
	return true
}

func lookRight(pos int, clean [3]bool) bool {
	if clean[pos] {
		if pos < 2 {
			return lookRight(pos+1, clean)
		}
		return false
	}
	return true
}
