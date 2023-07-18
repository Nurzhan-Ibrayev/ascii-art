package Modify

import (
	"bufio"
	"log"
	"os"
)

func FromTxt() [][]string { // from standard.txt take all symbols and put them to [] string
	file, err := os.Open("banners/standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)
	var standardText []string
	for scanned.Scan() {
		if scanned.Text() == "" {
			continue
		}
		standardText = append(standardText, scanned.Text())
	}
	var res [][]string
	var temp []string
	for i, w := range standardText {
		i++
		if i%8 == 0 {
			temp = append(temp, w)
			res = append(res, temp)
			temp = nil
		} else {
			temp = append(temp, w)
		}

	}
	return res
}
