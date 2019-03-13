package service

import (
	"graduation_project_socket/base/data_struct"
	"graduation_project_socket/base"
	"github.com/satori/go.uuid"
)

type ExamService struct {

}

// 登录service
func (examService ExamService) ExamLoginStatus(data data_struct.ExamLoginStruct) bool {
	var db = base.Db{}
	res :=db.Up("update student_exam_topic set login_status=?, uuid=? where student_id=? and exam_topic_id=? ",
		1,data.Uuid,data.Student_id,data.Exam_topic_id)
	if res > 0 {
		return true
	}
	return false
}

// 退出service
func (examService ExamService) ExamUnLoginStatus(uuid uuid.UUID) {
	var db = base.Db{}
	db.Up("update student_exam_topic set login_status=?,uuid=? where uuid=?",
		0,"",uuid)
}

