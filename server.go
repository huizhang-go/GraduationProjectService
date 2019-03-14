package main

import (
	"github.com/gorilla/websocket"
	"graduation_project_socket/app"
	"graduation_project_socket/extend/wbskt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 监听系统信号
func getSignal(conn *wbskt.Connection, event app.Event, sigs chan os.Signal) {
	sig := <-sigs
	switch sig {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT: // 监听kill、退出信号
		event.OnClose(conn) // 关闭处理
		os.Exit(1)
		break
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
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
	sigs := make(chan os.Signal, 1)
	if conn, err = wbskt.InitConnection(wsConn); err != nil {
		goto ERR
	}

	// 监听退出信号
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go getSignal(conn, event, sigs)

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

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
