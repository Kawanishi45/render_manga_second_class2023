package controller

import (
	"encoding/json"                                // JSONのエンコード・デコードに関するパッケージ
	"github.com/render_manga_api/controller/dto"   // コントローラのDTOのパッケージ
	"github.com/render_manga_api/model/repository" // リポジトリのパッケージ
	"log"                                          // ログ出力に関するパッケージ
	"net/http"                                     // HTTP通信に関するパッケージ
	"strconv"                                      // 文字列と数値の変換に関するパッケージ
)

type PageController interface {
	GetPages(w http.ResponseWriter, r *http.Request)
}

type pageController struct {
	pr repository.PageRepository // ページのリポジトリ
}

func NewPageController(pr repository.PageRepository) PageController {
	return &pageController{pr} // 新しいページコントローラを作成して返す
}

func (c pageController) GetPages(w http.ResponseWriter, r *http.Request) {
	titleId, err := strconv.Atoi(r.URL.Query().Get("title_id")) // URLのクエリパラメータからタイトルIDを取得
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // エラーレスポンスを返す
		log.Println(err)                     // エラーログを出力
		return
	}

	pageImageUrls, err := c.pr.GetPageImageUrlsByTitleId(titleId) // タイトルIDに基づいてページの画像URLを取得
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // エラーレスポンスを返す
		log.Println(err)                              // エラーログを出力
		return
	}

	var pagesResponse dto.PagesResponse
	pagesResponse = dto.PagesResponse{ImageUrls: pageImageUrls} // ページレスポンスの作成

	output, err := json.MarshalIndent(pagesResponse, "", "\t\t") // ページレスポンスをJSON形式にエンコード
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // エラーレスポンスを返す
		log.Println(err)                              // エラーログを出力
		return
	}

	// レスポンスヘッダーの設定
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
