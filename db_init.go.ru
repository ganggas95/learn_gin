package main

import (
	"learn_gin/db"
	"learn_gin/models"
)

func main(){
	engine := db.InitDB()
	var err error
	err = engine.Sync(new(models.User))
	println(err)
	err = engine.Sync(new(models.Todo))

	user := models.User{Username:"admin"}
	engine.Get(&user)
	println(&user)
	if &user == nil {
		user.Username = "admin"
		password, err := models.CreatePassword("admin")
		if err == nil{
			print(password)
			user.Password = password
			engine.Insert(&user)
			println("User admin has ben created!")
		}
		panic(err)
	}

}