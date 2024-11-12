package main

import "net"

type User struct {
	Name string
	Addr string
    cl chan string
	conn net.Conn
}

//创建一个用户API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:  userAddr,
        Addr:  userAddr,
        cl: make(chan string),
        conn: conn,
    }
	return user
}

// 监听发送
func (user *User) ListenMsg(){
	for {
		msg := <-user.cl
		user.conn.Write([]byte(msg + "\r\n"))
	}
}
