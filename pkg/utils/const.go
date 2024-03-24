package utils

const (
	//时间format格式
	TIME_FORMAT = "2006-01-02 15:04:05"
	//时间format格式
	TIME_FORMAT_PARSE = "2006-01-02T15:04:05.999999999Z07:00"
	TIME_FORMAT_NUM   = "20060102150405"

	//email 正则
	EMAIL_VERIFY_PATTERN = "^[a-zA-Z0-9]+([-_.][a-zA-Z0-9]+)*@[a-zA-Z0-9]+([-_.][a-zA-Z0-9]+)*\\.[a-z]{2,}$+"
	//phone 正则
	PHONE_VERIFY_PATTERN = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|(19[0-9])|165|166|(147))\\d{8}$"
	//允许的email 正则
	ALLOW_EMAIL_PATTERN = "^[a-zA-Z0-9]+([-_.][a-zA-Z0-9]+)*@(yahoo|gmail|outlook|hotmail|icloud|gmx|qq|163)(\\.[a-z]{2,}){1,2}$"
)
