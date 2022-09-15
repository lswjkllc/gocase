package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SrOnMac struct {
	C0 int    `json:"c0"`
	C1 string `json:"c1"`
	C2 string `json:"c2"`
	C3 string `json:"c3"`
}

var INSERT_DML = `
insert into sr_on_mac values (1, '2022-02-01', '2022-02-01 10:47:57', '111'),(2, '2022-02-02', '2022-02-02 10:47:57', '222'),(3, '2022-02-03', '2022-02-03 10:47:57', '333')
`
var SELECT_DML = `
select c0,c1,c2,c3 from sr_on_mac where c1 >= '2022-02-02'
`

func main() {
	fmt.Println("mysql test ...")

	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db, err := sql.Open("mysql", "root:@tcp(ai.wgine-dev.com:30013)/TEST") // 设置连接数据库的参数
	if err != nil {
		fmt.Println("数据库打开错误:", err)
		return
	}
	defer db.Close() //关闭数据库

	// ping
	err = db.Ping() //连接数据库
	if err != nil {
		fmt.Println("数据库连接失败:", err) //连接失败
		return
	} else {
		fmt.Println("数据库连接成功") //连接成功
	}

	// // 插入数据
	// result, err := db.Exec(INSERT_DML)
	// if err != nil {
	// 	fmt.Println("数据插入失败:", err)
	// 	return
	// } else {
	// 	fmt.Println("数据插入成功:", result)
	// }
	// 查询数据
	row := SrOnMac{}

	qRows, err := db.Query(SELECT_DML)
	if err != nil {
		fmt.Println("数据查询失败:", err)
		return
	}
	for qRows.Next() {
		qRows.Scan(&row.C0, &row.C1, &row.C2, &row.C3)
		fmt.Println(row.C0, row.C1, row.C2, row.C3)
	}

	// qRow := db.QueryRow(SELECT_DML, 1)
	// err = qRow.Scan(&row.C0, &row.C1, &row.C2, &row.C3)
	// if err != nil {
	// 	fmt.Println("数据查询失败:", err)
	// 	return
	// } else {
	// 	fmt.Println("数据查询成功:", row)
	// }
}
