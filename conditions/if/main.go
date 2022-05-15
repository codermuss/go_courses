package main

func main() {
	// we can define variable in if condition variable
	// if we can define like that we can't use the variable out of scope
	if foo := 2; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}
}
