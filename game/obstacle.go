package game

type obstacle struct {
	height int
	*Vec2
}

func newObstacle(h int, pos *Vec2) *obstacle {
	return &obstacle{
		height: h,
		Vec2:   pos,
	}
}

func (o *obstacle) overlaps(pos Vec2) bool {
	colOverlap := o.X == pos.X
	heightOverlap := pos.Y > o.Y-o.height
	return colOverlap && heightOverlap
}

type obstacles []*obstacle

func newObstacles(obs ...*obstacle) obstacles {
	return obstacles(obs)
}

func (obs *obstacles) remove() {
	(*obs)[0] = nil
	*obs = (*obs)[1:]
}

func (obs *obstacles) add(o *obstacle) {
	*obs = append(*obs, o)
}
