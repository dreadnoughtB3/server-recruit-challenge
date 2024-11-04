package model

type AlbumID int

type Album struct {
	ID       AlbumID  `json:"id"`
	Title    string   `json:"title"`
	SingerID SingerID `json:"singer_id"` // モデル Singer の ID と紐づきます
}

// 取得時用の型定義
type AlbumGet struct {
	ID       AlbumID  `json:"id"`
	Title    string   `json:"title"`
	Singer   *Singer  `json:"singer"`
}

func (a *Album) Validate() error {
	if a.Title == "" {
		return ErrInvalidParam
	}
	if len(a.Title) > 255 {
		return ErrInvalidParam
	}
	return nil
}
