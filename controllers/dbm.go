package controllers

import (
	. "bulaoge/models"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

type DbmController struct {
	beego.Controller
}

var (
	tables   = []string{"kc_name", "wh_item_case"}
	pageSize = 10
)

func (c *DbmController) Get() {
	c.Data["tables"] = tables
	c.TplName = "dbm/index.html"
}

func (c *DbmController) List() {

	tableName := c.Ctx.Input.Param(":tableName")
	sql := "select top 10 id from " + tableName + " where id not in(select top 0 id from  " + tableName + " order by id desc) order by id desc"

	if strings.EqualFold(c.Ctx.Request.Method, "POST") {
		sql = c.GetString("sql", sql)

		db, err := NewMssqlEngine()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		rows, err := db.Query(sql)

		if err != nil {
			fmt.Println(err)
			return
		}

		columns, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
			return
		}

		c.Data["columns"] = columns

		tableData := make([]interface{}, 0)

		values := make([]interface{}, len(columns))

		scanArgs := make([]interface{}, len(columns))

		for i := range columns {
			scanArgs[i] = &values[i]
		}

		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				fmt.Println(err)
			}

			m := make(map[string]interface{})
			for j := range values {
				m[columns[j]] = values[j]
			}
			tableData = append(tableData, m)
		}
		c.Data["list"] = tableData
	}

	c.Data["tables"] = tables
	c.Data["sql"] = sql
	c.TplName = "dbm/index.html"

}

//测试sqlserver
func testMssql() {
	db, err := NewMssqlEngine()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select id,name,py from kc_name where id<1010")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		drug := Drug{}
		err = rows.Scan(&drug.Id, &drug.Name, &drug.Py)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ID: %d \t 名字: %s \t 拼音: %s\n", drug.Id, drug.Name, drug.Py)
	}
}

//测试mysql
func testMysql() {
	db, err := NewMysqlEngine()

	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query("select * from kc_name")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows.Columns())
}
