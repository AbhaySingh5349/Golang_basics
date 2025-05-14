package main

import (
	"fmt"
	"time"
)

// storing timestamps in DB in Unix time simplifies sorting & querying data
// Epoc time is universal across platforms & programming languages

func main() {

	// 00:00:00 UTC on Jan 1, 1970

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time:", unixTime) // 1747233737

	t := time.Unix(unixTime, 0) // 0 represents "nanoseconds" for additional level of accuracy

	fmt.Println(t)                               // 2025-05-14 20:12:17 +0530 IST
	fmt.Println("Time:", t.Format("2006-01-02")) // 2025-05-14

}
