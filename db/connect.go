package db

import (
	"context"
	"database/sql"
	"log"
	"websocket/utils"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *bun.DB

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int64 `bun:",pk,autoincrement"`
	Name string
}


func Connect() {
		// Open a PostgreSQL database.
		dsn := utils.DBURL
		// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		defer sqldb.Close()
		// 使用するデータベースに合わせて、第二引数を変える。
		DB = bun.NewDB(sqldb, pgdialect.New())
		// クエリを標準出力するコードです。
		// 動作が分かりやすいため、入れておくことをオススメします。
		DB.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
		if _, err := DB.NewCreateTable().Model((*User)(nil)).Exec(context.TODO()); err != nil {
			log.Fatalln(err)
		}
		user := &User{Name: "鈴木太郎"}
		if _, err := DB.NewInsert().Model(user).Exec(context.TODO()); err != nil {
			log.Fatalln(err)
		}
}