package main

type X interface {
	testX()
}

type Y interface {
	testY()
}
type Z interface {
	X
	Y
	testZ()
}