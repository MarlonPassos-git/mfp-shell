package interfaces

type Command struct {
	Name    string
	Handler func(args *[]string)
}
