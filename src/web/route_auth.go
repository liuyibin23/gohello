package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"model"
	"net/http"
)

//POST /auth
func auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		//fmt.Fprintf(w,string(body))
		var authRequest model.AuthRequest
		err := json.Unmarshal(body, &authRequest)
		authResponse := model.AuthResponse{}
		if err == nil {
			db, err := sql.Open("sqlite3", "./appDB.db")
			checkErr(err)

			rows, err := db.Query("SELECT username,password FROM users WHERE username=?", authRequest.UserName)
			checkErr(err)
			if rows.Next() {
				var uname string
				var pwd string
				rows.Scan(&uname, &pwd)
				if pwd == authRequest.PassWord {
					authResponse.Status = "0"
					authResponse.DeveleperPermissionDevices = ""
					authResponse.RemotePermissionDevices = ""
				} else {
					authResponse.Status = "2"
					authResponse.DeveleperPermissionDevices = ""
					authResponse.RemotePermissionDevices = ""
				}
			} else {
				authResponse.Status = "1"
				authResponse.DeveleperPermissionDevices = ""
				authResponse.RemotePermissionDevices = ""
			}

			rows.Close()
			db.Close()

			////插入数据
			//stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
			//checkErr(err)
			//
			//res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
			//checkErr(err)
			//id, err := res.LastInsertId()
			//checkErr(err)
			//fmt.Println(id)

		} else {
			authResponse.Status = "256"
			authResponse.DeveleperPermissionDevices = ""
			authResponse.RemotePermissionDevices = ""
		}
		responseJson, _ := json.Marshal(authResponse)
		fmt.Fprintf(w, string(responseJson))
	}

}
