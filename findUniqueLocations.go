package main

func main() {}

func FindUniqueLocations(input string) int {
	x, y := 0, 0

	type coord struct { 
		X int 
		Y int 
	}

	gmap := make(map[coord]bool)
	for _, dir := range input {
		switch dir {
		case 'E':
			x++
		case 'W':
			x--
		case 'N':
			y++
		case 'S':
			y--
		}
		gmap[coord{x, y}] = true
	}

	return len(gmap)
}
