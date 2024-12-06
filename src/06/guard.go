package main

type Direction int

const (
	DirectionRight Direction = iota
	DirectionDown
	DirectionLeft
	DirectionUp
)

var directionRune = map[rune]Direction{
	'>': DirectionRight,
	'v': DirectionDown,
	'<': DirectionLeft,
	'^': DirectionUp,
}

type Position struct {
	x      int
	y      int
	facing Direction
}

type Guard struct {
	x             int
	y             int
	facing        Direction
	placesVisited map[[2]int]bool
}

func (g *Guard) IsOutOfBounds(world *World) bool {
	return g.x < 0 || g.y < 0 || g.x >= world.width || g.y >= world.height
}

func (g *Guard) TurnRight() {
	g.facing = (g.facing + 1) % 4
}

func (g *Guard) MoveForward(world *World) bool {
	newX := g.x
	newY := g.y
	if g.facing == DirectionRight {
		newX += 1
	} else if g.facing == DirectionDown {
		newY += 1
	} else if g.facing == DirectionLeft {
		newX -= 1
	} else if g.facing == DirectionUp {
		newY -= 1
	}
	if world.HasObstacle(newX, newY) {
		return false
	}
	g.x = newX
	g.y = newY
	g.placesVisited[[2]int{g.x, g.y}] = true
	return true
}

func (g *Guard) SetPosition(pos Position) {
	g.x = pos.x
	g.y = pos.y
	g.facing = pos.facing
	g.placesVisited[[2]int{g.x, g.y}] = true
}

func (g *Guard) GetPosition() Position {
	return Position{
		x:      g.x,
		y:      g.y,
		facing: g.facing,
	}
}
