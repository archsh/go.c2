package main

import (
	"database/sql"
	"fmt"
	c2 "github.com/archsh/go.c2"
	xql "github.com/archsh/go.xql"
)

func makeRequestProcessHandler(db *sql.DB) c2.RequestCmdHandleFunc {
	var ff = func(CSPID, LSPID, CorrelateID, CmdFileURL string) error {
		var req = ExecCommandRequest{
			CSPID:       CSPID,
			LSPID:       LSPID,
			CorrelateID: CorrelateID,
			CmdFileURL:  CmdFileURL,
		}
		if session := xql.MakeSession(db, "postgesql", true); nil == session {
			return fmt.Errorf("create session failed")
		} else if n, e := session.Table(CmdRequestTable).Where("CorrelateID", CorrelateID).Count(); nil != e {
			return e
		} else if n < 1 {
			if _, e := session.Table(CmdRequestTable).Insert(req); nil != e {
				return e
			} else {
				return nil
			}
		} else {
			_, ee := session.Table(CmdRequestTable).Where("CorrelateID", CorrelateID).Update(map[string]interface{}{"status": 0})
			return ee
		}
	}
	return ff
}
