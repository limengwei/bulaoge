package controllers

import (
	. "bulaoge/models"
	"fmt"

	"github.com/astaxie/beego"
)

type DbmController struct {
	beego.Controller
}

var (
	tables = []string{"kc_name", "wh_item_case"}
)

func (c *DbmController) Get() {
	c.Data["tables"] = tables
	c.TplName = "dbm/index.html"
}

func (c *DbmController) List() {
	fmt.Println("listttttttttttt")
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
