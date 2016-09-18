package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {

	_, fulleFilename, line, _ := runtime.Caller(0)
	fmt.Println(fulleFilename)
	fmt.Println(line)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fulleFilename)
	fmt.Println("filenameWithSuffix=", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	fmt.Println("fileSuffix=", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenameOnly=", filenameOnly)

	var logFilename string = filenameOnly + ".log"
	fmt.Println("logFilename=", logFilename)
	logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("open file error=%s\r\n", err.Error())
		os.Exit(-1)
	}

	defer logFile.Close()
	logger := log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	//logger := log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("normal log 1")
	logger.Println("normal log 2")

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8081", nil)

}
