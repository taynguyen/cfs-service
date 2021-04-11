package main

import (
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	// dateTime, err := gostradamus.Parse("2020-11-25 07:36:04.000001", "YYYY-MM-DD HH:mm:ss.S")
	// if err != nil {
	// 	panic(err)
	// }

	// pp.Println("dateTime:", dateTime.GoString())

	layout := "2006-01-02 15:04:05.000"
	d, e := time.Parse(layout, "2020-11-25 07:36:04.193")
	pp.Println("d:", d.Format(layout), " e:", e, " mil:", d.Nanosecond())
	// time.Parse("", "")
	rs := time.Now().Format("2006-01-02 15:04:05.000")
	pp.Println("rs:", rs)
}
