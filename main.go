package main

import (
	"MANOMASH-Server/database"
	"MANOMASH-Server/handler"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	c := cors.AllowAll()

	//DB初期化の前にGoがDBに接続しにいって失敗する
	time.Sleep(time.Second * 5)
	database.GormConnect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	fmt.Println("hallo,world")
	r.HandleFunc("/login", handler.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", handler.SignUpHandler).Methods(http.MethodPost)
	r.HandleFunc("/mypage/edit", handler.MyPageEditHandler).Methods(http.MethodPost)
	r.HandleFunc("/mypage", handler.MyPageHandler).Methods(http.MethodGet)
	r.HandleFunc("/profile", handler.ProfileHandler).Methods(http.MethodGet)
	r.HandleFunc("/profile/add", handler.ProfileAddHandler).Methods(http.MethodPost)
	r.HandleFunc("/profile/edit", handler.ProfileEditHandler).Methods(http.MethodPost)
	r.HandleFunc("/profile/delete", handler.ProfileDeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/profile/list", handler.ProfileListHandler).Methods(http.MethodPost)
	r.HandleFunc("/gallery", handler.GalleryHandler).Methods(http.MethodGet)

	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}
