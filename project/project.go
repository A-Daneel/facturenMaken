package project

const (
	format string = "%d : %s\n"
)

type Project struct {
    Name string
}

func (p Project) Title()  string { return p.Name }

func (p Project) FilerValue() string { return p.Name }
