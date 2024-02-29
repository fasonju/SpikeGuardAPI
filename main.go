package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"prototype_app_api/db"
)
import "prototype_app_api/endpoints"

func main() {

	if os.Getenv("env") == "dev" {
		gin.SetMode(gin.DebugMode)
		if godotenv.Load(".env") != nil {
			log.Fatalln("Could not load .env file")
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := setupRouter()

	var err error
	db.DB, err = setupConnection()
	if err != nil {
		log.Fatalln("Could not connect to the database")
	}

	var port = os.Getenv("PORT")

	if err := r.Run(":" + port); err != nil {
		log.Fatalln("Could not start the server")
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/markers", endpoints.MarkersGET)
	r.POST("/markers", endpoints.MarkersPOST)
	r.DELETE("/markers", endpoints.MarkersDelete)
	r.PUT("/markers", endpoints.MarkersPUT)
	r.PUT("/report", endpoints.ReportPUT)
	return r
}

func setupConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT"),
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}

	DB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return DB, nil
}
