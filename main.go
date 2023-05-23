package main

import (
	"log"
	"os"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	commandline "github.com/pius706975/backend/commandLine"
)

func init()  {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main()  {
	
	err := commandline.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}