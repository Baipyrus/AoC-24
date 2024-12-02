package registry

type Challenge struct {
	Name string
	Exec func()
}

var challenges []Challenge

func Register(name string, exec func()) {
	challenges = append(
		challenges,
		Challenge{name, exec})
}

func Get() []Challenge {
	return challenges
}
