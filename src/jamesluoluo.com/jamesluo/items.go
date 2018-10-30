package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    )


func searchByCategory(db *sql.DB) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        cat := r.URL.Query().Get("category")
        username := r.URL.Query().Get("username")
        uid := r.URL.Query().Get("uid")
        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        //query
        val, err := db.Query("SELECT DISTINCT category FROM items")
        var (
            category string
            out *Item
            )
        for val.Next() {
            err := val.Scan(&category)
            if err != nil {
                log.Fatal(err)
            }
            out = &Item{
                category:category
            }

            break
        //return json of sql read
    }
    defer val.Close()
        if err == nil && val != nil {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(out)
            w.Write(send)
        }else{
            w.Write([]byte("0"))
        }
    }
    return http.HandlerFunc(fn)

}


func searchByName(db *sql.DB ) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        username := r.URL.Query().Get("username")
        uid := r.URL.Query().Get("uid")
        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        //query
        val, err := db.Query("SELECT DISTINCT name FROM items")
        var (
            item_name string
            out *Item
            )
        for val.Next() {
            err := val.Scan(&item_name)
            if err != nil {
                log.Fatal(err)
            }
            out = &Item{
                item_name:item_name
            }

            break
        //return json of sql read
    }
    defer val.Close()
        if err == nil && val != nil {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(out)
            w.Write(send)
        }else{
            w.Write([]byte("0"))
        }
        
    }
    return http.HandlerFunc(fn)
}


func get_items(db *sql.DB ) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        item := r.URL.Query().Get("item")
        username := r.URL.Query().Get("username")
        uid := r.URL.Query().Get("uid")
        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        /////////////////////////////////////////////////
        val, err := db.Query("select LOCATION, TIMESTAMP, NAME, FULLDESCRIPTION, VALUE, CATEGORY from items")
        //verify first
        var (
            location string 
            timestamp string
            name string 
            fulldescription string
            value string
            category string
            out *ITEM
        )
        for val.Next() {
            err := val.Scan(&location,&timestamp,&name,&fulldescription,&value,&category)
            if err != nil {
                log.Fatal(err)
            }
            out = &Item{
                location:location,
                timestamp:timestamp,
                name:name,
                fulldescription:fulldescription,
                value:value,
                category:category,

            }
            
            break
        }

        defer val.Close()
        if err == nil && val != nil {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(out)
            w.Write(send)
        }else{
            w.Write([]byte("0"))
        }
        
    }
    return http.HandlerFunc(fn)
}


func add_items(db *sql.DB ) http.HandlerFunc{
    fn := func(w http.ResponseWriter, r *http.Request) {
        if (session_check(db, r.URL.Query().Get("sess")) != ""){

            }else{
                w.Write([]byte("0"))
            }
            m := post_request_resolver(db , r)
            location := m["location"][0]
            timestamp := m["timestamp"][0]
            name := m["name"][0]
            fulldescription := m["fulldescription"][0]
            value := m["value"][0]
            category := m["category"][0]

        //location := r.URL.Query().Get("location")
        //timestamp := r.URL.Query().Get("timestamp")
        //name := r.URL.Query().Get("name")
        //fulldescription := r.URL.Query().Get("fulldescription")
        //value := r.URL.Query().Get("value")
        //category := r.URL.Query().Get("category")
        val, err := db.Query(`insert into items (
            LOCATION
            ,TIMESTAMP
            ,NAME
            ,FULLDESCRIPTION
            ,VALUE
            ,CATEGORY
            `, location, timestamp, name, fulldescription, value, category)

        //verify first
        if err == nil && val != nil {
            w.Write([]byte("1"))
        }else{
            w.Write([]byte("0"))
        }

        
    }
    return http.HandlerFunc(fn)
}
