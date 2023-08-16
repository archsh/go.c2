package main

import (
	"database/sql"
	"fmt"
	c2 "github.com/archsh/go.c2"
	xql "github.com/archsh/go.xql"
	"github.com/sirupsen/logrus"
	"time"
)

func makeRequestProcessHandler(db *sql.DB, c2conf C2Config) c2.RequestCmdHandleFunc {
	var ff = func(CSPID, LSPID, CorrelateID, CmdFileURL string) error {
		var req = ExecRequest{
			CSPID:       CSPID,
			LSPID:       LSPID,
			CorrelateID: CorrelateID,
			CmdFileURL:  CmdFileURL,
		}
		if session := xql.MakeSession(db, "postgres", true); nil == session {
			logrus.Errorln("RequestCmdHandleFunc:> create session failed")
			return fmt.Errorf("create session failed")
		} else if n, e := session.Table(CmdRequestTable).Where("correlate_id", CorrelateID).Count(); nil != e {
			logrus.Errorln("RequestCmdHandleFunc:>", e)
			return e
		} else if n < 1 {
			if _, e := session.Table(CmdRequestTable).Insert(&req); nil != e {
				logrus.Errorln("RequestCmdHandleFunc:>", e)
				return e
			} else {
				return nil
			}
		} else {
			if _, ee := session.Table(CmdRequestTable).Where("correlate_id", CorrelateID).Update(map[string]interface{}{
				"status":  0,
				"updated": time.Now(),
			}); nil != ee {
				logrus.Errorln("RequestCmdHandleFunc:>", ee)
				return ee
			} else {
				return nil
			}
		}
	}
	return ff
}
