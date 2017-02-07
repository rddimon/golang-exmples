package raindrops

import "fmt"

const testVersion = 2

func Convert(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0 && n%7 == 0:
		return "PlingPlangPlong"
	case n%3 == 0 && n%5 == 0 && n%7 != 0:
		return "PlingPlang"
	case n%3 == 0 && n%5 != 0 && n%7 == 0:
		return "PlingPlong"
	case n%3 != 0 && n%5 == 0 && n%7 == 0:
		return "PlangPlong"
	case n%3 == 0:
		return "Pling"
	case n%5 == 0:
		return "Plang"
	case n%7 == 0:
		return "Plong"
	default:
		return fmt.Sprintf("%d", n)
	}
}
