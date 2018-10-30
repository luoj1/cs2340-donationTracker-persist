
//                            _ooOoo_  
//                           o8888888o  
//                           88" . "88  
//                           (| -_- |)  
//                            O\ = /O  
//                        ____/`---'\____  
//                      .   ' \\| |// `.  
//                       / \\||| : |||// \  
//                     / _||||| -:- |||||- \  
//                       | | \\\ - /// | |  
//                     | \_| ''\---/'' | |  
//                      \ .-\__ `-` ___/-. /  
//                   ___`. .' /--.--\ `. . __  
//                ."" '< `.___\_<|>_/___.' >'"".  
//               | | : `- \`.;`\ _ /`;.`/ - ` : | |  
//                 \ \ `-. \_ __\ /__ _/ .-` / /  
//         ======`-.____`-.___\_____/___.-`____.-'======  
//                            `=---='  
//  
//         .............................................  
//                  佛祖保佑             永无BUG 
//          佛曰:  
//                  写字楼里写字间，写字间里程序员；  
//                  程序人员写程序，又拿程序换酒钱。  
//                  酒醒只在网上坐，酒醉还来网下眠；  
//                  酒醉酒醒日复日，网上网下年复年。  
//                  但愿老死电脑间，不愿鞠躬老板前；  
//                  奔驰宝马贵者趣，公交自行程序员。  
//                  别人笑我忒疯癫，我笑自己命太贱；  
//                  不见满街漂亮妹，哪个归得程序员？
package main

import (
	"database/sql"
    _ "github.com/lib/pq"
	"fmt"
	"log"
    "net/http"
    "crypto/md5"
    "math/rand"
    "encoding/csv"
    //"encoding/json"
    "os"
    "golang.org/x/crypto/bcrypt"
    //"io/ioutil"
    //"net/url"
    
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(`
        in shenshaohua we trust

                        ::
                      :;J7, :,                        ::;7:
                      ,ivYi, ,                       ;LLLFS:
                      :iv7Yi                       :7ri;j5PL
                     ,:ivYLvr                    ,ivrrirrY2X,
                     :;r@Wwz.7r:                :ivu@kexianli.
                    :iL7::,:::iiirii:ii;::::,,irvF7rvvLujL7ur
                   ri::,:,::i:iiiiiii:i:irrv177JX7rYXqZEkvv17
                ;i:, , ::::iirrririi:i:::iiir2XXvii;L8OGJr71i
              :,, ,,:   ,::ir@mingyi.irii:i:::j1jri7ZBOS7ivv,
                 ,::,    ::rv77iiiriii:iii:i::,rvLq@huhao.Li
             ,,      ,, ,:ir7ir::,:::i;ir:::i:i::rSGGYri712:
           :::  ,v7r:: ::rrv77:, ,, ,:i7rrii:::::, ir7ri7Lri
          ,     2OBBOi,iiir;r::        ,irriiii::,, ,iv7Luur:
        ,,     i78MBBi,:,:::,:,  :7FSL: ,iriii:::i::,,:rLqXv::
        :      iuMMP: :,:::,:ii;2GY7OBB0viiii:i:iii:i:::iJqL;::
       ,     ::::i   ,,,,, ::LuBBu BBBBBErii:i:i:i:i:i:i:r77ii
      ,       :       , ,,:::rruBZ1MBBqi, :,,,:::,::::::iiriri:
     ,               ,,,,::::i:  @arqiao.       ,:,, ,:::ii;i7:
    :,       rjujLYLi   ,,:::::,:::::::::,,   ,:i,:,,,,,::i:iii
    ::      BBBBBBBBB0,    ,,::: , ,:::::: ,      ,,,, ,,:::::::
    i,  ,  ,8BMMBBBBBBi     ,,:,,     ,,, , ,   , , , :,::ii::i::
    :      iZMOMOMBBM2::::::::::,,,,     ,,,,,,:,,,::::i:irr:i:::,
    i   ,,:;u0MBMOG1L:::i::::::  ,,,::,   ,,, ::::::i:i:iirii:i:i:
    :    ,iuUuuXUkFu7i:iii:i:::, :,:,: ::::::::i:i:::::iirr7iiri::
    :     :rk@Yizero.i:::::, ,:ii:::::::i:::::i::,::::iirrriiiri::,
     :      5BMBBBBBBSr:,::rv2kuii:::iii::,:i:,, , ,,:,:i@petermu.,
          , :r50EZ8MBBBBGOBBBZP7::::i::,:::::,: :,:,::i;rrririiii::
              :jujYY7LS0ujJL7r::,::i::,::::::::::::::iirirrrrrrr:ii:
           ,:  :@kevensun.:,:,,,::::i:i:::::,,::::::iir;ii;7v77;ii;i,
           ,,,     ,,:,::::::i:iiiii:i::::,, ::::iiiir@xingjief.r;7:i,
        , , ,,,:,,::::::::iiiiiiiiii:,:,:::::::::iiir;ri7vL77rrirri::
         :,, , ::::::::i:::i:::i:i::,,,,,:,::i:i:::iir;@Secbone.ii:::


        `))
}

func verify_handler(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
        user := r.URL.Query().Get("user")
        uid := r.URL.Query().Get("uid")
        val, err := db.Query("select * where ID_EMAIL=$1 and PW=$2",user, md5.Sum([]byte(uid)))
        if err == nil && val.Next() {
            rand := random(250)
            db.Query("insert into sessions (ID_EMAIL, SESSION) values ($1, $2)",user, rand)
        	w.Write([]byte(rand))
        }else{
        	w.Write([]byte("0"))
        }

    }

    return http.HandlerFunc(fn)
}

