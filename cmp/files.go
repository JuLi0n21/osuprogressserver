package cmp

type helper struct {
	Name string
	Link string
}

// used to host static files on diffrent services
var files = []helper{
	//grades
	{"XS", "A.png"},
	{"SS", "A.png"},
	{"S", "A.png"},
	{"A", "A.png"},
	{"B", "A.png"},
	{"C", "A.png"},
	{"D", "A.png"},
	{"E", "A.png"},
	{"F", "A.png"},
	//transparent-grades

	//mods

	//additional files
}

func Img(name string) string {
	for _, file := range files {
		if file.Name == name {
			return file.Link
		}
	}
	return "404.png"
}
