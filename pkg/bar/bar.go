package bar

import "github.com/dt/externdemo/pkg/foo"

func Bar() int {
	foo.Foo()
	return 0
}