func random(n int) string {
    b := make([]byte, n)
    var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}


type UserInfo struct {
    user string
    email string
    user_type string
    pw string
}
func create_user(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        m := post_request_resolver(db , r)

        user := m["user"][0]
        email := m["email"][0]
        pw := m["pw"][0]
        //pw_md5 :=  md5.Sum([]byte(pw))
        pw_md5 , err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
        if err != nil {
            log.Println(err)
        }
        fmt.Println(string(pw_md5))
        u_t := m["user_type"][0]


        fmt.Println("pw"+pw)
        fmt.Println("email"+email)
        fmt.Println("user"+user)
        fmt.Println("u_t"+u_t)
        val , err2 := db.Query("select * from users where ID_EMAIL=$1 or PW=$2", email, string(pw_md5))
        if err2 != nil || val.Next() {
            log.Fatal(err2)
            w.Write([]byte("0"))
            return 
        }

        _, err = db.Query("insert into users (ID_EMAIL, USERNAME, PW, USER_TYPE, ACCOUNT_STATE) VALUES ($1, $2, $3, $4, $5)", email,user, string(pw_md5), u_t, 1)
        fmt.Println("finish write")
        if err == nil {
            fmt.Println("1")
            w.Write([]byte("1"))
        }else{
            fmt.Println("0")
            log.Fatal(err)
            w.Write([]byte("0"))
        }

    }
    return http.HandlerFunc(fn)
}

func edit_property(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        user := r.URL.Query().Get("user")
        pw := r.URL.Query().Get("pw")
        val, err := db.Query("select * where user=$1 and pw=$2",user, md5.Sum([]byte(pw)))
        //verify first
        if err == nil && val != nil {

        }else{
            w.Write([]byte("0"))
        }

    }
    return http.HandlerFunc(fn)
}


func get_user_type(db *sql.DB,  email string) string {
    var u_t string = ""
    val, err := db.Query("select USER_TYPE from users where ID_EMAIL=$1", email)
    //verify first
    if err == nil && val != nil {
        val.Next()
        val.Scan(&u_t)
        return u_t
    }else{
        return ""
    }
}
func session_check(db *sql.DB,  sess string) string {
    var id_email string = ""
    val, err := db.Query("select ID_EMAIL from sessions where SESSION=$1", sess)
    //verify first
    if err == nil && val != nil {
        val.Next()
        val.Scan(&id_email)
        return id_email
    }else{
        return ""
    }
}
func buildServer(db *sql.DB) {
    _, err:=db.Query(`
        CREATE TABLE IF NOT EXISTS users (
        KEY SERIAL PRIMARY KEY,
        ID_EMAIL VARCHAR(255),
        USERNAME VARCHAR(255),
        PW VARCHAR(255),
        USER_TYPE VARCHAR(255),
        LOCATION VARCHAR(255),
        ACCOUNT_STATE INTEGER,
        CONTACT_INFO VARCHAR(255)
    );
    
    CREATE TABLE IF NOT EXISTS locations (
        KEY SERIAL PRIMARY KEY,
        ID_EMAIL VARCHAR(255),
        NAME VARCHAR(255),
        LATITUDE FLOAT8,
        LONGITUDE FLOAT8,
        STREET_ADDR VARCHAR(255),
        CITY VARCHAR(255),
        STATE CHAR(2),
        ZIP INTEGER,
        TYPE VARCHAR(255),
        PHONE VARCHAR(255),
        WEBSITE VARCHAR(255)
    );
    CREATE TABLE IF NOT EXISTS sessions (
        KEY SERIAL PRIMARY KEY,
        ID_EMAIL VARCHAR(255),
        SESSION VARCHAR(255)
    );
    `)
    if err != nil {
        log.Fatal(err)
    }
}

func locationsIniter(path string, db *sql.DB) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    record, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error", err)
    }
    for value:= range record{ // for i:=0; i<len(record)
        if value == 0 {
            continue
        }
        row := record[value]
        name := row[1]
        fmt.Println("writing" + name)
        isRepeat, _ := db.Query("select NAME from locations where NAME=$1", name)
        //fmt.Println("writing" + isRepeat)
        if isRepeat.Next() {
            continue
        }
        latitude := row[2]
        longitude := row[3]
        street_addr := row[4]
        city := row[5]
        st :=row[6]
        zip := row[7]
        t := row[8]
        phone := row[9]
        website := row[10]
        _, err :=db.Query(`insert into locations (
            ID_EMAIL
            ,NAME
            ,LATITUDE
            ,LONGITUDE
            ,STREET_ADDR
            ,CITY
            ,STATE
            ,ZIP
            ,TYPE
            ,PHONE
            ,WEBSITE) VALUES (
            $1,$2,$3, $4, $5, $6, $7, $8, $9, $10, $11
                )`, "test@test.com", name, latitude, longitude, street_addr, city, st, zip, t, phone, website)
        if err != nil {
            fmt.Println("Error", err)
        }
    }

}
func main() {
	fmt.Println("start")
	connStr := "user=jamesluo dbname=cs2340 port=5432 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("server initialized!")
    buildServer(db)
    locationsIniter("../src/jamesluoluo.com/jamesluo/LocationData.csv",db)
    
    http.HandleFunc("/", handler)
    http.HandleFunc("/verify", verify_handler(db))
    http.HandleFunc("/register", create_user(db))
    http.HandleFunc("/getLocation", get_locations(db))
    http.HandleFunc("/editLocation", add_location(db))
    //http.HandleFunc("/add", add_handler)
    //http.HandleFunc("/contain", contain_handler)
    log.Fatal(http.ListenAndServe(":8080", nil))

}
