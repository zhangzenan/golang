package main

func getData[T any](value T) T {
	return value
}

func main() {
	str1 := getData[string]("this is str")
	println(str1)

	num := getData[int](123)
	println(num)

	//类型可以省略
	str2 := getData("this is str")
	println(len(str2))
}
