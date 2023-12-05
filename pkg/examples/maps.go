package examples

import "fmt"

func maps() {
	var m map[string]int

	fmt.Println(m)
	m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["bar"] = 99

	delete(m, "foo")

	m["baz"] = 418

	fmt.Println(m)

	fmt.Println(m["foo"])
	v, ok := m["foo"]
	fmt.Println(v, ok, m)
}
