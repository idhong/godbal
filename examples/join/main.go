package main

import (
	"encoding/json"
	"fmt"
	"github.com/xujiajun/godbal"
	"github.com/xujiajun/godbal/driver/mysql"
)

func main() {
	connection, _ := godbal.NewMysql("root:123@tcp(127.0.0.1:3306)/test?charset=utf8").Open()

	//queryBuilder := mysql.NewQueryBuilder(connection)
	//rows, _ := queryBuilder.Select("u.uid,u.username,p.address").From("userinfo", "u").SetFirstResult(0).
	//	SetMaxResults(3).InnerJoin("profile", "p", "u.uid = p.uid").Query()
	//
	//fmt.Print(ToJson(rows))

	queryBuilder2 := mysql.NewQueryBuilder(connection)
	queryBuilder2 = queryBuilder2.Select("u.uid,u.username,p.address").From("userinfo", "u").SetFirstResult(0).
		SetMaxResults(3).RightJoin("profile", "p", "u.uid = p.uid")

	//fmt.Print(queryBuilder2.GetSQL())

	rows, _ := queryBuilder2.Query()
	fmt.Print(ToJson(rows))
}

func ToJson(rows map[int]map[string]string) string {
	jsonString, _ := json.Marshal(&rows)
	return string(jsonString)
}