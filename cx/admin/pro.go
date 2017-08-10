package admin

import (
	"fmt"
	//	"strconv"
	"time"
	//	"reflect"
	//	"strconv"
	//	"log"
	"encoding/json"
	//	"fmt"
	"net/http"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

func Apply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	xmname := r.FormValue("xmname")
	xmtype := r.FormValue("xmtype")
	xmpeople := r.FormValue("xmpeople")
	direct := r.FormValue("direct")
	cont := r.FormValue("cont")
	acco := r.FormValue("acco")
	fmt.Println(xmname)
	fmt.Println(xmtype)
	fmt.Println(xmpeople)
	fmt.Println(direct)
	fmt.Println(cont)
	fmt.Println(acco)
	t := time.Now().Unix()
	engine.Exec("insert into t_apply(lx,nam,dire,tim,ppe,cntent,account) values(?,?,?,?,?,?,?);", xmtype, xmname, direct, t, xmpeople, cont, acco)
	w.Write([]byte("{'res':0}"))

}
func Anal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	sql := "select t_apply.nam,t_apply.account,t_apply.tim,t_apply.id from t_apply"
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

func Analid(w http.ResponseWriter, r *http.Request) {
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

func Yes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	anid := r.FormValue("anid")
	sql := "update t_apply set adopt=1 where id=" + anid
	_, _ = engine.Query(sql)
	w.Write([]byte("{'res':0}"))
}

func No(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	anid := r.FormValue("anid")
	sql := "update t_apply set adopt=0 where id=" + anid
	_, _ = engine.Query(sql)
	w.Write([]byte("{'res':0}"))
}
