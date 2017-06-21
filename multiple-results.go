package main

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("hello", "world")

	println(a, b)
}
