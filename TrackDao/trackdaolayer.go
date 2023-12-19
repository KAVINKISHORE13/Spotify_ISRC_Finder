// trackdao.go

package trackdao

import (
	"github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/model"
	"gorm.io/gorm"
)

type TrackDAL struct {
	DB *gorm.DB
}

func NewTrackDAL(db *gorm.DB) *TrackDAL {
	return &TrackDAL{
		DB: db,
	}
}

func (dao *TrackDAL) CreateTrack(track *model.Track) error {

	if err := dao.DB.Create(track).Error; err != nil {
		return err
	}
	return nil
}


func (dao *TrackDAL) GetTrackByISRC(isrc string) (*model.Track, error) {
	var track model.Track
	err := dao.DB.Where("isrc = ?", isrc).First(&track).Error
	if err != nil {
		return nil, err
	}
	return &track, nil
}

func (dao *TrackDAL) GetTracksByArtist(artist string) (*[]model.Track, error) {
	var tracks *[]model.Track
	err := dao.DB.Where("array_to_string(artist_names, '||') ILIKE ?", "%"+artist+"%").Order("popularity DESC").Find(&tracks).Error
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func (dao *TrackDAL) UpdateTrack(track *model.Track) error {

	err := dao.DB.Save(track).Error
	if err != nil {
		return err
	}

	return nil
}