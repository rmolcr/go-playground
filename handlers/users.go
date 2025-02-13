package handlers

import (
	"net/http"

	"com.rmolcr/go-playground/domain"
	"github.com/gin-gonic/gin"
)

var users = []domain.User{
	{ID: "1", Name: "Blue Train", Surname: "John Coltrane", Email: "btjc@myhost.com", Number: 666000222},
	{ID: "2", Name: "White Train", Surname: "John Doe", Email: "wtjd@myhost.com", Number: 666111333},
	{ID: "3", Name: "Yellow Train", Surname: "Alice Cooper", Email: "ytac@myhost.com", Number: 666222444},
	{ID: "4", Name: "Red Train", Surname: "Bob Marley", Email: "rtbm@myhost.com", Number: 666333555},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
