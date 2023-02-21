package commonfunction

import (
	"strconv"
	"time"
)

func ParseStringToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func ParseStringToInt(s string) (int, error) {
	f, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return f, nil
}
func ParseStringToBool(s string) (bool, error) {
	f, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return f, nil
}
func ParseStringToInt64(s string) (int64, error) {
	f, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}
func ParseStringToUint64(s string) (uint64, error) {
	f, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}
func ParseStringToByte(s string) ([]byte, error) {
	f, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return []byte{byte(f)}, nil
}
func CheckEroor(err error) bool {
	if err != nil {
		return true
	}
	return false
}
func CheckEroorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
func CheckEroorInt(err error) int {
	if err != nil {
		return 0
	}
	return 1
}
func CheckEroorInt64(err error) int64 {
	if err != nil {
		return 0
	}
	return 1
}
func CheckEroorFloat64(err error) float64 {
	if err != nil {
		return 0
	}
	return 1
}
func ParseStringToTime(s string,format string) (time.Time, error) {
	f, err := time.Parse(format, s)
	if err != nil {
		return time.Time{}, err
	}
	return f, nil
}
func ParseStringToTimeLayout(s string, layout string) (time.Time, error) {
	f, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, err
	}
	return f, nil
}
func ParseTimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}

