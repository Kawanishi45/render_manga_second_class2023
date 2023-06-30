package dto

type PagesRequest struct {
	TitleId int `json:"title_id"` // ページリクエストのタイトルID
}

type PagesResponse struct {
	ImageUrls []string `json:"page_urls"` // ページレスポンスの画像URLのリスト
}

/*
このコードは、データ転送オブジェクト（DTO）の定義です。
PagesRequestはページリクエストを表し、TitleIdフィールドはリクエストされるタイトルのIDを示します。
PagesResponseはページレスポンスを表し、ImageUrlsフィールドはページの画像のURLのリストを格納します。
これらのDTOはJSON形式でのデータの受け渡しに使用されることを想定しています。
*/
