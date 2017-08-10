package admin

import (
	"fmt"
	//	"strconv"
	//	"io"
	//	"os"
	//	"reflect"
	//	"strconv"
	//	"log"
	"encoding/json"
	//	"fmt"
	"net/http"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func Resultt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	account := r.FormValue("account")
	sql := "select t_apply.nam,t_apply.account,t_apply.tim,t_apply.id from t_apply where account=" + account
	affected1, _ := engine.Query(sql)
	getbook := make([]map[string]string, len(affected1))
	fmt.Println(len(affected1))
	for i, row := range affected1 {
		getbook[i] = make(map[string]string)
		for k, v := range row {
			getbook[i][k] = string(v)
		}
	}
	tr, er := json.Marshal(getbook)
	if er != nil {
		panic(er)
	}
	w.Write([]byte("{'res':0,'data':" + string(tr) + "}"))
}

func Resulttid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	anid := r.FormValue("anid")
	sql := "select * from t_apply where id=" + anid
	affected1, _ := engine.Query(sql)
	getbook := make([]map[string]string, len(affected1))
	fmt.Println(len(affected1))
	fmt.Println(sql)
	for i, row := range affected1 {
		getbook[i] = make(map[string]string)
		for k, v := range row {
			getbook[i][k] = string(v)
		}
	}
	tr, er := json.Marshal(getbook)
	if er != nil {
		panic(er)
	}
	w.Write([]byte("{'res':0,'data':" + string(tr) + "}"))
}

//func Move(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Origin", "*")

//	r.ParseForm()
//	//	move := r.FormValue("move")
//	originalFile, _ := os.Open("E:\test.txt")

//	defer originalFile.Close()

//	// Create new file
//	newFile, _ := os.Create("D:\test_copy.txt")

//	defer newFile.Close()

//	// Copy the bytes to destination from source
//	_, _ = io.Copy(newFile, originalFile)

//	// Commit the file contents
//	// Flushes memory to disk
//	_ = newFile.Sync()

//}
