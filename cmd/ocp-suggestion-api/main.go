package main

import (
	"fmt"
	"os"
	"time"
)

const ConfigFileName = "config.json"

func main() {
	fmt.Println("This is ocp-suggestion-api")

	//ReadConfigFile - функтор, реализует открытие и закрытие файла, используя defer
	ReadConfigFile := func(filename string) (err error) {
		var file *os.File
		file, err = os.Open(filename)
		if err != nil {
			return
		}
		defer func() {
			closeErr := file.Close()
			if err == nil { //Если не было ошибки открытия файла,
				err = closeErr //возвращаем в err ошибку закрытия
			}
		}()
		return
	}

	//бесконечный цикл, пока не будет прочитана информация из конфиг-файла
	for {
		if errConfig := ReadConfigFile(ConfigFileName); errConfig != nil {
			fmt.Printf("Error read config file '%s', error: %v\n", ConfigFileName, errConfig)
			fmt.Println("Wait 5 seconds and repeat...")
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Successful read config file:", ConfigFileName)
			break
		}
	}
}
