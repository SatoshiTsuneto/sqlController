### Go言語 version: 1.12.4

### 実行環境: Ubuntu19.04

### このアプリケーションについて
- このアプリケーションは、Go言語でデータベースを制御するためサンプルとなっています。   
基本的には、データの取得、データの作成、データの更新、データの削除ができます。   

- 使い方は、ソースコードを見ればわかるようになっています。

- このサンプルではMySQLを使用していますが、他のものに変更することも可能です。   
変更方法は、[こちら](https://golang.org/pkg/database/sql/)を参考にしてください。

### セットアップ
- `go get "github.com/go-sql-driver/mysql"` 

### テスト環境の作成および実行例
- Dockerを用いてSQLサーバを作成します。(Dockerを使用しなくても、以下と同じ環境を作成できれば大丈夫です)   
※Dockerのインストール方法については、[こちら](https://qiita.com/n-yamanaka/items/ddb18943f5e43ca5ac2e)を参考にしてください。   

　　以下のコマンドを実行して作成します。   
　　`$ docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=secret -d -p 3306:3306 mysql`   

　　次に、コンテナに入りデータベースおよびテーブルの作成をします。以下のコマンドを順に実行してください。   
```
$ docker exec -it test-mysql bash
$ mysql -u root -p secret
$ create database user_info;
$ create table posts(id int, name varchar(255), age int);
```

- ホストOSに戻り、以下のコマンドで、必要なライブラリをインストールしてください。   
`go get "github.com/go-sql-driver/mysql"` 

- 任意のディレクトリにこのリポジトリをクローンしてください。

- 以下のコマンドでテストプログラムを実行して、エラーが出なければ成功です。   
`$ go run main.go`
