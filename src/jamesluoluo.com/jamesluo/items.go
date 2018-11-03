package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    )
/*
        LOCATION VARCHAR(255),
        TIMESTAMP VARCHAR(255),
        NAME VARCHAR(255),
        FULLDESCRIPTION VARCHAR(255),
        VALUE VARCHAR(255),
        CATEGORY VARCHAR(255)
*/
type Item struct {
    LOCATION string `json:"location"`
    TIMESTAMP string `json:"timestamp"`
    NAME string`json:"name"`
    FULLDESCRIPTION string `json:"fulldescription"`
    VALUE string`json:"value"`
    CATEGORY string`json:"category"`

}
type Items []Item
func searchByCategory(db *sql.DB) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        m := post_request_resolver(db , r)
        user := m["username"][0]
        pw := m["pw"][0]
        cat := m["category"][0]

        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        //query
        val, err := db.Query("SELECT LOCATION, TIMESTAMP, NAME, FULLDESCRIPTION, VALUE, CATEGORY FROM items WHERE CATEGORY=$1", cat)
        var (
            LOCATION string
            TIMESTAMP string
            NAME string
            FULLDESCRIPTION string
            VALUE string
            CATEGORY string
            )
        var it_list Items
        for val.Next() {
            err := val.Scan(&LOCATION, &TIMESTAMP, &NAME, &FULLDESCRIPTION, &VALUE, &CATEGORY)
            if err != nil {
                log.Fatal(err)
            }
             
            out := &Item{
                LOCATION:LOCATION,
                TIMESTAMP: TIMESTAMP,
                NAME:NAME,
                FULLDESCRIPTION:FULLDESCRIPTION,
                VALUE:VALUE,
                CATEGORY:CATEGORY,
            }
            it_list = append(it_list, *out)
            
        //return json of sql read
        }
        defer val.Close()
        if err == nil && len(it_list)>0 {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(it_list)
            w.Write(send)
        }else{
            w.Write([]byte("0"))
        }
    }
    return http.HandlerFunc(fn)

}
func searchByCategory_loc(db *sql.DB) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        m := post_request_resolver(db , r)
        location := m["location"][0]
        user := m["username"][0]
        pw := m["pw"][0]
        cat := m["category"][0]
        fmt.Println("--cat_loc--")
        fmt.Println("location:"+location)
        fmt.Println("username:" + user)
        fmt.Println("pw:" + pw)
        fmt.Println("category:" + cat)
        fmt.Println("---------")
        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        //query
        val, err := db.Query("SELECT LOCATION, TIMESTAMP, NAME, FULLDESCRIPTION, VALUE, CATEGORY FROM items WHERE CATEGORY=$1 AND LOCATION=$2", cat, location)
        var (
            LOCATION string
            TIMESTAMP string
            NAME string
            FULLDESCRIPTION string
            VALUE string
            CATEGORY string
            )
        var it_list Items
        for val.Next() {
            err := val.Scan(&LOCATION, &TIMESTAMP, &NAME, &FULLDESCRIPTION, &VALUE, &CATEGORY)
            if err != nil {
                log.Fatal(err)
            }
             
            out := &Item{
                LOCATION:LOCATION,
                TIMESTAMP: TIMESTAMP,
                NAME:NAME,
                FULLDESCRIPTION:FULLDESCRIPTION,
                VALUE:VALUE,
                CATEGORY:CATEGORY,
            }
            it_list = append(it_list, *out)
            
        //return json of sql read
        }
        defer val.Close()
        if err == nil && len(it_list)>0 {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(it_list)
            w.Write(send)
        }else{
            w.Write([]byte("0"))
        }
    }
    return http.HandlerFunc(fn)

}


