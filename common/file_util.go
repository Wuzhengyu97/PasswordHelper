package common

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ImportFile2String(path string) (string, error) {
	// 读取file
	file, fileError := os.Open(path)
	defer file.Close()
	if fileError != nil {
		log.Fatalf("can not open the file, err is %+v", fileError)
		return "", fileError
	}
	fileReader := bufio.NewReader(file)
	var s string
	var buf []byte = make([]byte, 1000)
	for {
		n, readerError := fileReader.Read(buf)
		//hookfn(buf[:n])
		if readerError == io.EOF {
			break
		}
		s += string(buf[:n])
	}

	//fmt.Println(s)
	return s, nil
}

func File2Map(path string) map[string]interface{} {
	s, _ := ImportFile2String(path)
	return String2json(s)
}