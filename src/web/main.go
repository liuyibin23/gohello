package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"utilslib"
)

var (
	sessionMgr *SessionManager = nil
)

func hello(w http.ResponseWriter, r *http.Request) {

	sessionMgr.maxAge = 30
	session := sessionMgr.BeginSession(w, r)
	session.GetId()
	fmt.Fprintf(w, "Hello!,sessionid = %s", session.GetId())
	//for i:=0;i<300;i++{
	//	log.Printf("Hello!,sessionid = %s", session.GetId())
	//}

	for i:=0;i<300;i++{
		logger1.Info("Hello!,sessionid = %s", session.GetId())
	}
	//fmt.Fprintf(w,"Hello!")
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//time.AfterFunc(5 * time.Second, func() {
	//	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//})

}

func hello2(w http.ResponseWriter, r *http.Request){
	sessionMgr.maxAge = 30
	session := sessionMgr.BeginSession(w, r)
	session.GetId()
	fmt.Fprintf(w, "Hello2222222222!,sessionid = %s", session.GetId())
	//for i:=0;i<100;i++{
	//	log.Printf("Hello!,sessionid = %s", session.GetId())
	//}
	for i:=0;i<100;i++{
		logger1.Info("Hello!,sessionid = %s", session.GetId())
	}
}

func logFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		//log.Println("Handler function called - " + name)
		//logger.Println("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		//logger1.Info("Handler function called - " + name)
		h(w, r)
	}
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//var logger *log.Logger
var logger1 *utilslib.Logger

func init() {

	//f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//logger = log.New(f, "", log.LstdFlags|log.Lshortfile)
	//// 设置日志抬头信息
	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	sessionMgr = NewSessionManager()

	logger1 = utilslib.NewFileLogger("test.log", 1, log.LstdFlags|log.Lshortfile, 2)
	log.SetOutput(&lumberjack.Logger{
		Filename:"test1.log",
		MaxSize: 1,
		LocalTime:true,
	})
}

func main() {

	world := WorldHandler{}
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	//http.HandleFunc ("/hello", hello)
	http.HandleFunc("/hello", logFunc(hello))
	http.HandleFunc ("/hello2", hello2)
	http.Handle("/world", &world)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/resetpwd", resetpwd)
	server.ListenAndServe()
}
