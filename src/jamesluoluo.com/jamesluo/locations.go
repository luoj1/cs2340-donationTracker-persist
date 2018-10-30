package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    //"fmt"
    "log"
    "net/http"
    "encoding/json"
    )

func add_location(db *sql.DB) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        if (session_check(db, r.URL.Query().Get("sess")) != ""){

            }else{
                w.Write([]byte("0"))
            }
        email := r.URL.Query().Get("email")
        name := r.URL.Query().Get("name")
        latitude := r.URL.Query().Get("latitude")
        longitude := r.URL.Query().Get("longitude")
        street_addr := r.URL.Query().Get("street_addr")
        city := r.URL.Query().Get("city")
        st :=r.URL.Query().Get("state")
        zip := r.URL.Query().Get("zip")
        t := r.URL.Query().Get("type")
        phone := r.URL.Query().Get("phone")
        website := r.URL.Query().Get("website")
        val, err := db.Query(`insert into locations (
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
                )`, email, name, latitude, longitude, street_addr, city, st, zip, t, phone, website)

        //verify first
        if err == nil && val != nil {
            w.Write([]byte("1"))
        }else{
            w.Write([]byte("0"))
        }

    }
    return http.HandlerFunc(fn)
}
type Location struct {
    latitude string 
    longitude string
    street_addr string 
    city string
    state string
    zip string
    TYPE string`json:"type"`
    phone string
    website string

}
func get_locations(db *sql.DB)  http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        m := post_request_resolver(db , r)
        user := m["username"][0]
        pw := m["pw"][0]
        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        
        val, err := db.Query("select LATITUDE, LONGITUDE, STREET_ADDR, CITY, STATE, ZIP, TYPE, PHONE, WEBSITE from locations")
        //verify first
        var (
            latitude string 
            longitude string
            street_addr string 
            city string
            state string
            zip string
            TYPE string
            phone string
            website string
            out *Location
        )
        for val.Next() {
            err := val.Scan(&latitude,&longitude,&street_addr,&city,&state,&zip,&TYPE,&phone, &website)
            if err != nil {
                log.Fatal(err)
            }
            out = &Location{
                latitude:latitude,
                 longitude:longitude,
                street_addr:street_addr,
                TYPE: TYPE,
                city:city,
                state:state,
                zip:zip,
                phone: phone,
                website:website,

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