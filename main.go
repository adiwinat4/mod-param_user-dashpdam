package main

import (
	"context"
	"fmt"
	"log"
	"utils/mod-user-dash/config"
	"utils/mod-user-dash/pkg/mysql"
)

type User struct {
	Id        int    `db:"id_user"`
	ParamUser string `db:"param_user"`
}

func main() {
	log.Println("Starting app")

	v := config.ReadConfig("./config/app")
	cfg := config.ParseConfig(v)
	mysqlConn, _ := mysql.ConnectDB(cfg)
	defer mysqlConn.Close()
	ctx := context.Background()
	tx := mysqlConn.MustBegin()
	users := []User{}
	sql := `select id_user, param_user from tb_user`
	if err := tx.SelectContext(ctx, &users, sql); err != nil {
		fmt.Println("Error querying data:", err)
	}
	// fmt.Println(users)
	for _, u := range users {
		temp := u.ParamUser
		temp = temp[:len([]rune(temp))-1]
		temp += `,"IoT": {"status": false,"sub": {"swm": false}}}`
		fmt.Println(temp)
		break
		// sql = `update tb_user set param_user=:param_user where id_user=:id_user`
		// tx.NamedExecContext(ctx, sql, map[string]interface{}{"param_user": temp, "id_user": u.Id})
	}
	tx.Commit()
}
