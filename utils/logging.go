package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//os.O_RDWR|os.O_CREATEでファイルを読み書き、ファイルがなければ新しく開く、追記することを可能にする。
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//io.MultiWriterでログの書き込み先を標準出力先とログファイルに指定している。SetFlagsでlogのフォーマットを指定している。
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// SetFlagsでログのフォーマットを指定している。
	log.SetOutput(multiLogFile)
	//OpenFileでLogFileを読み込んで読み書き、作成、開くができlogfileとして変数作成。
	
}