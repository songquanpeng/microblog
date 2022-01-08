package main

import (
	"embed"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

var (
	port  = flag.Int("port", 3000, "specify the server listening port.")
	Token = flag.String("token", "token", "specify the private token.")
)

//go:embed public
var fs embed.FS

func init() {

}

func main() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	flag.Parse()

	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	server := gin.Default()
	SetIndexRouter(server)
	SetApiRouter(server)
	var realPort = os.Getenv("PORT")
	if realPort == "" {
		realPort = strconv.Itoa(*port)
	}
	err = server.Run(":" + realPort)
	if err != nil {
		log.Println(err)
	}
}
