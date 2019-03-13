package tools

import (
	_ "graduation_project_socket/base/inter"
)

type DataConvert struct {
	bData []byte
}

// 字节转字符串
func (dataConverInter DataConvert) TypeToString() string {
	return string(dataConverInter.data)
}