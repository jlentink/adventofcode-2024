package main

type Pair struct {
	Left  int
	Right int
}

func (p *Pair) Distance() int {
	d := p.Left - p.Right
	if d < 0 {
		return -d
	}
	return d
}
