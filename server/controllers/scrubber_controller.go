package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/truemail-rb/truemail-go"

func scrubMail(c *gin.Context) {
	mail := c.Param("mail")
	configuration, _ := truemail.NewConfiguration(
		truemail.ConfigurationAttr{
			VerifierEmail: "azerty@gmail.com",
		},
	)
	validate, err := truemail.Validate(mail, configuration)
	if err != nil {
		return
	}
	if validate.Success {
		c.Writer.WriteHeader(http.StatusOK)
	} else {
		fmt.Printf("%#v", validate)
		c.Writer.WriteHeader(http.StatusNotFound)
	}
}
