package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

//POST /resetpwd
func resetpwd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		var requestJsonMap map[string]string
		err := json.Unmarshal(body, &requestJsonMap)
		uname := requestJsonMap["username"]
		pwd := requestJsonMap["password"]
		if err == nil {
			db, _ := sql.Open("sqlite3", "./appDB.db")
			defer db.Close()
			//db.Query("SELECT COUNT(username) FROM users WHERE username = ?",uname)
			row := db.QueryRow("SELECT COUNT(username) FROM users WHERE username = ?", uname)
			var count int
			row.Scan(&count)
			if count != 0 {
				stmt, _ := db.Prepare("UPDATE users SET password=? WHERE username=?")
				res, _ := stmt.Exec(pwd, uname)
				rowCnt, _ := res.RowsAffected()
				if rowCnt != 0 {

				}
			}
		}

	}
}
