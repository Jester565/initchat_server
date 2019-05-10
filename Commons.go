package main

import (
	"./Messages"
	"github.com/golang/protobuf/proto")

func sendError(typeID string, err error, client *Client) {
	errorMsg := Messages.Error{Message: err.Error()}
	errorData, _ := proto.Marshal(&errorMsg)
	client.send(typeID, errorData)
	return
}