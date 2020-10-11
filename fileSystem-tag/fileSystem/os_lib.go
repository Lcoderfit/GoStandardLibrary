package fileSystem

import (
	"fmt"
	"log"
	"os"
)

func UseOs() {
	osOpen()
}

//打开文件操作
func osOpen() {
	path := "E:/SocialProject/Learn-Tags/GoStandardLibrary/fileSystem-tag/go.mod"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	//打印的是file变量的地址
	fmt.Println(file)

	data := make([]byte, 100)
	//count是data的字节数
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	//如果直接用data而不是用data[:count], 则末尾会输出很多\x00
	fmt.Printf("read %d bytes: %q", count, data[:count])
}
