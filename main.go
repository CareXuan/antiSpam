package main

import (
	"antispam/base"
	"antispam/http"
	"fmt"
)

type T1 struct {
	Id  int64
	Val int64
}

func main() {
	err := base.Init("./conf/local.yaml")
	if err != nil {
		fmt.Print(err)
	}
	//result, err := src.ContentCheckSecondStep("裸聊", "759599", "PORN_AD")
	//re, err := videoSpam.VideoContentResult("7595991314")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Print(re)
	r := http.InitGin()
	r.Run(":1234")
}

func mysqlTest() {
	//err := base.Init("./conf/local.yaml")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//mysqlConn := conf.Mysql
	//defer mysqlConn.Close()
	//sql := "select * from t1"
	//rows, err := mysqlConn.Query(sql)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//var T1s []T1
	//for rows.Next() {
	//	var id int64
	//	var val int64
	//	var t1 T1
	//	rows.Scan(&id, &val)
	//	t1.Id = id
	//	t1.Val = val
	//	T1s = append(T1s, t1)
	//}
	//fmt.Print(T1s)
}
