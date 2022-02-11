package tools

import "strconv"

func ParseFloat(data string) (int, error) {
	num, err := strconv.ParseFloat(data, 0)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}
