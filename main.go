package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	condition := true
	var notepad []string
	var maxSize int

	fmt.Print("Enter the maximum number of notes: ")
	strSize, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	maxSize, err = strconv.Atoi(strings.Trim(strSize, " \r\n"))

	if err != nil {
		log.Fatal(err)
	}

	for condition == true {
		fmt.Print("Enter command and data: ")

		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		command, data := split(input)

		switch command {

		case "exit":
			fmt.Println("[Info] Bye!")
			condition = false
		case "clear":
			notepad = nil
			fmt.Println("[OK] All notes were successfully deleted")
		case "create":
			if len(notepad) == maxSize {
				fmt.Println("[Error] Notepad is full")
			} else if data == "" {
				fmt.Println("[Error] Missing note argument")
			} else {
				notepad = append(notepad, data)
				fmt.Println("[OK] The note was successfully created")
			}
		case "list":
			if len(notepad) == 0 {
				fmt.Println("[Info] Notepad is empty")
			} else {
				for idx, data := range notepad {
					fmt.Printf("[Info] %d: %s\n", idx+1, data)
				}
			}
		case "update":
			if data == "" {
				fmt.Println("[Error] Missing position argument")
			} else {
				notepad = update(data, notepad)
			}
		case "delete":
			if data == "" {
				fmt.Println("[Error] Missing position argument")
			} else {
				notepad = delete(data, notepad)
			}
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func split(text string) (command string, data string) {
	textArr := strings.Split(text, " ")
	command = strings.Trim(textArr[0], " \r\n")
	data = strings.Trim(strings.Join(textArr[1:], " "), " \r\n")
	return
}

func update(data string, notepad []string) []string {
	position, note := split(data)

	positionIdx, err := strconv.Atoi(position)

	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", position)
	} else if note == "" {
		fmt.Println("[Error] Missing note argument")
	} else {
		notepad[positionIdx-1] = note
		fmt.Printf("[OK] The note at position %d was successfully updated\n", positionIdx)
	}

	return notepad
}

func delete(data string, notepad []string) []string {
	positionIdx, err := strconv.Atoi(data)

	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", data)
	} else if positionIdx-1 > len(notepad)-1 {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", positionIdx, len(notepad))
	} else {
		if len(notepad) == 1 {
			notepad = nil
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", positionIdx)
		} else {
			notepad = append(notepad[:positionIdx-1], notepad[positionIdx:]...)
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", positionIdx)
		}
	}

	return notepad
}
