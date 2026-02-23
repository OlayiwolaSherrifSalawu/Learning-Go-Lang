package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(messsage string) error
}

type ConsoleLogger struct{}

type FileLogger struct {
	fileName string
}

func (cl ConsoleLogger) Log(message string) error {
	fmt.Printf("[Console :] %s \n", message)
	return nil
}
func (fl FileLogger) Log(message string) error {
	f, err := os.OpenFile(fl.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("Failed to open %s : %w \n", fl.fileName, err)
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "[FILE] %s \n", message)
	if err != nil {
		return fmt.Errorf("Failed to open %s : %w \n", fl.fileName, err)
	}
	return nil
}

func RunApplication(logger Logger, message string) {
	err := logger.Log(message)
	if err != nil {
		fmt.Printf("Error logging message: %v\n", err)
	}
}

func main() {
	ConsoleMessage := ConsoleLogger{}
	FileMessage := FileLogger{fileName: "test.log"}
	fmt.Println("------Console Logger-------")
	RunApplication(ConsoleMessage, "By the end of this year Olayiwola would be a milloniare!")
	RunApplication(ConsoleMessage, "If I improve by 1percentage everyday I would be 37 percent better at the end of the year! ")
	fmt.Println("------File Logger-------")
	RunApplication(FileMessage,"I of the opinion that the little struggles in life makes, life more beautiful!")
	RunApplication(FileMessage,"I love afang especially with Pounded yam and goat meat!")
}
