package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"graduation_project_socket/extend/wbskt"
	"time"
	"graduation_project_socket/app"
)
var(
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter , r *http.Request) {
	//	w.Write([]byte("hello"))
	var (
		wsConn *websocket.Conn
		err    error
		conn   *wbskt.Connection
		data   []byte
		event  app.Event
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = wbskt.InitConnection(wsConn); err != nil {
		goto ERR
	}
	event.OnConn(conn) // 建立连接
	// 启动线程，发送心跳
	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(time.Duration(20) * time.Second)
		}
	}()

	// 循环从
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		event.OnMsg(conn, data)
	}
ERR:
	event.OnClose(conn)
}

func main(){
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe("0.0.0.0:7777",nil)
}