package vacuum

import (
	"fmt"
	"log"
)

func NewAgent() *Agent {
	return &Agent{
		memory: make([]Perception, 0),
	}
}

type Agent struct {
	memory []Perception
}

//for 2 places: 2 * (4+3+2+1) = 2 * 4*(5)/2  = 20 (4 = 2^2)
//for 3 places: 3 * (8+7+6+5+4+3+2+1) = 3 * 8 * 9/2 = 108 (8 = 2^3)
var table map[string]Action = map[string]Action{
	// first action states
	"[0, [true true]]":   Stop,
	"[0, [true false]]":  Right,
	"[0, [false true]]":  Clean,
	"[0, [false false]]": Clean,
	"[1, [true true]]":   Stop,
	"[1, [true false]]":  Clean,
	"[1, [false true]]":  Left,
	"[1, [false false]]": Clean,
	// second action states
	"[0, [true false]][1, [true false]]":  Clean,
	"[0, [false true]][0, [true true]]":   Stop,
	"[0, [false false]][0, [true false]]": Right,
	"[1, [true false]][1, [true true]]":   Stop,
	"[1, [false true]][0, [false true]]":  Clean,
	"[1, [false false]][1, [false true]]": Left,
	//third action states
	"[0, [true false]][1, [true false]][1, [true true]]":   Stop,
	"[0, [false false]][0, [true false]][1, [true false]]": Clean,
	"[1, [false true]][0, [false true]][0, [true true]]":   Stop,
	"[1, [false false]][1, [false true]][0, [false true]]": Clean,
	//fourth action states
	"[0, [false false]][0, [true false]][1, [true false]][1, [true true]]": Stop,
	"[1, [false false]][1, [false true]][0, [false true]][0, [true true]]": Stop,
}

func Flatten(memory ...Perception) string {
	concat := ""
	for _, perception := range memory {
		concat += "["
		concat += fmt.Sprint(perception.Pos) + ", "
		concat += fmt.Sprint(perception.Clean)
		concat += "]"
	}
	return concat
}

func (a *Agent) AgentAction(perception *Perception) Action {
	a.memory = append(a.memory, *perception)
	log.Printf("Agent percieved: %+v \n", *perception)
	log.Printf("Agent has memory: %+v \n", a.memory)
	action := lookUp(a.memory)
	return logDecision(action)

}

func lookUp(memory []Perception) Action {
	action, ok := table[Flatten(memory...)]
	if ok {
		return action
	}
	log.Fatal("no action found")
	return Stop
}

func logDecision(action Action) Action {
	log.Printf("Agent decided to %v", action)
	return action
}
