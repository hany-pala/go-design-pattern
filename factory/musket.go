package factory

type musket struct {
	gun
}

func (m musket) setPower(power int) {
	//TODO implement me
	m.power = power
	//panic("implement me")
}

func (m musket) GetName() string {
	//TODO implement me
	return m.name
	//panic("implement me")
}

func newMusket() IGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
