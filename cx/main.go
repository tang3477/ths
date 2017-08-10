package main

import (
	bb "cx/admin"
	"fmt"

	"log"
	"net/http"

	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

type dispath struct{}

func Listen() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &dispath{},
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/admin/login"] = bb.Login
	mux["/admin/register"] = bb.Register
	mux["/admin/apply"] = bb.Apply
	mux["/admin/anal"] = bb.Anal
	mux["/admin/analid"] = bb.Analid
	mux["/admin/yes"] = bb.Yes
	mux["/admin/no"] = bb.No
	mux["/admin/resultt"] = bb.Resultt
	mux["/admin/resulttid"] = bb.Resulttid
	mux["/admin/xysq"] = bb.Xysq
	mux["/admin/glxz"] = bb.Glxz
	fmt.Println("Listen is prot :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (*dispath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hf, ok := mux[r.URL.Path]; ok {
		hf(w, r)
	}
}

func main() {
	//时间戳

	t := time.Now().Unix()
	fmt.Println(t)

	//时间戳到具体显示的转化
	fmt.Println(time.Unix(t, 0).String())

	fmt.Println(time.Now().Format("2006年 01月 02日"))
	Listen()

}
