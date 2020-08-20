package utils

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/proto"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/repl"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func NewRequest(opCode uint8) *repl.Packet {
	request := repl.NewPacket()
	request.ReqID = proto.GenerateRequestID()
	request.Opcode = opCode
	return request
}

func DoRequest(request *repl.Packet, addr string, timeoutSec int) (err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("close tcp connection fail, host(%v) error(%v)", addr, err))
		}
	}()

	err = request.WriteToConn(conn)
	if err != nil {
		err = fmt.Errorf("write to host(%v) error(%v)", addr, err)
		return
	}

	err = request.ReadFromConn(conn, timeoutSec) // read the response
	if err != nil {
		err = fmt.Errorf("read from host(%v) error(%v)", addr, err)
		return
	}

	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func HttpGetRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}