package main

import (
	"./Messages"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getGroupFileCount(groupName string, sql *sql.DB) (int, error) {
	stmt, _ := sql.Prepare("SELECT COUNT(*) FROM Files WHERE groupName=?")
	rows, execErr := stmt.Query(groupName)
	if execErr != nil {
		return -1, execErr
	}
	if rows.Next() {
		var count int
		rows.Scan(&count)
		return count, nil
	} else {
		return -1, errors.New("No count was returned")
	}
}

func addFileInfoToDatabase(groupName string, id string, name string, sql *sql.DB) error {
	stmt, _ := sql.Prepare("INSERT INTO Files VALUES (?, ?, ?)")
	_, execErr := stmt.Exec(groupName, id, name)
	if execErr != nil {
		return execErr
	}
	return nil
}

func getFilePathFromInfo(groupName string, fileID string) string {
	return "./data/" + base64.StdEncoding.EncodeToString([]byte(groupName)) + "/" + fileID
}

func UploadHandler(message *Message, sql *sql.DB) {
	fileMsg := Messages.FileMessageReq{}
	proto.Unmarshal(message.body, &fileMsg)
	groupName := *message.client.groupName
	fileCount, _ := getGroupFileCount(groupName, sql)
	fileID := strconv.Itoa(fileCount + 1)
	addFileInfoToDatabase(groupName, fileID, fileMsg.Name, sql)
	path := getFilePathFromInfo(groupName, fileID)
	go func() {
		folderPath := path[:(strings.LastIndex(path, "/"))]
		mkDirErr := os.MkdirAll(folderPath, 0644)
		if mkDirErr != nil {
			fmt.Println("MkdirError: ", mkDirErr)
			return
		}
		saveErr := ioutil.WriteFile(path, fileMsg.Contents, 0644)
		if saveErr != nil {
			fmt.Println("SaveError: ", saveErr)
			return
		}
		addGroupMessage("Uploaded File " + fileID + " (" + fileMsg.Name + ")", groupName, *message.client.username, sql)
	}()
}

func doesFileInfoExist(groupName string, fileID string, sql *sql.DB) (bool, error) {
	stmt, _ := sql.Prepare("SELECT COUNT(*) FROM Files WHERE groupName=? AND id=?")
	rows, execErr := stmt.Query(groupName, fileID)
	if execErr != nil {
		return false, execErr
	}
	if rows.Next() {
		var count int
		rows.Scan(&count)
		return count == 1, nil
	} else {
		return false, errors.New("No count was returned")
	}
}

func DownloadHandler(message *Message, sql *sql.DB) {
	go func() {
		downloadMsg := Messages.DownloadReq{}
		proto.Unmarshal(message.body, &downloadMsg)
		groupName := *message.client.groupName
		fileID := downloadMsg.FileID
		exists, _ := doesFileInfoExist(groupName, fileID, sql)
		if exists {
			filePath := getFilePathFromInfo(groupName, fileID)
			fileData, err := ioutil.ReadFile(filePath)
			if err == nil {
				downloadRespMsg := Messages.DownloadResp{
					FileID: fileID,
					Contents: fileData,
				}
				respData, _ := proto.Marshal(&downloadRespMsg)
				message.client.send("downloadResp", respData)
			} else {
				sendError("downloadErr", err, message.client)
			}
		} else {
			sendError("downloadErr", errors.New("FileInfo does not exist"), message.client)
		}
	}()
}
