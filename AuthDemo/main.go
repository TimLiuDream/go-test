package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/password"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/twitter"
	"github.com/qor/auth/providers/phone"
	"github.com/qor/session/manager"
)

var (
	// gormDB,_ = gorm.Open("sqlite3","sample.db")

	gormDB,_ = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", "root", "lft8306488", "localhost", "3306", "qor_example"))

	Auth = auth.New(&auth.Config{
		DB:gormDB,
	})
)

func init()  {
	gormDB.AutoMigrate(&auth_identity.AuthIdentity{})

	Auth.RegisterProvider(password.New(&password.Config{}))

	Auth.RegisterProvider(github.New(&github.Config{
		ClientID:"github client id",
		ClientSecret :"github client secret",
	}))

	Auth.RegisterProvider(facebook.New(&facebook.Config{
		ClientID:     "facebook client id",
    ClientSecret: "facebook client secret",
	}))

	Auth.RegisterProvider(google.New(&google.Config{
		ClientID:     "facebook client id",
    	ClientSecret: "facebook client secret",
	}))

	Auth.RegisterProvider(twitter.New(&twitter.Config{
		ClientID:     "twitter client id",
		ClientSecret: "twitter client secret",
	  }))

	  Auth.RegisterProvider(phone.New(&phone.Config{}))
}

func main()  {
	mux:=http.NewServeMux()

	mux.Handle("/auth/",Auth.NewServeMux())

	http.ListenAndServe(":9000",manager.SessionManager.Middleware(mux))
}