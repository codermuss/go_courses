package main

func main() {
	foo := 33

	switch {
	case foo == 1:
		println("one")

	case foo == 2:
		println("two")
	case foo == 3:
		println("three")

	default:
		println("other...")

	}
}
