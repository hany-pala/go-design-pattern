package factory

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) SetPower(power int) {
	g.power = power
}

func (g *gun) GetPower() int {
	return g.power
}
