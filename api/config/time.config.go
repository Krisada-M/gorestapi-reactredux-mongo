package config

import (
	"strconv"
	"time"
)

var monthlist = [12]string{
	"มกราคม",
	"กุมภาพันธ์",
	"มีนาคม",
	"เมษายน",
	"พฤษภาคม",
	"มิถุนายน",
	"กรกฎาคม",
	"สิงหาคม",
	"กันยายน",
	"ตุลาคม",
	"พฤศจิกายน",
	"ธันวาคม",
}

var (
	tn    = time.Now()
	Day   = strconv.Itoa(tn.Day())
	Month = monthlist[tn.Month()-1]
	Year  = strconv.Itoa(tn.Year() + 543)
	Time  = tn.Format("15:04:05")
)
