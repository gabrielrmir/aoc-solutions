package main

type Region struct {
	letter    rune
	plots     map[[2]int]int
	area      int
	perimeter int
}

func NewRegion(letter rune, plots map[[2]int]int) *Region {
	region := Region{}
	region.letter = letter
	region.plots = plots
	for _, perimeter := range plots {
		region.area += 1
		region.perimeter += perimeter
	}
	return &region
}
