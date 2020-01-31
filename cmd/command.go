package cmd

type Command interface {
	Run() error
}
