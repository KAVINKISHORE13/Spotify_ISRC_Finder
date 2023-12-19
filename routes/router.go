package routes

import (
	// "os"

	"github.com/gin-gonic/gin"
	trackdao "github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/TrackDao"
	trackservice "github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/TrackService"
	trackhandler "github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/Trackhandler"
	"gorm.io/gorm"
)


func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	trackDAO := trackdao.NewTrackDAO(db)
	trackService := trackservice.NewtrackService(trackDAO, trackservice.NewSpotifyClient("cliient_ID","Client_secret" ))
	trackHandler := trackhandler.NewTrackHandler(trackService, trackService.GetSpotifyClient())  
	

	router.POST("/tracks/create", trackHandler.CreateTrackHandler)

	router.GET("/tracks/:isrc", trackHandler.GetTrackByISRCHandler)

	router.GET("/tracks/artist/:artist", trackHandler.GetTracksByArtistHandler)

	router.PUT("/tracks/update/:isrc", trackHandler.UpdateTrackByISRCHandler)
}
