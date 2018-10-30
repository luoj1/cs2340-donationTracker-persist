package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
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
    NAME string `json:"name"`
    LATITUDE string `json:"latitude"`
    LONGITUDE string`json:"longitude"`
    STREET_ADDR string `json:"street_addr"`
    CITY string`json:"city"`
    STATE string`json:"state"`
    ZIP string`json:"zip"`
    TYPE string`json:"type"`
    PHONE string`json:"phone"`
    WEBSITE string`json:"website"`

}
type Locations []Location

func get_locations(db *sql.DB)  http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        m := post_request_resolver(db , r)
        user := m["username"][0]
        pw := m["pw"][0]
        fmt.Println("pw in locations"+pw)
        if !check_uid(db, user, pw ) {
            w.Write([]byte("0"))
            return
        }
        
        val, err := db.Query("select NAME, LATITUDE, LONGITUDE, STREET_ADDR, CITY, STATE, ZIP, TYPE, PHONE, WEBSITE from locations")
        //verify first
        var (
            NAME string
            LATITUDE string 
            LONGITUDE string
            STREET_ADDR string 
            CITY string
            STATE string
            ZIP string
            TYPE string
            PHONE string
            WEBSITE string
            //out *Location
        )
        var loc_list Locations 
        for val.Next() {
            err := val.Scan(&NAME,&LATITUDE,&LONGITUDE,&STREET_ADDR,&CITY,&STATE,&ZIP,&TYPE,&PHONE, &WEBSITE)
            if err != nil {
                log.Fatal(err)
            }
            out := &Location{
                NAME:NAME,
                LATITUDE:LATITUDE,
                LONGITUDE:LONGITUDE,
                STREET_ADDR:STREET_ADDR,
                TYPE: TYPE,
                CITY:CITY,
                STATE:STATE,
                ZIP:ZIP,
                PHONE: PHONE,
                WEBSITE: WEBSITE,
            }
            
            loc_list = append(loc_list, *out)
            
        }

        defer val.Close()
        if err == nil {
            w.Header().Set("Content-Type", "application/json")
            send , _ := json.Marshal(loc_list)
            w.Write(send)
        }else{
            panic(err)
            w.Write([]byte("0"))
        }

    }
    return http.HandlerFunc(fn)
}