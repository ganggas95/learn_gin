package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type User struct{
	IdUser int64 `xorm:"'id' pk autoincr" json:"id_user"`
	Username string `xorm:"varchar(20) not null unique 'username'" json: "username" binding:"required"`
	Password []byte `xorm:"varchar(255) not null 'password'" json:"password"`
}

func CreatePassword(password string) ([]byte, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func CheckPassword(has_pass []byte, password string) error{
	err := bcrypt.CompareHashAndPassword(has_pass, []byte(password))
	if err != nil {
		panic(err)
		return err
	}
	return err
}

func GenerateToken(user *User) (string, error){
	token:= jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"Id":  string(user.Username),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	return token.SignedString([]byte("b1sm1llahirahmanirrahim1234567890987676765656"))
}
