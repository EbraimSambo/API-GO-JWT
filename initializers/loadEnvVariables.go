package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("ERRO AO CARREGAR AS VARIAVEIS DE AMBIENTE")
	}

	fmt.Println("Variaveis carregada com sucesso")
}