package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("booted")

	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var user UnregisteredUser

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("Registering", user.Username)

		pass, err := registerUser(user)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			fmt.Println("returning")
			return
		}

		fmt.Println("function ended")
		c.JSON(200, gin.H{
			"Message": strings.Join([]string{"Your secret key is", pass, "note it down and do not lose it, otherwise you will lose access to further account actions (your previous mail will be accessible on your local device still)"}, " "),
		})
	})

	r.POST("/sendmail", func(c *gin.Context) {
		type SendRequest struct {
			Mail       Mail   `json:"mail"`
			SenderName string `json:"sender"`
			SecretKey  string `json:"secretKey"`
			Assignee   string `json:"assignee"`
		}

		var request SendRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if taken, err := isUsernameTaken(request.SenderName); taken == false || err != nil {
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			} else {
				fmt.Println(taken)
				c.JSON(400, gin.H{
					"error": "sender name does not exist",
				})
			}
			return
		}

		if validKey, err := verifySecretKey(request.SenderName, request.SecretKey); validKey == false || err != nil {
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(400, gin.H{
					"error": "invalid secret key",
				})
			}
			return
		}

		if err := registerMailForUser(request.Mail, request.Assignee); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			fmt.Println("returning")
			return
		}

		fmt.Println("function ended")
		c.JSON(200, request.Mail)
	})

	r.Handle("QUERY", "/mail", func(c *gin.Context) {
		type SendRequest struct {
			Username  string `json:"username"`
			SecretKey string `json:"secretKey"`
		}

		var request SendRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if taken, err := isUsernameTaken(request.Username); taken == false || err != nil {
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			} else {
				fmt.Println(taken)
				c.JSON(400, gin.H{
					"error": "sender name does not exist",
				})
			}
			return
		}

		if validKey, err := verifySecretKey(request.Username, request.SecretKey); validKey == false || err != nil {
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(400, gin.H{
					"error": "invalid secret key",
				})
			}
			return
		}

		mail, err := readMailForUser(request.Username)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(200, mail)
	})

	r.Run(":2019")
}
