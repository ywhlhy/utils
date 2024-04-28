package cus_time

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestFmtNowTime(t *testing.T) {
	fmtTimeArr := [3]string{"yyyyMMdd", "yyyyMMdd HH:mm:ss", "yyyy-MM-dd HH:mm:ss"}
	for _, fmtTimeStr := range fmtTimeArr {
		if _, ok := fmtMap[fmtTimeStr]; !ok {
			t.Errorf("未找到%s对应的日期格式", fmtTimeStr)
		}

		timeStr := FmtNowTime(fmtTimeStr)

		if len(timeStr) != len(fmtTimeStr) {
			t.Errorf("格式化长度[%d]不对", len(timeStr))
		}

		timelocal = time.FixedZone("CST", 3600*8)

		curTime := time.Now().In(timelocal)
		year := strconv.Itoa(curTime.Year())
		month := fmt.Sprintf("%02d", int(curTime.Month()))
		day := fmt.Sprintf("%02d", curTime.Day())
		hour := fmt.Sprintf("%02d", curTime.Hour())
		minute := fmt.Sprintf("%02d", curTime.Minute())
		second := fmt.Sprintf("%2d", curTime.Second())

		switch fmtTimeStr {
		case "yyyyMMdd":
			if year != timeStr[0:4] || month != timeStr[4:6] || day != timeStr[6:8] {
				t.Errorf("%s类型格式化失败", fmtTimeStr)
			}
		case "yyyyMMdd HH:mm:ss":
			if year != timeStr[0:4] || month != timeStr[4:6] || day != timeStr[6:8] ||
				hour != timeStr[9:11] || minute != timeStr[12:14] || second != timeStr[15:17] {
				t.Errorf("%s类型格式化失败", fmtTimeStr)
			}
		case "yyyy-MM-dd HH:mm:ss":
			if year != timeStr[0:4] || month != timeStr[5:7] || day != timeStr[8:10] ||
				hour != timeStr[11:13] || minute != timeStr[14:16] || second != timeStr[17:19] {
				t.Errorf("%s类型格式化失败", fmtTimeStr)
			}
		default:

		}

	}

}
