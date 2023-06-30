package controller

import (
	"encoding/json"                                // JSONのエンコード・デコードに関するパッケージ
	"github.com/render_manga_api/controller/dto"   // コントローラのDTOのパッケージ
	"github.com/render_manga_api/model/repository" // リポジトリのパッケージ
	"log"                                          // ログ出力に関するパッケージ
	"net/http"                                     // HTTP通信に関するパッケージ
)

type TitleController interface {
	GetTitles(w http.ResponseWriter, r *http.Request)
}

type titleController struct {
	tr repository.TitleRepository // タイトルのリポジトリ
}

func NewTitleController(tr repository.TitleRepository) TitleController {
	return &titleController{tr} // 新しいタイトルコントローラを作成して返す
}

func (tc *titleController) GetTitles(w http.ResponseWriter, r *http.Request) {
	titles, err := tc.tr.GetTitles() // タイトルの取得を行う
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // エラーレスポンスを返す
		log.Println(err)                              // エラーログを出力
		return
	}

	var pickUpTitleResponses []dto.TitleResponse
	var normalTitleResponses []dto.TitleResponse
	for _, v := range titles {
		if v.Type == repository.PickUp {
			pickUpTitleResponses = append(pickUpTitleResponses, dto.TitleResponse{TitleId: v.Id, Name: v.Name, ThumbnailUrl: v.ThumbnailUrl}) // ピックアップタイトルのレスポンスを追加
		} else {
			normalTitleResponses = append(normalTitleResponses, dto.TitleResponse{TitleId: v.Id, Name: v.Name, ThumbnailUrl: v.ThumbnailUrl}) // 通常タイトルのレスポンスを追加
		}
	}

	var titlesResponse dto.TitlesResponse
	titlesResponse.PickupTitles = pickUpTitleResponses
	titlesResponse.NormalTitles = normalTitleResponses

	output, _ := json.MarshalIndent(titlesResponse, "", "\t\t") // タイトルレスポンスをJSON形式にエンコード

	w.Header().Set("Content-Type", "application/json")                                                                                   // レスポンスのコンテンツタイプを設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                   // CORSを許可するオリジンを設定
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                                                    // CORSを許可するHTTPメソッドを設定
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization") // CORSを許可するヘッダーを設定

	_, err = w.Write(output) // レスポンスの書き込み
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // エラーレスポンスを返す
		log.Println(err)                              // エラーログを出力
		return
	}
}
