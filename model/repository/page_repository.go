package repository

import (
	"log" // ログ出力に関するパッケージ
)

type PageRepository interface {
	GetPageImageUrlsByTitleId(titleId int) (pageImages []string, err error)
}

type pageRepository struct {
}

func NewPageRepository() PageRepository {
	return &pageRepository{} // 新しいページリポジトリを作成して返す
}

func (c pageRepository) GetPageImageUrlsByTitleId(titleId int) (pageImageUrls []string, err error) {
	rows, err := Db.
		Query(`
SELECT page_image_url
FROM pages
WHERE title_id = $1
ORDER BY id
`, titleId) // タイトルIDに基づいてページの画像URLを取得するクエリを実行

	if err != nil {
		log.Print(err) // エラーログを出力
		return
	}

	for rows.Next() {
		var pageImageUrl string
		err = rows.Scan(&pageImageUrl)
		if err != nil {
			log.Print(err) // エラーログを出力
			return
		}
		pageImageUrls = append(pageImageUrls, pageImageUrl)
	}

	return
}
