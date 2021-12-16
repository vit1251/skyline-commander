package util

import (
	"fmt"
	"time"
)

func FormatTime(cur time.Time) string {

	/* Date */
	year := cur.Year()
	month := convertMonth2(cur.Month())
	day := cur.Day()

	/* Time */
	hour := cur.Hour()
	minute := cur.Minute()

	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d", year, month, day, hour, minute)

}

//func convertMonth(month time.Month) int {
//	switch month {
//	case time.January:
//		return 1
//	case time.February:
//		return 2
//	case time.March:
//		return 3
//	case time.April:
//		return 4
//	case time.May:
//		return 5
//	case time.July:
//		return 6
//	case time.June:
//		return 7
//	case time.August:
//		return 8
//	case time.September:
//		return 9
//	case time.October:
//		return 10
//	case time.November:
//		return 11
//	case time.December:
//		return 12
//	}
//	return -1
//}

func convertMonth2(month time.Month) int {
	return int(month)
}