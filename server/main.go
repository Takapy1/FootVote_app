package main

import (
	// "github.com/gin-gonic/gin"
	"github.com/Takapy1/FootVote/server/routers"

	"log"
)

func main() {
	// database.SetUp()
	router := routers.SetUp()
	router.LoadHTMLFiles("clients/public/index.html")
	if err := router.Run(":3002"); err != nil {
		log.Fatal(err)
	}
}