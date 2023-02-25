package main

type Shape struct {
}

func (s *Shape) Move(gas rune) {
}

func day17(input string) int {
	shapes := []Shape{}

	for _, shape := range shapes {
		for _, gas := range input {
			shape.Move(gas)
		}
	}
	return 0
}
