package main

import (
	"fmt"
	"sqlController/goMySql"
	"time"
)

func main() {
	// データの取得（GET）
	data := goMySql.GetData()
	fmt.Println(data)
	time.Sleep(1 * time.Second)

	// データの挿入（POST）
	goMySql.PostData("Guest", 1)
	data = goMySql.GetData()
	fmt.Println(data)
	time.Sleep(1 * time.Second)

	// データの更新（PUT）
	update := data[0]
	update.Name = "Administrator"
	update.Age = 99
	goMySql.PutData(update)
	data = goMySql.GetData()
	fmt.Println(data)
	time.Sleep(1 * time.Second)

	// データの削除
	goMySql.DeleteData(1)
	data = goMySql.GetData()
	fmt.Println(data)
	time.Sleep(1 * time.Second)
}
