package main

import "myapp/sub"

func main() {
	sub.Export()

	// sub.unexporte() // This will cause a compile-time error because unexporte is unexported
}
