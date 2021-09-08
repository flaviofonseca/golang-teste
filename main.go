package main

import (
	"fmt"

	"br.com.teste/config/database"
	"br.com.teste/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Teste")

	r := gin.Default()

	controllers.DefineRouters(r)
	database.ConnectDataBase()

	r.Run()
}
