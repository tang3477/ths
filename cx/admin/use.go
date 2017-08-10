package admin

import (
	"fmt"
	//	"reflect"
	//	"strconv"
	//	"log"
	//	"encoding/json"
	//	"fmt"
	"net/http"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	name := r.FormValue("username")
	pwd := r.FormValue("password")
	xiala := r.FormValue("xiala")
	fmt.Println(name)
	fmt.Println(pwd)
	fmt.Println(xiala)
	if xiala == "0" {
		sql := "select * from t_stu where account=" + "'" + name + "'" + "and" + " " + "pwd=" + "'" + pwd + "';"
		affected1, _ := engine.Query(sql)
		if len(affected1) == 0 {
			w.Write([]byte("{'res':'密码或账号错误'}"))
		} else {
			w.Write([]byte("{'res':0}"))
		}
	} else {
		sql := "select * from t_adm where account=" + "'" + name + "'" + "and" + " " + "pwd=" + "'" + pwd + "';"
		affected1, _ := engine.Query(sql)
		if len(affected1) == 0 {
			w.Write([]byte("{'res':'密码或账号错误'}"))
		} else {
			w.Write([]byte("{'res':0}"))
		}
	}
	//	login := make([]map[string]string, len(affected1))
	//	for i, row := range affected1 {
	//		login[i] = make(map[string]string)
	//		for k, v := range row {
	//			login[i][k] = string(v)
	//			if login[i][k] == "0" {
	//				w.Write([]byte("{'res':1}"))
	//			} else {
	//				w.Write([]byte("{'res':0}"))
	//			}
	//		}
	//	}

}
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	active := r.FormValue("active")
	zhanghao := r.FormValue("zhanghao")
	shoujihao := r.FormValue("shoujihao")
	email := r.FormValue("email")
	mima := r.FormValue("mima")
	fmt.Println(active)
	if active == "0" {
		sql := "select * from t_stu where account=" + zhanghao
		affected1, _ := engine.Query(sql)
		if len(affected1) == 0 {
			engine.Exec("insert into t_stu(account,pwd,tel,e_mail) values(?,?,?,?);", zhanghao, mima, shoujihao, email)
			w.Write([]byte("{'res':0}"))
		} else {
			w.Write([]byte("{'res':'账户冲突'}"))
		}
	} else {
		sql := "select * from t_aid where code=" + "'" + active + "'"
		affected1, _ := engine.Query(sql)
		if len(affected1) == 0 {
			w.Write([]byte("{'res':'工号不存在'}"))
		} else {
			sql := "select * from t_adm where code=" + "'" + active + "'"
			affected1, _ := engine.Query(sql)
			if len(affected1) == 0 {
				sql := "select * from t_adm where account=" + zhanghao
				affected1, _ := engine.Query(sql)
				if len(affected1) == 0 {
					engine.Exec("insert into t_adm(account,pwd,tel,e_mail,code) values(?,?,?,?,?);", zhanghao, mima, shoujihao, email, active)
					w.Write([]byte("{'res':0}"))
				} else {
					w.Write([]byte("{'res':'账户冲突'}"))
				}
			} else {
				w.Write([]byte("{'res':'工号冲突'}"))
			}
		}
	}

}
