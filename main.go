package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JoaoOliveira-Dev/primeiro-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	// O router := gin.New() creiará um novo roteador gin. Os roteadores podem ser inicializados de duas maneiras, uma usando o gin.New() e a outra usando o gin.Default()
	// A diferença é que o gin.New() inicializa um roteador sem nenhum middleware enquanto o gin.Default() inicializa o roteador com logger e middleware de recovery
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	router.Run(":3001")

	if err := router.Run(":3001"); err != nil{
		log.Fatal(err)
	}

	fmt.Println(os.Getenv("TESTE"))
}