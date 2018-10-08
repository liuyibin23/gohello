package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"model"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello!")
}

func log(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w,r)
	}
}

type WorldHandler struct {}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"World!")
}

func auth(w http.ResponseWriter,r *http.Request){
	if r.Method == "POST" {
		len := r.ContentLength
		body := make([]byte,len)
		r.Body.Read(body)
		//fmt.Fprintf(w,string(body))
		var authRequest model.AuthRequest
		err := json.Unmarshal(body,&authRequest)
		authResponse := model.AuthResponse{}
		if err == nil{
			db,err := sql.Open("sqlite3","./appDB.db")
			checkErr(err)

			rows,err := db.Query("SELECT username,password FROM users WHERE username=?",authRequest.UserName)
			checkErr(err)
			if rows.Next(){
				var uname string
				var pwd string
				rows.Scan(&uname,&pwd)
				if pwd == authRequest.PassWord{
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
			////插入数据
			//stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
			//checkErr(err)
			//
			//res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
			//checkErr(err)
			//id, err := res.LastInsertId()
			//checkErr(err)
			//fmt.Println(id)


		} else{
			authResponse.Status = "256"
			authResponse.DeveleperPermissionDevices = ""
			authResponse.RemotePermissionDevices = ""
		}
		responseJson,_ := json.Marshal(authResponse)
		fmt.Fprintf(w,string(responseJson))
	}

}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

func main(){
	world := WorldHandler{}
	server := http.Server{
		Addr:"0.0.0.0:8080",
	}
	//http.HandleFunc ("/hello", hello)
	http.HandleFunc("/hello",log(hello))
	http.Handle("/world", &world)
	http.HandleFunc("/auth",auth)
	server.ListenAndServe()
}
