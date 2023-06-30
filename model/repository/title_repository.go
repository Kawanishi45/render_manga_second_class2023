package repository

import (
	"github.com/render_manga_api/model/entity" // モデルのエンティティのパッケージ
	"log"                                      // ログ出力に関するパッケージ
)

const (
	PickUp = "pick_up" // ピックアップタイトルの種類を表す定数
	Normal = "normal"  // 通常タイトルの種類を表す定数
)

type TitleRepository interface {
	GetTitles() (todos []entity.TitleEntity, err error)
}

type titleRepository struct {
}

func NewTitleRepository() TitleRepository {
	return &titleRepository{} // 新しいタイトルリポジトリを作成して返す
}

func (t *titleRepository) GetTitles() (titles []entity.TitleEntity, err error) {
	titles = []entity.TitleEntity{} // タイトルのエンティティのスライスを初期化
	rows, err := Db.
		Query(`
SELECT id, name, type, thumbnail_url
FROM titles ORDER BY id DESC
`) // タイトルの情報を取得するクエリを実行

	if err != nil {
		log.Print(err) // エラーログを出力
		return
	}

	for rows.Next() {
		title := entity.TitleEntity{} // タイトルのエンティティを初期化
		err = rows.Scan(&title.Id, &title.Name, &title.Type, &title.ThumbnailUrl)
		if err != nil {
			log.Print(err) // エラーログを出力
			return
		}
		titles = append(titles, title) // タイトルのエンティティをスライスに追加
	}

	return
}
