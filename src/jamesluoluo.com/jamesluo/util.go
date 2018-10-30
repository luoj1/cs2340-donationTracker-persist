package main
import (
	"database/sql"
    _ "github.com/lib/pq"
	"fmt"
	//"log"
	"encoding/hex"
    "net/http"
    "net/url"
    "io/ioutil"
    "crypto/md5"
    )
func md5_gen(text string) string {
    algorithm := md5.New()
    algorithm.Write([]byte(text))
    return hex.EncodeToString(algorithm.Sum(nil))
}
func check_uid(db *sql.DB, user string, pw string) bool{

	pw_md5 := md5_gen(pw)
        
    fmt.Println(string(pw_md5))
    fmt.Println(user)
    val, err := db.Query("select * from users where ID_EMAIL=$1 and PW=$2",user, string(pw_md5))
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