package models

import (
	"database/sql"
	"log"
	"main/config"
)

//user登録
func KariUserRegistration(username, email, hashpassword, token string) string {
	var err error
	DbConnection, err = sql.Open(config.Config.DBdriver, ConnectionInfo())
	defer DbConnection.Close()

	if err != nil {
		log.Fatalln(err)
	}
	//userにINSERTする

	var Token string
	err = DbConnection.QueryRow("INSERT INTO kari_users(username, password, email, token) VALUES($1, $2, $3, $4) RETURNING token", username, hashpassword, email, token).Scan(&Token)
	if err != nil {
		log.Println(err)
	}
	return Token
}

func KariUserCheck(karitoken string) (string, string, string) {
	var err error
	DbConnection, err = sql.Open(config.Config.DBdriver, ConnectionInfo())

	if err != nil {
		log.Fatalln(err)
	}

	var id, uname, email, pass, token string
	err = DbConnection.QueryRow("SELECT * FROM kari_users WHERE token = $1", karitoken).Scan(&id, &uname, &pass, &email, &token)
	if err != nil {
		log.Println(err)
	}
	return uname, email, token
}

func KariTokenCheck(token string) bool {
	var err error
	DbConnection, err = sql.Open(config.Config.DBdriver, ConnectionInfo())

	if err != nil {
		log.Fatalln(err)
	}

	var Token string
	err = DbConnection.QueryRow("SELECT token FROM kari_users WHERE token = $1", token).Scan(&Token)
	if err != nil {
		log.Println(err)
		return true
	} else {
		return false
	}
}

func GetKariUserALL(token string) (string, string, string) {
	var err error
	DbConnection, err = sql.Open(config.Config.DBdriver, ConnectionInfo())

	if err != nil {
		log.Fatalln(err)
	}

	var id, uname, email, pass, Token string
	err = DbConnection.QueryRow("SELECT * FROM kari_users WHERE token = $1", token).Scan(&id, &uname, &pass, &email, &Token)
	if err != nil {
		log.Println("GetKariUserALL", err)
	}
	return uname, email, pass
}

func DeleteKariUser(token string) {
	var err error
	DbConnection, err = sql.Open(config.Config.DBdriver, ConnectionInfo())
	defer DbConnection.Close()

	if err != nil {
		log.Fatalln(err)
	}

	cmd, err := DbConnection.Prepare("DELETE FROM kari_users WHERE token = $1")
	if err != nil {
		log.Println(err)
	}
	_, err = cmd.Exec(token)
	if err != nil {
		log.Println(err)
	}
}
