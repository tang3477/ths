package admin

import (
	"fmt"
	"io"
	"os"
	//	"reflect"
	//	"strconv"
	//	"log"
	//	"encoding/json"
	//	"fmt"
	"net/http"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func Xysq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()
	diz := r.FormValue("account")
	anid := r.FormValue("anid")

	sql := "update t_apply set dizhi=" + "'" + diz + "'" + " where id=" + anid

	_, _ = engine.Query(sql)

	w.Write([]byte("{'res':0}"))
}
func Glxz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=515810963 dbname=csngdzb host=127.0.0.1 port=5432 sslmode=disable")
	r.ParseForm()

	xzdz := r.FormValue("xz")
	fmt.Println(xzdz)
	originalFile, _ := os.Open("../hehe/" + xzdz)

	defer originalFile.Close()

	// Create new file
	newFile, _ := os.Create("../haha/" + xzdz)

	defer newFile.Close()

	// Copy the bytes to destination from source
	_, _ = io.Copy(newFile, originalFile)

	// Commit the file contents
	// Flushes memory to disk
	_ = newFile.Sync()
	w.Write([]byte("{'res':0}"))
}
