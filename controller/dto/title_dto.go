package dto

type TitleRequest struct {
	Name string `json:"name"` // タイトルリクエストの名前
}

type TitleResponse struct {
	TitleId      int    `json:"title_id"`  // タイトルレスポンスのタイトルID
	Name         string `json:"name"`      // タイトルレスポンスのタイトル名
	ThumbnailUrl string `json:"image_url"` // タイトルレスポンスのサムネイル画像のURL
}

type TitlesResponse struct {
	PickupTitles []TitleResponse `json:"pick_up_titles"` // タイトルレスポンスのピックアップタイトルのリスト
	NormalTitles []TitleResponse `json:"normal_titles"`  // タイトルレスポンスの通常タイトルのリスト
}

/*
このコードは、データ転送オブジェクト（DTO）の定義です。
TitleRequestはタイトルリクエストを表し、Nameフィールドはリクエストされるタイトルの名前を示します。
TitleResponseはタイトルレスポンスを表し、TitleIdフィールドはタイトルのID、Nameフィールドはタイトルの名前、ThumbnailUrlフィールドはサムネイル画像のURLを格納します。
TitlesResponseはタイトルリストのレスポンスを表し、PickupTitlesフィールドはピックアップタイトルのリスト、NormalTitlesフィールドは通常のタイトルのリストを格納します。
これらのDTOはJSON形式でのデータの受け渡しに使用されることを想定しています。
*/
