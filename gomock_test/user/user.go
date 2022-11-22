package user

import "github.com/timliudream/go-test/gomock_test/doer"

const (
	filtered   = "OK"
	unfiltered = "spam"
	Nice       = "nice"
	Bad        = "bad"
)

type User struct {
	// struct while mocking the doer interface
	Doer doer.Doer
}

// method Use using it
func (u *User) Use() error {
	return u.Doer.DoSomething(123, "Hello GoMock")
}

func (u *User) SaySomething(num int) string {
	if num == 3 {
		return u.Doer.SaySomething(unfiltered)
	}
	return u.Doer.SaySomething(filtered)
}

type Student struct{}

func (s *Student) DoSomething(_ int, _ string) error {
	panic("not implemented") // TODO: Implement
}

func (s *Student) SaySomething(kata string) string {
	if kata == filtered {
		return Nice
	}
	return Bad
}
