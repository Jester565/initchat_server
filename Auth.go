package main

import (
	"database/sql"
	"errors"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/bcrypt"
	"log"
	"./Messages"
)


//Passwords should be converted to secure salted hash
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Hash Failed")
	}
	return string(hash)
}

func doesPasswordMatchHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createAccount(username string, password string, sql *sql.DB) error {
	pwdHash := hashPassword(password)
	stmt, prepErr := sql.Prepare("INSERT INTO Users VALUES (?, ?)")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return prepErr
	}
	_, execErr := stmt.Exec(username, pwdHash)
	if  execErr != nil {
		log.Println("INSERT USER Err: ", execErr)
		return execErr
	}
	return nil
}

func login(username string, password string, sql *sql.DB) error {
	log.Println("LOGIN START")
	stmt, prepErr := sql.Prepare("SELECT pwdHash FROM Users WHERE name=?")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return prepErr
	}
	log.Println("LOGIN QUERY")
	rows, execErr := stmt.Query(username)
	if execErr != nil {
		log.Println("Login Query Error: ", execErr)
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

func createAccountHandler(message *Message, sql *sql.DB) {
	var createAccountRunner = func() {
		signUpMsg := Messages.SignUpReq{}
		parseErr := proto.Unmarshal(message.body, &signUpMsg)
		if parseErr != nil {
			log.Fatalln("PARSE ERROR: ", parseErr)
			return
		}
		err := createAccount(signUpMsg.Username, signUpMsg.Password, sql)
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("authError", errorData)
			return
		}
		authMsg := Messages.AuthResp{}
		authData, serializeErr := proto.Marshal(&authMsg)
		if serializeErr != nil {
			log.Fatalln("SERIALIZE ERROR: ", serializeErr)
			return
		}
		message.client.username = &signUpMsg.Username
		message.client.send("auth", authData)
	}
	go createAccountRunner()
}

func loginHandler(message *Message, sql *sql.DB) {
	var loginRunner = func() {
		log.Println("LOGIN RUNNER")
		loginMsg := Messages.LoginReq{}
		parseErr := proto.Unmarshal(message.body, &loginMsg)
		if parseErr != nil {
			log.Fatalln("PARSE ERROR: ", parseErr)
			return
		}
		err := login(loginMsg.Username, loginMsg.Password, sql)
		if err != nil {
			log.Println("Login Err: ", err)
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("authError", errorData)
			return
		}
		authMsg := Messages.AuthResp{}
		authData, serializeErr := proto.Marshal(&authMsg)
		if serializeErr != nil {
			log.Fatalln("SERIALIZE ERROR: ", serializeErr)
			return
		}
		message.client.username = &loginMsg.Username
		message.client.send("auth", authData)
	}
	log.Println("LOGIN HANDLER CALLED")
	go loginRunner()
}

func searchUsers(usernamePrefix string, username string, sql *sql.DB) ([]string, error) {
	usernamePrefix += "%"
	stmt, prepErr := sql.Prepare("SELECT name FROM Users WHERE LCASE(name) LIKE ? AND name!=? ORDER BY LENGTH(name), name LIMIT 50")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return nil, prepErr
	}
	rows, execErr := stmt.Query(usernamePrefix, username)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
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

func searchHandler(message *Message, sql *sql.DB) {
	var searchRunner = func() {
		searchMsg := Messages.UserSearchReq{}
		parseErr := proto.Unmarshal(message.body, &searchMsg)
		if parseErr != nil {
			log.Fatalln("PARSE ERROR: ", parseErr)
			return
		}
		usernames, err := searchUsers(searchMsg.UsernamePrefix, *message.client.username, sql)
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("userSearchErr", errorData)
			return
		}
		usernamesMsg := Messages.UserSearchResp{ Usernames: usernames }
		usernamesData, serializeErr := proto.Marshal(&usernamesMsg)
		if serializeErr != nil {
			log.Fatalln("SERIALIZE ERROR: ", serializeErr)
			return
		}
		message.client.send("userSearchResp", usernamesData)
	}
	go searchRunner()
}
