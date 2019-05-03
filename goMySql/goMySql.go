package goMySql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int
	Name string
	Age  int
}

// データベースに接続する関数
func connectingDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:secret@/user_info")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

// テーブルの行を取得する関数
func getRow(db *sql.DB) (count int, err error) {
	rows, err := db.Query("SELECT COUNT(id) FROM posts")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		_ = rows.Scan(&count)
	}
	return
}

// テーブルからデータの取得
func GetData() (data []Post) {
	// データベースに接続
	db, err := connectingDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// データの取得
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	// 取得したデータの格納
	var line Post
	for rows.Next() {
		_ = rows.Scan(&line.Id, &line.Name, &line.Age)
		data = append(data, line)
	}

	return
}

// テーブルにデータを挿入する関数
func PostData(name string, age int) {
	// データベースに接続
	db, err := connectingDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// 行の取得（IDに使う）
	row, err := getRow(db)
	if err != nil {
		return
	}

	// 準備
	statement := "INSERT INTO posts VALUES (?, ?, ?)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()

	// データの挿入
	_, err = stmt.Exec(row+1, name, age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

// テーブルのデータの更新
func PutData(post Post) {
	// データベースに接続
	db, err := connectingDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// 準備
	statement := "UPDATE posts SET name = ?, age = ? WHERE id = ?"
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()

	// データの更新
	_, err = stmt.Exec(post.Name, post.Age, post.Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

// テーブルからデータを削除
func DeleteData(id int) {
	// データベースに接続
	db, err := connectingDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// 準備
	statement := "DELETE FROM posts WHERE id = ?"
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()

	// データの更新
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
