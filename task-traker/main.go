package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type TaskData struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

const fileName = "task.json"

var instruction string = ""
var taskInfo string // Variable global para agregar mas tareas

func createFile(info TaskData) {
	var tasks []TaskData
	isFileExist := false

	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)
		isFileExist = true

		if err == nil && len(data) > 0 {
			_ = json.Unmarshal(data, &tasks)
		}

	}

	tasks = append(tasks, info)
	newJson, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileName, newJson, 0644)

	if err != nil {
		panic(err)
	}

	if isFileExist {
		fmt.Println("Se ha agregado una nueva tarea...")
	} else {
		fmt.Println("Archivo creado exitosamente...")
	}

}

func updateDescription(id string, newDescription string) {
	var tasks []TaskData
	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)

		if err != nil {
			panic(err)
		}

		_ = json.Unmarshal(data, &tasks)

		convertedId, err := strconv.Atoi(id)

		if err != nil {
			panic(err)
		}

		for index, updateTaks := range tasks {
			if convertedId == updateTaks.Id {
				tasks[index].Description = newDescription
				break
			} else if index == len(tasks)-1 {
				fmt.Println("No se ha podido encontrar la tarea")
				return
			}
		}
		newJson, err := json.MarshalIndent(tasks, "", " ")

		if err != nil {
			panic(err)
		}

		err = os.WriteFile(fileName, newJson, 0644)

		if err != nil {
			panic(err)
		}

		fmt.Println("Descripcion actualizada correctamente")
	}

}

func updateStatus(id string, newStatus string) {
	var tasks []TaskData
	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)

		if err != nil {
			panic(err)
		}

		_ = json.Unmarshal(data, &tasks)
		convertdId, err := strconv.Atoi(id)

		if err != nil {
			panic(err)
		}

		for index, tasksValues := range tasks {
			if convertdId == tasksValues.Id {
				tasks[index].Status = newStatus
				break
			} else if index == len(tasks)-1 {
				fmt.Println("No se pudo encontrar el elemento...")
				return
			}
		}
		newJson, err := json.MarshalIndent(tasks, "", " ")

		if err != nil {
			panic(err)
		}

		err = os.WriteFile(fileName, newJson, 0644)

		if err != nil {
			panic(err)
		}

		fmt.Println("Tarea actualizada correctamente")
		return
	}

	fmt.Println("No se pudo encontrar el archivo")
}

func deleteTasks() {}

func main() {
	reader := bufio.NewReader(os.Stdin)
	idGlobal := 1

	for {
		fmt.Println("Ingrese un comando")

		instruction, _ = reader.ReadString('\n')
		instruction = strings.TrimSpace(instruction)

		if instruction == "add" {
			fmt.Println("Ingrese el nombre de la tarea")
			taskInfo, _ = reader.ReadString('\n')
			taskInfo = strings.TrimRight(taskInfo, "\r\n")
			currentTime := time.Now()
			setTime := currentTime.Format("02-01-2006 15:04:05")

			newTask := TaskData{Id: idGlobal, Description: taskInfo, Status: "todo", CreatedAt: setTime, UpdatedAt: setTime}
			createFile(newTask)

			idGlobal++
		}

		if instruction == "update" {
			updateInstruction := ""
			taskId := ""
			newTaksDescription := ""
			newTasksStatus := ""

			fmt.Println("Escoja que quiere actualizar (tarea, estado)")
			updateInstruction, _ = reader.ReadString('\n')
			updateInstruction = strings.Trim(updateInstruction, "\r\n")

			switch updateInstruction {
			case "tarea":
				fmt.Println("Ingrese el id de la tarea")
				taskId, _ = reader.ReadString('\n')
				taskId = strings.Trim(taskId, "\r\n")

				fmt.Println("Ingrese la nueva descripcion")
				newTaksDescription, _ = reader.ReadString('\n')
				newTaksDescription = strings.Trim(newTaksDescription, "\r\n")

				updateDescription(taskId, newTaksDescription)

			case "estado":
				fmt.Println("Ingrese el id de la tarea")
				taskId, _ = reader.ReadString('\n')
				taskId = strings.TrimRight(taskId, "\r\n")

				fmt.Println("Escriba \"in-progress\" o \"done\" segun su estado")
				newTasksStatus, _ = reader.ReadString('\n')
				newTasksStatus = strings.TrimRight(newTasksStatus, "\r\n")

				if newTasksStatus == "in-progress" || newTasksStatus == "done" {
					updateStatus(taskId, newTasksStatus)
				} else {
					fmt.Println("Debe ingresar solamente \"in-progress\" o \"done\"")
				}

			default:
				fmt.Println("Comando incorrecto, escriba tarea o estado..")
			}
		}

		if instruction == "exit" {
			fmt.Println("Saliendo del programa...")
			time.Sleep(1 * time.Second)
			break
		}
	}
}
