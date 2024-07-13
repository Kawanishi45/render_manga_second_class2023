package controller

import "net/http"

type Router interface {
	HandleTitlesRequest(w http.ResponseWriter, r *http.Request)
	HandlePagesRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc TitleController
	pc PageController
}

func NewRouter(tc TitleController, cc PageController) Router {
	return &router{tc, cc} // 新しいルーターを作成して返す
}

func (ro *router) HandleTitlesRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.tc.GetTitles(w, r) // タイトルリクエストを処理する
	default:
		w.WriteHeader(http.StatusMethodNotAllowed) // サポートされていないHTTPメソッドへのレスポンスを返す
	}
}

func (ro *router) HandlePagesRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.pc.GetPages(w, r) // ページリクエストを処理する
	default:
		w.WriteHeader(http.StatusMethodNotAllowed) // サポートされていないHTTPメソッドへのレスポンスを返す
	}
}
