package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    const filename = "tasks.txt"

    tasks, err := LoadTasks(filename)
    if err != nil {
        fmt.Println("Ошибка загрузки задач:", err)
        return
    }

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("1. Показать все задачи")
        fmt.Println("2. Добавить задачу")
        fmt.Println("3. Удалить задачу")
        fmt.Println("4. Выйти")

        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            fmt.Println("Список задач:")
            for _, task := range tasks {
                fmt.Printf("%d. %s\n", task.ID, task.Name)
            }
        case "2":
            fmt.Println("Введите название задачи:")
            scanner.Scan()
            name := scanner.Text()
            tasks = AddTask(tasks, name)
            err := SaveTasks(filename, tasks)
            if err != nil {
                fmt.Println("Ошибка сохранения задачи:", err)
            }
        case "3":
            fmt.Println("Введите ID задачи для удаления:")
            scanner.Scan()
            idStr := scanner.Text()
            id, err := strconv.Atoi(idStr)
            if err != nil {
                fmt.Println("Неверный ID")
                continue
            }
            tasks = DeleteTask(tasks, id)
            err = SaveTasks(filename, tasks)
            if err != nil {
                fmt.Println("Ошибка удаления задачи:", err)
            }
        case "4":
            return
        default:
            fmt.Println("Неверный выбор, попробуйте снова.")
        }
    }
}
