package config

import (
	"log"
	"go-todo/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}
//↑これらの設定のためにconfig.iniファイルを作成。

var Config ConfigList
//構造体ConfigListを外部で使えるように変数宣言しておく。

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
	//↑utils/logging.goファイルのLoggingSettingsを呼び出して引数にConfigのLogfileを渡す。
}
// LoadConfig関数をmain関数より先に使用するためにinit関数で呼び出す。
// LoadConfigが実行される。

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	//⇧config.iniファイルを指定する。
	if err != nil {
		log.Fatalln(err)
	}
	// エラーハンドリング

	Config = ConfigList {
		// グローバルに宣言したConfigを使う。
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
		Static: cfg.Section("web").Key("static").String(),
		// 読み込んだiniファイルの値をConfigListに設定する。
	}
}