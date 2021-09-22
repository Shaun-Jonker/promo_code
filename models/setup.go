package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")

	viper.AutomaticEnv()

	viper_user := viper.Get("POSTGRES_USER")
	viper_password := viper.Get("POSTGRES_PASSWORD")
	viper_db := viper.Get("POSTGRES_DB")
	viper_host := viper.Get("POSTGRES_HOST")
	viper_port := viper.Get("POSTGRES_PORT")

	fmt.Print(viper_user)

	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)

	fmt.Println("conname is:", prosgret_conname)

	db, err := gorm.Open("postgres", prosgret_conname)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Promo{})

	m := Promo{Name: "Herbalife", Date_from: "20/09/2021", Date_to: "30/09/2021", Available: 20, Amount: 250.00, Allocated: 30}
	db.Create(&m)
	return db

}
