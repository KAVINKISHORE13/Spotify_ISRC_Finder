// service.go

package trackservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/TrackDao"
	"github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/model"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)


type SpotifyClient struct {
	ClientID     string
	ClientSecret string
}


type TrackService struct {
	trackDAO *trackdao.TrackDAL
	spotifyClient *SpotifyClient
}

func NewtrackService(trackDAO *trackdao.TrackDAL, spotifyClient *SpotifyClient) *TrackService {
	return &TrackService{trackDAO: trackDAO, spotifyClient: spotifyClient}
}

func NewSpotifyClient(clientID, clientSecret string) *SpotifyClient {
	return &SpotifyClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}
func (s *TrackService) GetSpotifyClient() *SpotifyClient {
	return s.spotifyClient
}


func (s *TrackService) CreateTrack(isrc string) (*model.Track, error) {

	existingTrack, err := s.trackDAO.GetTrackByISRC(isrc)
	if err != nil {
		fmt.Println("error occured")
	}

	if existingTrack != nil {
		return nil,errors.New("track already exit")
	}
	
	trackMetadata, err := s.spotifyClient.GetTrackMetadata(isrc)
	if err != nil {
		return nil, err
	}

	track := &model.Track{
		ISRC:        isrc,
		SpotifyImage: trackMetadata.SpotifyImage,
		Title:       trackMetadata.Title,
		ArtistNames: trackMetadata.ArtistNames,
		Popularity:  trackMetadata.Popularity,
	}

	// Store the track in the database
	err = s.trackDAO.CreateTrack(track)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (sc *SpotifyClient) GetTrackMetadata(isrc string) (*model.Track, error) {
	config := &clientcredentials.Config{
		ClientID:     sc.ClientID,
		ClientSecret: sc.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return nil, err
	}

	client := spotify.Authenticator{}.NewClient(token)

	result, err := client.Search(fmt.Sprintf("isrc:%s", isrc), spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	}


	if len(result.Tracks.Tracks) == 0 {
		return nil, errors.New("track not found on Spotify")
	}


	highestPopularityTrack := result.Tracks.Tracks[0]
	for _, track := range result.Tracks.Tracks {
		if track.Popularity > highestPopularityTrack.Popularity {
			highestPopularityTrack = track
		}
	}

	
	trackMetadata := &model.Track{
		SpotifyImage: highestPopularityTrack.Album.Images[0].URL,
		Title:        highestPopularityTrack.Name,
		ArtistNames:  getArtistNames(highestPopularityTrack.Artists),
		Popularity:   highestPopularityTrack.Popularity,
	}

	return trackMetadata, nil
}

func getArtistNames(artists []spotify.SimpleArtist) []string {
	var names []string
	for _, artist := range artists {
		names = append(names, artist.Name)
	}
	return names
}

func (s *TrackService) GetTrackByISRC(isrc string) (*model.Track, error) {
	track, err := s.trackDAO.GetTrackByISRC(isrc)
	if err == nil {
		return track, nil
	}

	trackMetadata, err := s.spotifyClient.GetTrackMetadata(isrc)
	if err != nil {
		return nil, err
	}

	track = &model.Track{
		ISRC:        isrc,
		SpotifyImage: trackMetadata.SpotifyImage,
		Title:       trackMetadata.Title,
		ArtistNames: trackMetadata.ArtistNames,
		Popularity:  trackMetadata.Popularity,
	}

	err = s.trackDAO.CreateTrack(track)
	if err != nil {
		return nil, err
	}

	return track, nil
}


func (s *TrackService) GetTracksByArtist(artist string) (*[]model.Track, error) {

	tracks, err := s.trackDAO.GetTracksByArtist(artist)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
	

func (s *TrackService) UpdateTrackByISRC(isrc string, updatedTrack *model.Track) (*model.Track, error) {
	// Check if the track exists in the database
	existingTrack, err := s.trackDAO.GetTrackByISRC(isrc)
	if err != nil {
		return nil, errors.New("track not found")
	}

	existingTrack.SpotifyImage = updatedTrack.SpotifyImage
	existingTrack.Title = updatedTrack.Title
	existingTrack.ArtistNames = updatedTrack.ArtistNames
	existingTrack.Popularity = updatedTrack.Popularity

	err = s.trackDAO.UpdateTrack(existingTrack)
	if err != nil {
		return nil, errors.New("failed to update track")
	}

	return existingTrack, nil
}
