package control

import (
	"graduation_project_socket/base/data_struct"
	"graduation_project_socket/app/index/service"
	"graduation_project_socket/extend/wbskt"
	"encoding/json"
)

type ExamControl struct {
	examService service.ExamService
	returnStruct data_struct.ReturnStruct
}

// 登录
func (examControl ExamControl) ExamLoginStatus(con *wbskt.Connection, data []byte) {
	// 考试登录结构体
	var examLoginStruct = data_struct.ExamLoginStruct{}
	json.Unmarshal(data, &examLoginStruct) // 将数据转为json
	examLoginStruct.Uuid = con.Id
	// 开始处理
	res := examControl.examService.ExamLoginStatus(examLoginStruct)
	// 处理结果
	if res {
		examControl.returnStruct = data_struct.ReturnStruct{"exam_login", 200, "连接成功", ""}
	} else {
		examControl.returnStruct = data_struct.ReturnStruct{"exam_login", 400, "连接失败，刷新重试", ""}
	}
	// 转为json
	returnData,_ := json.Marshal(examControl.returnStruct)
	// 返回数据
	con.WriteMessage(returnData)
}

// 退出
func (examControl ExamControl) ExamUnLoginStatus(con *wbskt.Connection) {
	examControl.examService.ExamUnLoginStatus(con.Id)
}


