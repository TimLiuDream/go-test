package lib

import (
	"os"
	in "GolangTraining/jiketime/demo6/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
