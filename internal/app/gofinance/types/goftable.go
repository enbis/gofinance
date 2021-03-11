package types

type Holders struct {
	Majors []Major
	Tops   []Top
}

type Major struct {
	Percentage string //needs percentage
	Detail     string
}

type Top struct {
	Holder string
	Shares string
	Data   string
	Out    string //needs percentage
	Value  string
}
