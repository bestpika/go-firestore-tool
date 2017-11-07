package hack

import (
	"log"
)

// PrintError 印出錯誤
func PrintError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// PrintErrorExit 印出錯誤並退出
func PrintErrorExit(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
