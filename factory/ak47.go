package factory

type Ak47 struct {
	gun
}

func (a Ak47) GetName() string {
	//TODO implement me
	return a.name
	//panic("implement me")
}

func (a Ak47) setPower(power int) {
	//TODO implement me
	a.power = power
	//panic("implement me")
}

func newAk47() IGun {
	return &Ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
