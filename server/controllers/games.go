package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	fmt.Println("Hello getGames") //terminal output
	c.JSON(http.StatusOK, gin.H{ "data": "hello getGames" })
}

func GetGame(c *gin.Context) {
	fmt.Println("Hello getGame") 
	c.JSON(http.StatusOK, gin.H{ "data": "hello getGame" })
}