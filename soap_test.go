package c2

import (
	"fmt"
	"os"
	"testing"
)

var client *SOAPClient

func TestMain(m *testing.M) {
	var retCode int

	client = NewSOAPClient("http://localhost:8080/soap", false, nil)
	retCode = m.Run()

	os.Exit(retCode)
}

func TestSOAPClient_Test(t *testing.T) {
	var req = ExecCmdReq{
		CSPID:       "1",
		LSPID:       "2",
		CorrelateID: "3",
		CmdFileURL:  "http://12314.12414.12412.4/aaa",
	}
	var ret = ExecCmdRes{}
	if e := client.Call("ExecCmdReq", req, &ret); nil != e {
		fmt.Println("Failed:>", e)
	}
}
