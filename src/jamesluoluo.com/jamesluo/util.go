package main
import (
	"database/sql"
    _ "github.com/lib/pq"
	"fmt"
	"log"
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "net/url"
    "io/ioutil"
    )
func check_uid(db *sql.DB, user string, pw string) bool{
	pw_md5 , err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
        if err != nil {
            log.Println(err)
        }
    val, err := db.Query("select * where ID_EMAIL=$1 and PW=$2",user, string(pw_md5))
        if err == nil && val.Next() {
            return true
        }else{
            return false
        }
}

func post_request_resolver(db *sql.DB, r *http.Request) map[string][]string {
		b, _ := ioutil.ReadAll(r.Body)
        defer r.Body.Close()
        
        fmt.Println(string(b))
        // Unmarshal
        tempurl:= "http://www.jamesluoluo.com/sql?"+string(b)
        u, err2 := url.Parse(tempurl)
        if err2 != nil {
            panic(err2)
        }

        m, _ := url.ParseQuery(u.RawQuery)
        return m

}