package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"serasa/routes"
	"serasa/utils"

	"github.com/joho/godotenv"
)



func main() {
  godotenv.Load()
  
	router := routes.GetRouter()
	PORT := os.Getenv("APP_PORT")
  fmt.Println("API Started.")
	err := http.ListenAndServe(PORT, utils.IsAuthorized(router))
	log.Fatal(err)

}
