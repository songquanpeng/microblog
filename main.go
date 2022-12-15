package main

import (
	"flag"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"microblog/common"
	"microblog/model"
	"microblog/router"
	"os"
	"strconv"
)

var (
	port     = flag.Int("port", 3000, "the listening port")
	username = flag.String("username", "admin", "username for authentication")
	password = flag.String("password", "123456", "password for authentication")
)

func init() {
	//if os.Getenv("USERNAME") != "" {
	//	*username = os.Getenv("USERNAME")
	//}
	common.Username = *username
	if os.Getenv("PASSWORD") != "" {
		*password = os.Getenv("PASSWORD")
	}
	common.Password = *password
	if os.Getenv("SESSION_SECRET") != "" {
		common.SessionSecret = os.Getenv("SESSION_SECRET")
	}
	if os.Getenv("SQLITE_PATH") != "" {
		common.SQLitePath = os.Getenv("SQLITE_PATH")
	}
}

func main() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	flag.Parse()
	err := model.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := model.CloseDB()
		if err != nil {
			log.Fatal(err)
		}
	}()
	server := gin.Default()
	store := cookie.NewStore([]byte(common.SessionSecret))
	server.Use(sessions.Sessions("session", store))
	router.SetRouter(server)
	var realPort = os.Getenv("PORT")
	if realPort == "" {
		realPort = strconv.Itoa(*port)
	}
	err = server.Run(":" + realPort)
	if err != nil {
		log.Println(err)
	}
}
