package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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
var taskInfo string

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

	fmt.Println("Archivo creado exitosamente...")
}

func writeFIle() {

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Ingrese un comando")
		idGlobal := 1

		instruction, _ = reader.ReadString('\n')
		instruction = strings.TrimSpace(instruction)

		if instruction == "add" {
			// fmt.Println("Ingrese el nombre de la tarea")
			taskInfo, _ = reader.ReadString('\n')
			taskInfo = strings.TrimRight(taskInfo, "\r\n")
			newTask := TaskData{Id: idGlobal, Description: taskInfo, Status: "To Do", CreatedAt: 12, UpdatedAt: 12}
			createFile(newTask)
			idGlobal++
		}

		if instruction == "exit" {
			fmt.Println("Saliendo del programa...")
			time.Sleep(1 * time.Second)
			break
		}
	}
}
