package main

import "fmt"

// Проблемы которые были в прошлом коде:
// 1) Не объявлена функция createHugeString, не объявлен пакет main и не было импорта библотеки fmt
// 2) Переменная justString пусть и хранила в себе только первые 100 символов но ссылалась на всю строку вцелом
// 	  из-за чего большая строка из которой делался срез занимала много памяти и к тому же в дальнейшем в коде эта большая строка не использовалась.

// Функция для создания большой строки
func createHugeString(size int) string {
	// Заполняем строку пробелами
	return string(make([]byte, size))
}

var justString string

func someFunc() {
	// Создаем строку размером 1024 байта
	v := createHugeString(1 << 10)

	// Берем только первые 100 символов
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
	fmt.Println("justString len:", len(justString))
}
