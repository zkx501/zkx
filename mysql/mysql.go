package main

import (
	"database/sql"
	"fmt"
	//需要在本地配置gobin，并且在gitbub上搞到驱动，并且本地编译通过，只要配置好
	//path,cmd下执行命令：go get github.com/go-sql-driver/mysql
	//就可以再你配置的gobin下看到打包好的可以使用的代码
	//项目主页 https://github.com/Go-SQL-Driver/MySQL ，里面的文档讲解的非常详细
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 打开数据库，sns是我的数据库名字，需要替换你自己的名字，（官网给的没有加tcp，跑不起来，具体有时 间看看源码分析下为何）
	db, err := sql.Open("mysql", "yagou_yggx_user:**v7D7997RQR53F495uKCq76Znp99S2C@tcp(115.236.61.51:33306)/yagou_yggx?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// topic是我本地数据库的表名，需要替换你自己的表名，这里面的英文注释都是引用github官网的~~
	//  嘿嘿 我只是想跑起来看看
	rows, err := db.Query("SELECT * FROM yggx_bianmin_order")
	if err != nil {
		panic(err.Error())
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}
