package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Takapy1/FootVote/server/models"
)

func GetGames(c *gin.Context) {
	games := models.GetGames()
	c.JSON(http.StatusOK, gin.H{ "data": games })
}

func GetGame(c *gin.Context) {
	fmt.Println("Hello getGame") 
	c.JSON(http.StatusOK, gin.H{ "data": "hello getGame" })
}