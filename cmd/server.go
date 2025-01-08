package main

import (
	"fmt"
)

func intToAlpha(n int) string {
	// For uppercase letters (A-Z)
	return string(rune('A' + n - 1))
}

func main() {
	fmt.Println("Your server code goes here.")
	// sls := skiplist.NewSkiplist()
	// for i := 1; i <= 8; i++ {
	// 	sls.Insert(intToAlpha(i*3), strconv.Itoa(i*3))
	// }
	// sls.Print()
	// sst, err := sstable.NewSSTable(sls, "/tmp/data/sstables/0001.sst", "/tmp/data/indexes/001.idx")
	// if err != nil {
	// 	fmt.Printf("Unable to create sstable: %v", err)
	// }
	// if err = sst.Write(); err != nil {
	// 	fmt.Printf("Unable to write data: %v", err)
	// }
	// sst.Close()
}
