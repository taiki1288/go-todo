package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"go-todo/config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB 
// 変数でDbをポインタのsql.DBでグローバルに宣言
 
var err error 
// エラーもグローバルに宣言

const (
	tableNameUser = "users"
	// usersテーブルの宣言
	tableNameTodo = "todos"
	// todosテーブルの宣言
	tableNameSession = "sessions"
	// sessionsテーブルの宣言
)
// テーブル名の宣言

func init() {
	// main関数より先に作成する必要があるのでinit関数で宣言する。
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
		// cmdUを作成。（userstableを作成するもの。
	Db.Exec(cmdU)
	// ↑cmdを実行

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)
		// cmdTを作成。（Todotableを作成するもの。）
	Db.Exec(cmdT)
	// ↑cmdを実行

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	// uuidを作成している。
	return uuidobj
	// rerurnで返している
}
// UUIDを作成する関数。返り値をuuidのUUIDとしている。（uuidパッケージのUUIDを使用している。）

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	// sha1でハッシュ値にする
	return cryptext
}
// passwordの保存はハッシュ値にする必要がある。