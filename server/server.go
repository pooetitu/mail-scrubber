package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mail-scrubber/server/controllers"
	"os"
)

func StartServer() {
	engine := gin.Default()

	controllers.SourceControllers(engine)
	port := "8080"
	if value, isPresent := os.LookupEnv("PORT"); isPresent {
		port = value
	}
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("error %s", err)
	}
}
