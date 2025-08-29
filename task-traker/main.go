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
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
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

func updateDescription(id int, newDescription string) {
	var tasks []TaskData
	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)

		if err != nil {
			panic(err)
		}

		_ = json.Unmarshal(data, &tasks)

		for index, updateTaks := range tasks {
			if id == updateTaks.Id {
				tasks[index].Description = newDescription
			} else {
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

func updateStatus(id int, newStatus string) {

}

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

			newTask := TaskData{Id: idGlobal, Description: taskInfo, Status: "todo", CreatedAt: 12, UpdatedAt: 12}
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

			if updateInstruction == "tarea" {
				fmt.Println("Ingrese el id de la tarea")
				taskId, _ = reader.ReadString('\n')
				taskId = strings.Trim(taskId, "\r\n")

				fmt.Println("Ingrese la nueva descripcion")
				newTaksDescription, _ = reader.ReadString('\n')
				newTaksDescription = strings.Trim(newTaksDescription, "\r\n")
				convertValue, err := strconv.Atoi(taskId)

				if err != nil {
					panic(err)
				}

				updateDescription(convertValue, newTaksDescription)

			} else if updateInstruction == "estado" {
				fmt.Println("Cambiando el estado de la tarea")
			} else {
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
