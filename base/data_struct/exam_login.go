package data_struct

import "github.com/satori/go.uuid"

// 登录时的参数
type ExamLoginStruct struct {
	Student_id string
	Exam_topic_id string
	Uuid uuid.UUID
}