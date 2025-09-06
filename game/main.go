package main // определили пакет main

import "fmt" // импортировали пакет fmt

// функция для входа в программу
func main() {
	var name string // объявили переменную name типа string
	fmt.Scan(&name) // считали переменную name с потока ввода

	// функция Print() печатает сообщение в консоль
	fmt.Print("Hello, ", name)
}
