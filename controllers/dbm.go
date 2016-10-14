package controllers

import (
	. "bulaoge/models"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type DbmController struct {
	beego.Controller
}

var (
	MssqlDbName  = "dentist"
	MysqlDbName  = "dentist"
	pageSize     = 10
	tables       = make([]string, 0)
	currentTable = ""
	tableData    = make([]interface{}, 0)
	columns      = make([]string, 0)
)

func (c *DbmController) Get() {
	db, err := NewMssqlEngine(MssqlDbName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select name from sys.tables")

	if err != nil {
		fmt.Println(err)
		return
	}

	var i int
	var v string
	for rows.Next() {
		rows.Scan(&v)
		tables = append(tables, v)
		i++
	}

	c.Data["tables"] = tables
	c.TplName = "dbm/index.html"
}

func (c *DbmController) List() {

	currentTable = c.Ctx.Input.Param(":tableName")
	sql := "select top 10 * from " + currentTable

	c.Data["columns"] = []string{"*"}

	if strings.EqualFold(c.Ctx.Request.Method, "POST") {
		sql = c.GetString("sql", sql)

		db, err := NewMssqlEngine(MssqlDbName)
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

		columns, err = rows.Columns()
		if err != nil {
			fmt.Println(err)
			return
		}

		c.Data["columns"] = columns

		values := make([]interface{}, len(columns))

		scanArgs := make([]interface{}, len(columns))

		c.Data["attrs"] = strings.Join(columns, ",")
		for i := range columns {
			scanArgs[i] = &values[i]
		}

		tableData = tableData[:0]

		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				fmt.Println(err)
			}

			m := make([]interface{}, 0)
			for j := range values {
				m = append(m, values[j])
			}
			tableData = append(tableData, m)
		}
		c.Data["list"] = tableData
	}

	c.Data["tables"] = tables
	c.Data["currentTable"] = currentTable
	c.Data["sql"] = sql
	c.TplName = "dbm/index.html"
}

func (c *DbmController) Move() {
	var sql = ""

	for i, v := range tableData {

		vs := make([]interface{}, 0)

		vs = append(vs, v)

		b, _ := json.Marshal(vs)
		values := string(b)
		values = strings.Replace(values, "[", "", -1)
		values = strings.Replace(values, "]", "", -1)

		sql = "insert into " + currentTable + "(" + strings.Join(columns, ",") + ") values (" + values + ")"

		func() {
			db, err := NewMysqlEngine(MysqlDbName)

			if err != nil {
				fmt.Println(err)
				return
			}
			defer db.Close()

			_, err = db.Exec(sql)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(i)
			}
			time.Sleep(time.Millisecond * 100) //100ms插入一次
		}()
	}
	c.Ctx.Output.Body([]byte("迁移成功"))
}

//测试sqlserver
func testMssql() {
	db, err := NewMssqlEngine(MssqlDbName)
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
	db, err := NewMysqlEngine(MysqlDbName)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from kc_name")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows.Columns())
}
