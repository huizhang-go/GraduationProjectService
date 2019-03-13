package app

import (
	"graduation_project_socket/extend/wbskt"
	"encoding/json"
	"graduation_project_socket/base/data_struct"
	"graduation_project_socket/app/index/control"
)

type Event struct {

}

// 连接
func (event Event) OnConn(con *wbskt.Connection)  {

}

// 接收消息
func (event Event) OnMsg(con *wbskt.Connection, data []byte) {
	var curlStruct data_struct.CurlStruct
	json.Unmarshal(data, &curlStruct)
	switch curlStruct.Curl_type {
		case "exam_login": // 更改登录状态
			var examControl control.ExamControl = control.ExamControl{}
			examControl.ExamLoginStatus(con, data)
		break
	}
}

// 关闭连接
func (event Event) OnClose(con *wbskt.Connection) {
	var examControl control.ExamControl = control.ExamControl{}
	examControl.ExamUnLoginStatus(con)
}