package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Task структура для хранения информации о задаче
type Task struct {
    ID   int
    Name string
}

// Функция для загрузки задач из файла
func LoadTasks(filename string) ([]Task, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var tasks []Task
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, ":", 2)
        if len(parts) == 2 {
            id := len(tasks) + 1
            tasks = append(tasks, Task{ID: id, Name: parts[1]})
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}

// Функция для сохранения задач в файл
func SaveTasks(filename string, tasks []Task) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    for _, task := range tasks {
        fmt.Fprintf(file, "%d:%s\n", task.ID, task.Name)
    }

    return nil
}

// Функция для добавления новой задачи
func AddTask(tasks []Task, name string) []Task {
    id := len(tasks) + 1
    tasks = append(tasks, Task{ID: id, Name: name})
    return tasks
}

// Функция для удаления задачи по ID
func DeleteTask(tasks []Task, id int) []Task {
    for i, task := range tasks {
        if task.ID == id {
            return append(tasks[:i], tasks[i+1:]...)
        }
    }
    return tasks
}
