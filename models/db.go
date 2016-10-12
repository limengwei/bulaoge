package models

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-adodb"

	_ "github.com/go-sql-driver/mysql"
)

////////mssql/////////
type MssqlConfig struct {
	*sql.DB
	DataSource string
	Database   string
	Windows    bool
	SAUser     string
	SAPassword string
}

func (m *MssqlConfig) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.DataSource)
	conf = append(conf, "Initial Catalog="+m.Database)
	if m.Windows { //windows身份验证
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "user id="+m.SAUser)
	conf = append(conf, "password="+m.SAPassword)

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}

func NewMssqlEngine() (dbb *sql.DB, err error) {
	db := MssqlConfig{
		DataSource: "DESKTOP-OI76G5U",
		Database:   "dentist",
		Windows:    false,
		SAUser:     "sa",
		SAPassword: "sa"}

	err = db.Open()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db.DB, nil
}

////////mysql/////////
type MySqlConfig struct {
	*sql.DB
	Addr   string
	DBName string
	User   string
	Passwd string
}

func (m *MySqlConfig) Open() (err error) {
	var url = m.User + ":" + m.Passwd + "@tcp(" + m.Addr + ")/" + m.DBName + "?charset=utf8"

	m.DB, err = sql.Open("mysql", url)
	if err != nil {
		return err
	}
	return nil
}

func NewMysqlEngine() (dbb *sql.DB, err error) {
	db := MySqlConfig{
		Addr:   "localhost:3306",
		DBName: "dentist",
		User:   "root",
		Passwd: ""}

	err = db.Open()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db.DB, nil
}
