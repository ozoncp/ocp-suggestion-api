package main

import (
	"fmt"
	"os"
	"time"
)

const ConfigFileName = "config.json"

func main() {
	fmt.Println("This is ocp-suggestion-api")

	//readConfigFile - функтор, реализует открытие и закрытие файла, используя defer
	readConfigFile := func(filename string) error {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println("error when closing config file: ", err)
			}
		}()
		return nil
	}

	//бесконечный цикл, пока не будет прочитана информация из конфиг-файла
	for {
		if errConfig := readConfigFile(ConfigFileName); errConfig != nil {
			fmt.Printf("Error read config file '%s', error: %v\n", ConfigFileName, errConfig)
			fmt.Println("Wait 5 seconds and repeat...")
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Successful read config file:", ConfigFileName)
			break
		}
	}
}
