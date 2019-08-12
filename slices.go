package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	yslice := make([][]uint8, dy)

	for i := range yslice {
		yslice[i] = make([]uint8, dx)

		for j := range yslice[i] {
			yslice[i][j] = uint8((i + j) / 2)
		}
	}

	return yslice
}

func main() {
	pic.Show(Pic)
}
