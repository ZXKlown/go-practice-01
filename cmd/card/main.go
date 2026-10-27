package main

import "fmt"

func main() {
	name := "Алексей"
	group := "ЭФБО-02-26"
	const cours = "Бекенд разработка на языке Go"
	completed := 1
	completed = 0
	interest := true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", cours)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Участник: %t\n", interest)

}
