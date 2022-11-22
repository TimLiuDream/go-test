package doer

type Doer interface {
	DoSomething(int, string) error
	SaySomething(string) string
}
