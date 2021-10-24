package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)



func getCreateProductHandler(b *broker) func(c *gin.Context) {
	return func(c *gin.Context) {
		var p product

		err := json.NewDecoder(c.Request.Body).Decode(&p)
		if err != nil {
			log.Println("error parsing request body", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("error parsing body: %s", err.Error()),
			})
			return
		}

		p.ID = uuid.New().String()

		body, err := json.Marshal(p)
		if err != nil {
			log.Println("error encoding event", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("error encoding event: %s", err.Error()),
			})
			return
		}

		if err := b.publish(body); err != nil {
			log.Println("error publishing product event")
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("error sending event: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "product created - id(" + p.ID + ")",
		})
	}
}

