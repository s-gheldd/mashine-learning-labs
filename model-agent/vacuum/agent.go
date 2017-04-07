package vacuum

import (
	"log"
)

type Agent struct {
	beenThere bool
}

func (a *Agent) AgentAction(perception *Perception) Action {
	log.Printf("Agent percieved: %+v agent state %v\n", *perception, a.beenThere)
	if perception.Clean {
		if a.beenThere {
			if perception.Edge {
				return logDecision(Stop)
			}
			return logDecision(Right)
		}

		if perception.Edge {
			a.beenThere = true
			return logDecision(Right)
		}
		return logDecision(Left)
	}
	return logDecision(Clean)
}

func logDecision(action Action) Action {
	log.Printf("Agent decided to %v", action)
	return action
}
