package main

import (
	"database/sql"
	"fmt"
	c2 "github.com/archsh/go.c2"
	xql "github.com/archsh/go.xql"
	"github.com/sirupsen/logrus"
	"time"
)

func makeTaskHandler(db *sql.DB) func() {
	var f = func() {
		for {
			session := xql.MakeSession(db, "postgres", true)
			if n, e := session.Table(CmdRequestTable).Where("status", 0, "=").Count(); nil != e {
				logrus.Errorln("Count failed:", e)
				session.Close()
				time.Sleep(time.Minute)
			} else if n < 1 {
				time.Sleep(time.Second * 15)
			} else {
				var req ExecCommandRequest
				if e := session.Table(CmdRequestTable).Where("status", 0, "=").OrderBy("created").Limit(1).One().Scan(&req); nil != e {
					logrus.Errorln("Select failed:", e)
					session.Close()
					continue
				}
				if e := session.Begin(); nil != e {
					logrus.Errorln("Session begin:", e)
					session.Close()
					continue
				}
				var up = make(map[string]interface{}, 3)
				if adi, e := c2.FTPGetADI(req.CmdFileURL); nil != e {
					logrus.Errorln("Get ADI failed:>", e)
					up["status"] = -1
					up["result"] = e.Error()
				} else if e := saveADI(session, adi); nil != e {
					logrus.Errorln("Save ADI failed:>", e)
					up["status"] = -2
					up["result"] = e.Error()
				} else {
					up["status"] = 1
				}
				up["updated"] = time.Now()
				_, _ = session.Table(CmdRequestTable).Where("correlate_id", req.CorrelateID).Update(up)
				if e := session.Commit(); nil != e {
					logrus.Errorln("Session begin:", e)
					_ = session.Rollback()
				}
				// Notify CMS SOAP ...
				session.Close()
			}
		}
	}
	return f
}

func saveADI(session *xql.Session, adi c2.ADI) error {
	logrus.Infoln("Saving ADI:>", adi.BizDomain, adi.CheckFlag, adi.StaffID)
	return fmt.Errorf("not implemented")
}
