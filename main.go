package main

import (
	"promo_code/models"

	controllers "promo_code/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/promos", controllers.FindPromos)
	r.POST("/promos", controllers.CreatePromo)
	r.GET("/promos/:id", controllers.FindPromo)
	r.PATCH("/promos/:id", controllers.UpdatePromo)
	r.DELETE("/promos/:id", controllers.DeletePromo)
	r.POST("/promos/:id", controllers.UsePromo)

	r.Run()
}
