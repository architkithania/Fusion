package utils

const (
	North = iota
	East
	South
	West
)

var DirectionOrder = [4]int{North, East, South, West}

func DirectionToString(direction int) string {
	switch direction {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return ""
	}
}
