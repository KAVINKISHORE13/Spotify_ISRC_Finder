package routes

import (
	"github.com/gin-gonic/gin"
	trackdao "github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/TrackDao"
	trackservice "github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/TrackService"
	trackhandler "github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/Trackhandler"
	"gorm.io/gorm"
)


func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	trackDAO := trackdao.NewTrackDAL(db)
	trackService := trackservice.NewtrackService(trackDAO, trackservice.NewSpotifyClient("b65d33b1cb2049719f65fb43f1b74dc1", "7f380918e96f4b1a804ac2b37a542097"))
	trackHandler := trackhandler.NewTrackHandler(trackService, trackService.GetSpotifyClient())  // Assuming you have a GetSpotifyClient() method in your TrackService
	

	router.POST("/tracks/create", trackHandler.CreateTrackHandler)

	router.GET("/tracks/:isrc", trackHandler.GetTrackByISRCHandler)

	router.GET("/tracks/artist/:artist", trackHandler.GetTracksByArtistHandler)

	router.PUT("/tracks/update/:isrc", trackHandler.UpdateTrackByISRCHandler)



}