func searchByName(db *sql.DB ) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        m := post_request_resolver(db , r)
        user := m["username"][0]
        pw := m["pw"][0]
        name := m["name"][0]

        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        //query
        val, err := db.Query("SELECT LOCATION, TIMESTAMP, NAME, FULLDESCRIPTION, VALUE, CATEGORY FROM items WHERE NAME LIKE $1", string('%')+name+string('%'))
        var (
            LOCATION string
            TIMESTAMP string
            NAME string
            FULLDESCRIPTION string
            VALUE string
            CATEGORY string
            )
        var it_list Items

        for val.Next() {
            err := val.Scan(&LOCATION, &TIMESTAMP, &NAME, &FULLDESCRIPTION, &VALUE, &CATEGORY)
            if err != nil {
                log.Fatal(err)
            }
             
            out := &Item{
                LOCATION:LOCATION,
                TIMESTAMP: TIMESTAMP,
                NAME:NAME,
                FULLDESCRIPTION:FULLDESCRIPTION,
                VALUE:VALUE,
                CATEGORY:CATEGORY,
            }
            it_list = append(it_list, *out)
            
        //return json of sql read
    }
    defer val.Close()
        if err == nil && len(it_list)>0  {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(it_list)
            w.Write(send)
        }else{
            w.Write([]byte("0"))
        }
        
    }
    return http.HandlerFunc(fn)
}
func searchByName_loc(db *sql.DB ) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        m := post_request_resolver(db , r)
        location := m["location"][0]
        user := m["username"][0]
        pw := m["pw"][0]
        name := m["name"][0]
        fmt.Println("--name_loc--")
        fmt.Println("location:"+location)
        fmt.Println("username:" + user)
        fmt.Println("pw:" + pw)
        fmt.Println("name:" + name)
        fmt.Println("---------")

        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        //query
        val, err := db.Query("SELECT LOCATION, TIMESTAMP, NAME, FULLDESCRIPTION, VALUE, CATEGORY FROM items WHERE NAME LIKE $1 AND LOCATION=$2", string('%')+name+string('%'), location)
        var (
            LOCATION string
            TIMESTAMP string
            NAME string
            FULLDESCRIPTION string
            VALUE string
            CATEGORY string
            )
        var it_list Items

        for val.Next() {
            err := val.Scan(&LOCATION, &TIMESTAMP, &NAME, &FULLDESCRIPTION, &VALUE, &CATEGORY)
            if err != nil {
                log.Fatal(err)
            }
             
            out := &Item{
                LOCATION:LOCATION,
                TIMESTAMP: TIMESTAMP,
                NAME:NAME,
                FULLDESCRIPTION:FULLDESCRIPTION,
                VALUE:VALUE,
                CATEGORY:CATEGORY,
            }
            it_list = append(it_list, *out)
            
        //return json of sql read
    }
    defer val.Close()
        if err == nil && len(it_list)>0 {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(it_list)
            w.Write(send)
        }else{
            w.Write([]byte("0"))
        }
        
    }
    return http.HandlerFunc(fn)
}


func get_items(db *sql.DB ) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        m := post_request_resolver(db , r)
        location := m["location"][0]
        user := m["username"][0]
        pw := m["pw"][0]
        fmt.Println("--getitems--")
        fmt.Println("location:"+location)
        fmt.Println("username:" + user)
        fmt.Println("pw:" + pw)
        fmt.Println("---------")
        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        /////////////////////////////////////////////////
        val, err := db.Query("select LOCATION, TIMESTAMP, NAME, FULLDESCRIPTION, VALUE, CATEGORY from items WHERE LOCATION=$1", location)
        //verify first
        var (
            LOCATION string
            TIMESTAMP string
            NAME string
            FULLDESCRIPTION string
            VALUE string
            CATEGORY string
            )
        var it_list Items

        for val.Next() {
            err := val.Scan(&LOCATION,&TIMESTAMP,&NAME,&FULLDESCRIPTION,&VALUE,&CATEGORY)
            if err != nil {
                log.Fatal(err)
            }
            out := &Item{
                LOCATION:LOCATION,
                TIMESTAMP: TIMESTAMP,
                NAME:NAME,
                FULLDESCRIPTION:FULLDESCRIPTION,
                VALUE:VALUE,
                CATEGORY:CATEGORY,
            }
            it_list = append(it_list, *out)
            
           
        }

        defer val.Close()
        if err == nil && len(it_list)>0  {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(it_list)
            w.Write(send)
        }else{
            w.Write([]byte("0"))
        }
        
    }
    return http.HandlerFunc(fn)
}

func add_item(db *sql.DB ) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        /*if (session_check(db, r.URL.Query().Get("sess")) != ""){

            }else{
                w.Write([]byte("0"))
            }*/
            m := post_request_resolver(db , r)
            location := m["location"][0]
            timestamp := m["timestamp"][0]
            name := m["name"][0]
            fulldescription := m["fulldescription"][0]
            value := m["value"][0]
            category := m["category"][0]

        _, err := db.Query(`insert into items (
            LOCATION
            ,TIMESTAMP
            ,NAME
            ,FULLDESCRIPTION
            ,VALUE
            ,CATEGORY
            ) values ($1, $2, $3, $4, $5, $6)
            `, location, timestamp, name, fulldescription, value, category)

        //verify first
        if err == nil {
            w.Write([]byte("1"))
        }else{
            w.Write([]byte("0"))
        }

        
    }
    return http.HandlerFunc(fn)
}
