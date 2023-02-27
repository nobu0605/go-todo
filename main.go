package main

import (
	"fmt"
	"go-todo/app/controllers"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	controllers.StartMainServer()
}

func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
 
	// models.CreateTodo()
}
