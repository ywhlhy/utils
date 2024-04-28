package cus_time

import (
	"time"
)

var (
	fmtMap    = map[string]string{"yyyyMMdd": "20060102", "yyyyMMdd HH:mm:ss": "20060102 15:04:05", "yyyy-MM-dd HH:mm:ss": "2006-01-02 15:04:05"}
	timelocal *time.Location
)

func init() {
	timelocal = time.FixedZone("CST", 3600*8) // 东八
}

// FmtNowTime 按照预定的格式化字符串格式化当前时间
func FmtNowTime(fmtStr string) string {
	fmtTime, ok := fmtMap[fmtStr]
	if !ok {
		fmtTime = "2006-01-02 15:04:05"
	}
	localTime := time.Now()

	return localTime.In(timelocal).Format(fmtTime)
}
