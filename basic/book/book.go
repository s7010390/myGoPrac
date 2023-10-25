package book

type book struct {
	Name string
}

func New() book {
	return book{Name: "som"}
}
