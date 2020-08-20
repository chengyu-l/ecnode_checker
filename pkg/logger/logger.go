package logger

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/utils"
	"os"
	"time"
)

const (
	outPath   = "out"
	closeByte = 1
)

type Logger struct {
	fileName  string
	filePath  string
	outFile   *os.File
	closeChan chan byte
}

func NewLogger(fileName string) *Logger {
	filePath := getOutFilePath(fileName)
	exists, _ := utils.PathExists(filePath)
	if exists {
		_ = os.Remove(filePath)
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("create logger file fail. err:%v", err))
		os.Exit(1)
	}

	closeChan := make(chan byte)
	closeOutFileAfter(outFile, closeChan)
	return &Logger{
		fileName:  fileName,
		filePath:  filePath,
		outFile:   outFile,
		closeChan: closeChan,
	}
}

func closeOutFileAfter(file *os.File, close chan byte) {
	select {
	case <-close:
		err := file.Sync()
		if err != nil {
			fmt.Println("sync logger file fail.")
		}
		err = file.Close()
		if err != nil {
			fmt.Println("close logger file fail.")
		}

		fmt.Println("close logger file")
	}
}

func (l *Logger) RecordF(format string, a ...interface{}) {
	l.Record(fmt.Sprintf(format, a))
}

func (l *Logger) Record(msg string) {
	_, err := l.outFile.WriteString(msg)
	if err != nil {
		fmt.Println("write to logger file fail", err)
	}
}

func (l *Logger) Close() {
	l.closeChan <- closeByte
}

func getOutFilePath(fileName string) string {
	path := fmt.Sprintf("%v/%v", outPath, time.Now().Format("2006010215"))
	exists, _ := utils.PathExists(path)
	if !exists {
		_ = os.MkdirAll(path, 0666)
	}

	return fmt.Sprintf("%v/%v", path, fileName)
}
