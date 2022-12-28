package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/truemail-rb/truemail-go"

func scrubMail(c *gin.Context) {
	mail := c.Param("mail")
	configuration, _ := truemail.NewConfiguration(
		truemail.ConfigurationAttr{
			VerifierEmail:      "pooetitu@gmail.com",
			NotRfcMxLookupFlow: true,
			SmtpFailFast:       true,
		},
	)
	if truemail.IsValid(mail, configuration) {
		c.Writer.WriteHeader(http.StatusOK)
	} else {
		c.Writer.WriteHeader(http.StatusNotFound)
	}
}
