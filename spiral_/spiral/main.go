package main

type canvas [][]string

func main() {
	a := 10
	h := 3
	canvas := canvas{}
	for i := 0; i < a; i++ {
		canvas = append(canvas, []string{})
		for j := 0; j < a; j++ {
			canvas[i] = append(canvas[i], " ")
		}
	}
	canvas.drawSpiralLine(a, h, 1, 0, 0, "right")
	canvas.print()
}

func (c canvas) drawSpiralLine(a, h, z, i, j int, direction string) {
	if z >= 4 && z%2 == 0 {
		a = a - h - 1
		if a < h {
			return
		}
	}
	if direction == "right" {
		for t := 0; t < a; t++ {
			c[i][j] = "#"
			j++
		}
		c.drawSpiralLine(a, h, z+1, i, j-1, "down")
		return
	} else if direction == "down" {
		for t := 0; t < a; t++ {
			c[i][j] = "#"
			i++
		}
		c.drawSpiralLine(a, h, z+1, i-1, j, "left")
		return
	} else if direction == "left" {
		for t := 0; t < a; t++ {
			c[i][j] = "#"
			j--
		}
		c.drawSpiralLine(a, h, z+1, i, j+1, "up")
		return
	} else if direction == "up" {
		for t := 0; t < a; t++ {
			c[i][j] = "#"
			i--
		}
		c.drawSpiralLine(a, h, z+1, i+1, j, "right")
		return
	} else {
		panic("Invalid direction")
	}
}

func (c canvas) print() {
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c); j++ {
			print(c[i][j])
			print(" ")
		}
		println()
	}
}
