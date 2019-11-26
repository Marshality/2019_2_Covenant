package usecase

import (
	"2019_2_Covenant/internal/models"
	"2019_2_Covenant/internal/playlist"
)

type PlaylistUsecase struct {
	playlistRepo playlist.Repository
}

func NewPlaylistUsecase(repo playlist.Repository) playlist.Usecase {
	return &PlaylistUsecase{
		playlistRepo: repo,
	}
}

func (pUC *PlaylistUsecase) Store(playlist *models.Playlist) error {
	if err := pUC.playlistRepo.Store(playlist); err != nil {
		return err
	}

	return nil
}

func (pUC *PlaylistUsecase) Fetch(userID uint64, count uint64, offset uint64) ([]*models.Playlist, uint64, error) {
	playlists, total, err := pUC.playlistRepo.Fetch(userID, count, offset)

	if err != nil {
		return nil, total, err
	}

	if playlists == nil {
		playlists = []*models.Playlist{}
	}

	return playlists, total, nil
}

func (pUC *PlaylistUsecase) DeleteByID(playlistID uint64) error {
	if err := pUC.playlistRepo.DeleteByID(playlistID); err != nil {
		return err
	}

	return nil
}
