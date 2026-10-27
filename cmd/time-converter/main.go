package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}
	const m = 60
	hours := totalMinutes / m
	mins := totalMinutes % m
	// TODO: объявите константу — количество минут в одном часе.
	// TODO: вычислите полные часы и оставшиеся минуты.
	// TODO: выведите результаты в формате задания 3.

	fmt.Printf("Полных часов: %d\n", hours)
	fmt.Printf("Оставшихся минут: %d\n", mins)

}
