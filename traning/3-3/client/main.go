package main

import (
	"log"
	"net/rpc"
)

type (
	GetUserReq struct {
		Id string `json:"id"`
	}
	GetUserResp struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dial error")
	}
	defer client.Close()
	var (
		req = GetUserReq{
			Id: "1",
		}
		resp GetUserResp
	)

	err = client.Call("UserServer.GetUser", req, &resp)
	if err != nil {
		log.Fatalf("call error:%v", err)
		return
	}
	log.Printf("resp: %+v", resp)
}
