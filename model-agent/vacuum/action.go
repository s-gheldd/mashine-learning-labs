package vacuum

const (
	Left Action = iota
	Right
	Clean
	Stop
)

type Action int

func (a Action) String() string {
	switch a {
	case Left:
		return "LEFT"
	case Right:
		return "RIGHT"
	case Clean:
		return "CLEAN"
	case Stop:
		return "STOP"
	}
	return "UNKNOWN"
}
