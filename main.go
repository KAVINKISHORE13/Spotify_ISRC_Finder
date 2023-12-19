package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/model"
	"github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/routes"
	docs "github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Spotify ISRC Finder
// @version 1.0
// @description This is a RESTful API for managing music tracks using Gin, Gorm, and PostgreSQL.
// @contact.name kavinkishore
// @contact.url iamkavin1309@gmqil.com
// @externalDocs.url https://github.com/KAVINKISHORE13/Spotify_ISRC_Finder/docs
// @externalDocs.description Documentation for this project can be found on GitHub.
// @host localhost:8080
// @BasePath /

func main() {
	
// DB creation and migration
	db,err:= gorm.Open(postgres.Open("postgres://postgres:K@vin1309@localhost:5432/postgres"), &gorm.Config{});
	if err != nil {
		panic(err)
	}
	if db != nil {
		fmt.Println("database created succesfully")
	}
	db.AutoMigrate(&model.Track{})
	docs.SwaggerInfo.BasePath = "/"


	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	routes.SetupRoutes(router, db)
	router.Run(":8080")
}