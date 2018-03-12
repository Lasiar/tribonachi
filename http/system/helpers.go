package system

import (
	"fmt"
	"for_job/http/db"
	"strconv"
)

func Trib(n uint32) (string, error) {
	if n <= 76 {
		return tribOnLocal(int(n)), nil
	}

	if !db.Check() {
		return "", fmt.Errorf("[%v] %v", "Redis", "unavailable")
	}
	str := db.Cash.Get(strconv.FormatUint(uint64(n), 10)).Val()
	if str == "0" {
		return "", fmt.Errorf("[%v] %v", "Redis", "not counted yet")
	}
	return str, nil

}
