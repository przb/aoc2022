package util

type Point struct {
	X int
	Y int
}

func (p *Point) SameRow(other *Point) bool {
	return p.Y == other.Y
}
func (p *Point) SameColumn(other *Point) bool {
	return p.X == other.X
}

func (p *Point) Adjacent(other *Point) bool {
	return (IntAbs(p.X-other.X) == 1 && IntAbs(p.Y-other.Y) == 1) ||
		(p.SameColumn(other) && IntAbs(p.Y-other.Y) == 1) ||
		(p.SameRow(other) && IntAbs(p.X-other.X) == 1)
}

func (p *Point) Add(v Vector) *Point {

	return &Point{
		X: p.X + v.X,
		Y: p.Y + v.Y,
	}
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}
