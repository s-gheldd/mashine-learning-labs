package vacuum

import (
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	field int
	clean [2]bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewWorld() *World {
	world := &World{}
	world.field = rand.Intn(2)
	for i := range world.clean {
		world.clean[i] = rand.Intn(2) == 0
	}
	return world
}

func (w *World) Perception() *Perception {
	cpy := w.clean
	return &Perception{
		Pos:   w.field,
		Clean: cpy,
	}
}

func (w *World) Update(action Action) error {

	switch action {
	case Left:
		if w.field > 0 {
			w.field--
			return nil
		}
	case Right:
		if w.field < 1 {
			w.field++
			return nil
		}
	case Clean:
		w.clean[w.field] = true
		return nil
	case Stop:
		return fmt.Errorf("Agent stopped on field %d, world state %v", w.field, w.clean)
	}
	return fmt.Errorf("Could not perform action %v on position %d in world state %v", action, w.field, w.clean)
}

func (w *World) IsClean() bool {
	return w.clean[0] && w.clean[1]
}
