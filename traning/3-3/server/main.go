package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type UserServer struct {
}
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

// UserServer.GetUser
func (*UserServer) GetUser(req GetUserReq, resp *GetUserResp) error {
	if u, ok := users[req.Id]; ok {
		*resp = GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return nil
	}
	return errors.New("user not found")
}
func main() {
	userServer := new(UserServer)
	// 注册服务到rpc服务
	rpc.Register(userServer)
	//监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error")
	}
	defer listen.Close()
	log.Print("listen on 1234")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("accept error")
			continue
		}
		go rpc.ServeConn(conn)
	}

}
