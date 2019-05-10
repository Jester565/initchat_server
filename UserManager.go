package main

import (
	"database/sql"
	"errors"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/bcrypt"
	"log"
	"./Messages"
)


func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Hash Failed")
	}
	return string(hash)
}

func createAccount(username string, password string, sql *sql.DB) error {
	pwdHash := hashPassword(password)
	stmt, prepErr := sql.Prepare("INSERT INTO Users VALUES (?, ?)")
	if prepErr != nil {
		return prepErr
	}
	_, execErr := stmt.Exec(username, pwdHash)
	if  execErr != nil {
		return execErr
	}
	return nil
}

func CreateAccountHandler(message *Message, sql *sql.DB) {
	go func() {
		signUpMsg := Messages.SignUpReq{}
		proto.Unmarshal(message.body, &signUpMsg)
		err := createAccount(signUpMsg.Username, signUpMsg.Password, sql)
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, _ := proto.Marshal(&errorMsg)
			message.client.send("authErr", errorData)
			return
		}
		authMsg := Messages.AuthResp{}
		authData, _ := proto.Marshal(&authMsg)
		message.client.username = &signUpMsg.Username
		message.client.send("auth", authData)
	}()
}


func doesPasswordMatchHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func login(username string, password string, sql *sql.DB) error {
	stmt, _ := sql.Prepare("SELECT pwdHash FROM Users WHERE name=?")
	rows, execErr := stmt.Query(username)
	if execErr != nil {
		return execErr
	}
	if rows.Next() {
		var pwdHash string
		rows.Scan(&pwdHash)
		if doesPasswordMatchHash(password, pwdHash) {
			return nil
		} else {
			return errors.New("invalid password")
		}
	} else {
		return errors.New("username not found")
	}
}

func LoginHandler(message *Message, sql *sql.DB) {
	go func() {
		loginMsg := Messages.LoginReq{}
		proto.Unmarshal(message.body, &loginMsg)
		err := login(loginMsg.Username, loginMsg.Password, sql)
		if err != nil {
			sendError("authErr", err, message.client)
			return
		}
		authMsg := Messages.AuthResp{}
		authData, _ := proto.Marshal(&authMsg)
		message.client.username = &loginMsg.Username
		message.client.send("auth", authData)
	}()
}

func searchUsers(usernamePrefix string, username string, sql *sql.DB) ([]string, error) {
	usernamePrefix += "%"
	stmt, _ := sql.Prepare("SELECT name FROM Users WHERE LCASE(name) LIKE ? AND name!=? ORDER BY LENGTH(name), name LIMIT 50")
	rows, execErr := stmt.Query(usernamePrefix, username)
	if execErr != nil {
		return nil, execErr
	}
	var usernames []string
	for rows.Next() {
		var searchName string
		rows.Scan(&searchName)
		usernames = append(usernames, searchName)
	}
	return usernames, nil
}

func SearchHandler(message *Message, sql *sql.DB) {
	go func() {
		searchMsg := Messages.UserSearchReq{}
		proto.Unmarshal(message.body, &searchMsg)
		usernames, err := searchUsers(searchMsg.UsernamePrefix, *message.client.username, sql)
		if err != nil {
			sendError("userSearchErr", err, message.client)
			return
		}
		usernamesMsg := Messages.UserSearchResp{ Usernames: usernames }
		usernamesData, _ := proto.Marshal(&usernamesMsg)
		message.client.send("userSearchResp", usernamesData)
	}()
}
