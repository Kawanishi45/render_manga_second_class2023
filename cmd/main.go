package main

import (
	"database/sql"                                 // データベース操作に関するパッケージ
	"fmt"                                          // フォーマット出力に関するパッケージ
	"github.com/render_manga_api/controller"       // コントローラのパッケージ
	"github.com/render_manga_api/model/repository" // リポジトリのパッケージ
	"log"                                          // ログ出力に関するパッケージ
	"net/http"                                     // HTTP通信に関するパッケージ
	"os"                                           // オペレーティングシステムの機能に関するパッケージ
)

var (
	tr = repository.NewTitleRepository()   // タイトルのリポジトリを作成
	tc = controller.NewTitleController(tr) // タイトルのコントローラを作成

	pr = repository.NewPageRepository()   // ページのリポジトリを作成
	pc = controller.NewPageController(pr) // ページのコントローラを作成

	ro = controller.NewRouter(tc, pc) // タイトルとページのコントローラを指定してルーターを作成
)

func main() {
	var err error

	// データベース接続情報を設定
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DbHost"), os.Getenv("DbPort"), os.Getenv("DbUser"), os.Getenv("DbPassword"), os.Getenv("DbName"))

	repository.Db, err = sql.Open("postgres", psqlInfo) // PostgreSQLデータベースへの接続を開く
	if err != nil {
		log.Println(err)
	}

	defer repository.Db.Close() // main 関数の終了時にデータベース接続を閉じる

	err = repository.Db.Ping() // データベースへの接続を確認する
	if err != nil {
		log.Println(err)
		os.Exit(1) // 接続できない場合はプログラムを終了する
	}

	fmt.Println("Successfully connected to db!")

	server := http.Server{
		Addr: ":8080", // サーバーのアドレスを指定
	}

	http.HandleFunc("/home", ro.HandleTitlesRequest)  // "/home" へのリクエストをタイトルハンドラに処理させる
	http.HandleFunc("/viewer", ro.HandlePagesRequest) // "/viewer" へのリクエストをページハンドラに処理させる

	fmt.Println("ListenAndServe...")

	err = server.ListenAndServe() // サーバーを起動してリクエストを受け付ける
	if err != nil {
		log.Println(err)
	}
}
