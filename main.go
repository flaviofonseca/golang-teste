package main

import (
	"net/http"
	"time"

	"br.com.teste/config/database"
	"br.com.teste/controllers"
	"github.com/gin-gonic/gin"
	timeout "github.com/vearne/gin-timeout"
)

func main() {
	defaultMsg := `{"code": -1, "msg":"http: Handler timeout"}`

	r := gin.Default()
	r.Use(timeout.Timeout(
		timeout.WithTimeout(3*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout),
		timeout.WithDefaultMsg(defaultMsg)))

	controllers.DefineRouters(r)
	database.ConnectDataBase()

	r.Run()
}
