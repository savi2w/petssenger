package http

import (
	"fmt"
	"net"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/nglabo/petssenger/services/user/config"
	"github.com/nglabo/petssenger/services/user/models"
)

type User struct {
	Email string `json:"email"`
}

type Response struct {
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func createUser(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
			Payload: nil,
		})

		return
	}

	email := "^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:.[a-zA-Z0-9-]+)*$"
	matched, err := regexp.MatchString(email, json.Email)
	if err != nil {
		panic(err)
	}

	if !matched {
		c.JSON(http.StatusBadRequest, Response{
			Message: fmt.Sprintf(`The email "%v" is invalid`, json.Email),
			Payload: nil,
		})

		return
	}

	user, err := models.CreateUser(json.Email)
	if err != nil {
		pgErr, ok := err.(pg.Error)
		if ok && pgErr.IntegrityViolation() {
			c.JSON(http.StatusInternalServerError, Response{
				Message: fmt.Sprintf(`The email "%v" already exists`, json.Email),
				Payload: nil,
			})

			return
		}

		panic(err)
	}

	c.JSON(http.StatusCreated, Response{
		Message: http.StatusText(http.StatusCreated),
		Payload: user,
	})
}

func UserHTTPListen() (net.Listener, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/user", createUser)

	lis, err := net.Listen("tcp", config.Default.HTTPPort)
	if err != nil {
		return nil, err
	}

	if err := r.RunListener(lis); err != nil {
		lis.Close()
		return nil, err
	}

	return lis, nil
}
