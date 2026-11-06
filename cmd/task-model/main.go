package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	// TODO: верните "#ID Название [статус], приоритет N".
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
	// TODO: присвойте статус done исходной задаче.
}

func (t *Task) SetPriority(priority int) bool {
	// TODO: проверьте диапазон 1–3, при ошибке объект не менять.
	if priority >= 1 && priority <= 3 {
		t.Priority = priority
		return true
	}
	return false
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	// TODO: верните &tasks[i] при совпадении, иначе nil.
	return nil
}

func main() {
	counter := 0
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}
	task2 := findTaskByID(tasks, 2)
	if task2 != nil {
		task2.Complete()
		task2.SetPriority(3)
	}

	for i := range tasks {
		if tasks[i].Status == "done" {
			counter++
		}
	}
	fmt.Printf("Завершено задач: %d\n\n", counter)
	// TODO: найдите ID=2, завершите и установите приоритет 3.
	// TODO: посчитайте завершённые задачи.
	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
	fmt.Println(tasks[0].Title)
	fmt.Println(tasks[1].Title)
	fmt.Println(tasks[2].Title)
}
