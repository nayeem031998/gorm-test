package api

import (
	"fmt"
	"log"
	"os"
    "fitpass/api/controller"
	"github.com/joho/godotenv"
)

var server = controller.Server{}

func doesFileExist(fileName string) {
      
	// check if error is "file not exists"
	if _, err := os.Stat(fileName); err == nil {
		if err = godotenv.Load(fileName); err != nil {
			log.Fatalf("Error getting env, not comming through %v", err)
		} else {
			fmt.Println(fmt.Sprintf("We are getting the %s values", fileName))
		}
	}
}
func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	doesFileExist(".fitpass")
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	fmt.Println("NAYEEM", os.Getenv("DB_DRIVER"))
	server.Run(":1234")

}
