package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.AlbumGet, error)
	GetAlbumService(ctx context.Context, singerID model.AlbumID) (*model.AlbumGet, error)
	PostAlbumService(ctx context.Context, singer *model.Album) error
	DeleteAlbumService(ctx context.Context, singerID model.AlbumID) error
}

type albumService struct {
	albumRepository repository.AlbumRepository
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(albumRepository repository.AlbumRepository) *albumService {
	return &albumService{albumRepository: albumRepository}
}

func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.AlbumGet, error) {
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumGet, error) {
	singer, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}
	return singer, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
	if err := album.Validate(); err != nil {
		return err
	}

	if err := s.albumRepository.Add(ctx, album); err != nil {
		return err
	}
	return nil
}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumRepository.Delete(ctx, albumID); err != nil {
		return err
	}
	return nil
}