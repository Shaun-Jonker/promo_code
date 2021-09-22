package controllers

import (
	"net/http"

	models "promo_code/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindPromos(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var promos []models.Promo

	db.Find(&promos)
	c.JSON(http.StatusOK, gin.H{"data": promos})

}

func CreatePromo(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var input models.CreatePromo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	promo := models.Promo{Name: input.Name, Date_from: input.Date_from, Date_to: input.Date_to, Available: input.Available, Amount: input.Amount, Allocated: input.Allocated}
	db.Create(&promo)
	c.JSON(http.StatusOK, gin.H{"data": promo})
}

func FindPromo(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var promo models.Promo

	if err := db.Where("id = ?", c.Param("id")).First(&promo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": promo})
}

func UpdatePromo(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var promo models.Promo

	if err := db.Where("id = ?", c.Param("id")).First(&promo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdatePromo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&promo).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": promo})
}

func DeletePromo(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var promo models.Promo

	if err := db.Where("id = ?", c.Param("id")).First(&promo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&promo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UsePromo(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var promo models.Promo

	if err := db.Where("id = ?", c.Param("id")).First(&promo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UsePromo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&promo).Where("available >= 1").UpdateColumn("available", gorm.Expr("available - ?", 1))

	c.JSON(http.StatusOK, gin.H{"data": true})
}
