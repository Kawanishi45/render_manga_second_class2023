package repository

import (
	"database/sql"        // データベースの操作に関するパッケージ
	_ "github.com/lib/pq" // PostgreSQLドライバのパッケージ（ブランク識別子を使用してimportし、ドライバを初期化）
)

var Db *sql.DB // データベース接続のためのグローバル変数
