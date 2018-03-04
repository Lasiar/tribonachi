package system

import (
	"for_job/http/db"
	"strconv"
)

func Trib(n int) string {
	if n < 100 {
		return tribOnLocal(n)
	} else {
		return db.Cash.Get(strconv.Itoa(n)).Val()
	}
}
