package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
)

type StackDriverJson struct {
	message  string
	cpf      string
	module   string
	severity string
}

func main() {

	//https://github.com/buger/jsonparser
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		r := regexp.MustCompile(`(?P<timestamp>\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}).\d+Z\s(?P<regularexpression>{.*}$)`)
		match := r.FindStringSubmatch(text)
		if len(match) > 2 {
			timestamp := match[1]
			jsonString := match[2]
			var jsonObject map[string]interface{}
			err := json.Unmarshal([]byte(jsonString), &jsonObject)
			if err != nil {
				fmt.Println(text)
			} else {
				cpf := jsonObject["cpf"]
				module := jsonObject["module"]
				if cpf == nil {
					cpf = "None"
				}
				logLine := fmt.Sprintf("%s", jsonObject["message"])
				//print timestamp without newline
				fmt.Printf("\033[1;90m%s [%s] [%s] \033[0m", timestamp, cpf, module)

				if jsonObject["severity"] == "INFO" {
					color.White(logLine)
				} else if jsonObject["severity"] == "ERROR" {
					color.Red(logLine)
				} else {
					color.Yellow(logLine)
				}
			}
		} else {
			fmt.Println(text)
		}
	}
}
