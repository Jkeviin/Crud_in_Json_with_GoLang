package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/Jkeviin/go-cli-crud/tasks"
)

func main() {

	// os es un paquete que nos permite leer y escribir archivos

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666) // RDWR: Permisos de lectura y escritura, CREATE: Si no existe el archivo lo crea, 0666: Permisos de lectura y escritura

	if err != nil {
		panic(err)
	}

	defer file.Close() // Cerramos la conexión con el archivo

	var tasks []task.Task

	info, err := file.Stat() // Obtenemos información del archivo
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 { // Si el archivo no está vacío
		bytes, err := io.ReadAll(file) // Leemos el archivo, da los datos en bytes
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks) // Convertimos los bytes a un json y lo guardamos en el slice de tareas
		if err != nil {
			panic(err)
		}

	} else {
		tasks = []task.Task{} // Si el archivo está vacío, inicializamos el slice de tareas vacío
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.List(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ingresa el título de la tarea: ")
		titulo, err := reader.ReadString('\n') // \n es para
		if err != nil {
			panic(err)
		}
		titulo = strings.TrimSpace(titulo) // Quitamos los espacios en blanco

		tasks = task.Add(tasks, titulo)

		fmt.Println("Tarea agregada")

		task.SaveTasks(file, tasks)

	case "complete":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ingresa el id de la tarea a completar: ")
		id, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		id = strings.TrimSpace(id)

		num, err := strconv.Atoi(id) // se convierte el string a int
		if err != nil {
			fmt.Println("El id debe ser un número")
			return
		}

		tasks = task.Complete(tasks, num)

		fmt.Println("Tarea completada")

		task.SaveTasks(file, tasks)

	case "delete":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ingresa el id de la tarea a eliminar: ")
		id, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		id = strings.TrimSpace(id)

		num, err := strconv.Atoi(id) // se convierte el string a int
		if err != nil {
			fmt.Println("El id debe ser un número")
			return
		}

		tasks = task.Delete(tasks, num)

		fmt.Println("Tarea eliminada")

		task.SaveTasks(file, tasks)

	default:
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Uso: go-cli-crud [list|add|complete|delete]")
}

/* func printTasks(tasks []task.Task) {
	for _, task := range tasks {
		fmt.Printf("%d - %s\n", task.ID, task.Titulo)
	}
} */
