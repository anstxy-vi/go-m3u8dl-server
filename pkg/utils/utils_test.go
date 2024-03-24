package utils

import "testing"

func TestEmailCheck(t *testing.T) {
	arr := []string{
		"静", "安静", "超安静", "超级安静", "l", "lv", "lvw", "12345", "1234@gmail.com", "13@outlook.com", "123@hotmail.com", "123@qq.com", "123@yahoo.com", "123@163.com",
	}
	// for _, email := range arr {

	// 	err := EmailCheck(email)
	// 	if err != nil {
	// 		println(email, err.Error())
	// 	}
	// }

	for _, str := range arr {
		result := HandleEmailMask(str)
		println(result)
	}

}
