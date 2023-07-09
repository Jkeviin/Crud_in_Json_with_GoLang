package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Hecho  bool   `json:"hecho"`
}

func List(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {

		status := " "

		if task.Hecho {
			status = "✓"
		}

		fmt.Printf("[%s] %d - %s\n", status, task.ID, task.Titulo)

	}
}

func Add(tasks []Task, titulo string) []Task {
	newTask := Task{
		ID:     GetNextID(tasks),
		Titulo: titulo,
		Hecho:  false,
	}
	tasks = append(tasks, newTask)
	return tasks
}

func Delete(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			/*
				Explicacion de lo siguiente
				tasks[:i] -> Toma desde el inicio hasta la posición i
				tasks[i+1:] -> Toma desde la posición i+1 hasta el final

				Entonces, lo que hace es tomar desde el inicio hasta la posición i y luego
				toma desde la posición i+1 hasta el final y los une en un solo slice

				["a", "b", "c", "d", "e"]
				tasks[:2] -> ["a", "b"]
				tasks[3:] -> ["d", "e"]
				["a", "b"] + ["d", "e"] -> ["a", "b", "d", "e"]

			*/
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks
		}
	}
	return tasks
}

func Complete(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Hecho = true
			break
		}
	}
	return tasks
}

/////////////

func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	} // Esto es para cuando no hay tareas

	lastTask := tasks[len(tasks)-1].ID + 1 // Toma la el id de la última tarea y le suma 1
	return lastTask
}

func SaveTasks(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks) // Marshal convierte un arreglo de bytes a un json
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0) // Truncate borra el contenido del archivo
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.Write(bytes) // Escribimos los bytes en el archivo
	if err != nil {
		panic(err)
	}

	err = writer.Flush() // Flush escribe los bytes en el archivo
	if err != nil {
		panic(err)
	}
}
