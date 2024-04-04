package common

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Build Up MySQL Connections
func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:imooc@tcp(127.0.0.1:3306)/imooc?charset=utf8")
	return
}

// GetOne
func GetResultRow(rows *sql.Rows) map[string]string {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	record := make(map[string]string)
	for rows.Next() {
		//Saving row data to the record dictionary
		rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				//fmt.Println(reflect.TypeOf(col))
				record[columns[i]] = string(v)
			}
		}
	}
	return record
}

// GetAll
func GetResultRows(rows *sql.Rows) map[int]map[string]string {

	columns, _ := rows.Columns()

	vals := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))

	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {

		rows.Scan(scans...)
		row := make(map[string]string)
		for k, v := range vals {
			key := columns[k]
			row[key] = string(v)
		}
		result[i] = row
		i++
	}
	return result
}
