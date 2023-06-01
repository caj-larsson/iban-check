package main

import (
	"git.sg.caj.me/caj/iban-check/v2/iban"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.Writer.Write([]byte("ok"))
		c.Writer.Flush()
	})

	r.GET("/v1/iban/:iban", func(c *gin.Context) {
		form := struct {
			Iban string `uri:"iban" binding:"required"`
		}{}
		c.ShouldBindUri(&form)

		ibanValue := iban.New(form.Iban)

		ibanError := ibanValue.ValidationError()
		var errorMessage = ""

		if ibanError != nil {
			errorMessage = ibanError.Error()
		}

		c.JSON(200, gin.H{
			"valid": ibanError == nil,
			"error": errorMessage,
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
