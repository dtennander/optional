# Optional

This is a small library for us that is strong users of `Optional<>` in Java
and would like the same paradigm in Go.

## Usage
```go
package example

import opt "github.com/DiTo04/optional"

type Queue interface{
	peekNext() (person opt.Optional)
}

type Person interface{
	GetName() string
}

func getNameOfNext(queue Queue) (str opt.Optional) {
	person := queue.peekNext()
	return person.Map(func(p Person) string {
		return p.GetName()
	})
}

func getNameOfNextOrGopher(queue Queue) string {
	return getNameOfNext(queue).OrElse("Gopher!").(string)
}

func isFirstPersonGopher(queue Queue) bool {
	return getNameOfNext(queue).Filter(func(s string) bool {
		return s == "Gopher"
	}).OrElse(false).(bool)
}
```