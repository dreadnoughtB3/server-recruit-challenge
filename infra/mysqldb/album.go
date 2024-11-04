package mysqldb

import (
	"context"
	"database/sql"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

func NewAlbumRepository(db *sql.DB) *albumRepository {
	return &albumRepository{
		db: db,
	}
}

type albumRepository struct {
	db *sql.DB
}

// GetAll implements repository.AlbumRepository.
func (a *albumRepository) GetAll(ctx context.Context) ([]*model.AlbumGet, error) {
	albums := []*model.AlbumGet{}
	query := `
		SELECT
			albums.id,
			albums.title,
			singers.id,
			singers.name
		FROM
			albums
		INNER JOIN
			singers
		ON
			albums.singer_id = singers.id
		ORDER BY
			albums.id ASC
	`
	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		album := &model.AlbumGet{}
		singer := &model.Singer{}
		if err := rows.Scan(&album.ID, &album.Title, &singer.ID, &singer.Name); err != nil {
			return nil, err
		}
		if album.ID != 0 {
			album.Singer = singer
			albums = append(albums, album)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

// Get implements repository.AlbumRepository.
func (a *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.AlbumGet, error) {
	album := &model.AlbumGet{}
	singer := &model.Singer{}

	query := `
		SELECT
			albums.id,
			albums.title,
			singers.id,
			singers.name
		FROM
			albums
		INNER JOIN
			singers
		ON
			albums.singer_id = singers.id
		WHERE
			albums.id = ?
		ORDER BY
			albums.id ASC
	`
	rows, err := a.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&album.ID, &album.Title, &singer.ID, &singer.Name); err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if album.ID == 0 {
		return nil, model.ErrNotFound
	}
	album.Singer = singer
	return album, nil
}

// Add implements repository.AlbumRepository.
func (a *albumRepository) Add(ctx context.Context, album *model.Album) error {
	query := "INSERT INTO albums (id, title, singer_id) VALUES (?, ?, ?)"
	if _, err := a.db.ExecContext(ctx, query, album.ID, album.Title, album.SingerID); err != nil {
		return err
	}
	return nil
}

// Delete implements repository.AlbumRepository.
func (a *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
	query := "DELETE FROM albums WHERE id = ?"
	if _, err := a.db.ExecContext(ctx, query, id); err != nil {
		return err
	}
	return nil
}
