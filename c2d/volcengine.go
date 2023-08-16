package main

import (
	"database/sql"
	"github.com/archsh/go.xql"
	"github.com/sirupsen/logrus"
	"github.com/volcengine/volcengine-sdk-go-rec/byteair"
	"github.com/volcengine/volcengine-sdk-go-rec/core"
	"github.com/volcengine/volcengine-sdk-go-rec/core/logs"
	"github.com/volcengine/volcengine-sdk-go-rec/core/metrics"
	"time"
)

type VolEngineConfig struct {
	TenantId      string
	ApplicationId string
	AK            string
	SK            string
}

func setupVolClient(vc VolEngineConfig) (byteair.Client, error) {
	if c, e := (&byteair.ClientBuilder{}).
		// 必传,租户id.
		TenantId(vc.TenantId).
		// 必传,应用id.
		ApplicationId(vc.ApplicationId).
		// 必传,密钥AK,获取方式:【火山引擎控制台】->【个人信息】->【密钥管理】中获取.
		AK(vc.AK).
		// 必传,密钥SK,获取方式:【火山引擎控制台】->【个人信息】->【密钥管理】中获取.
		SK(vc.SK).
		// 必传,国内使用RegionAirCn.
		Region(core.RegionAirCn).
		Build(); nil != e {
		return c, e
	} else {
		// metrics上报初始化.建议开启,方便火山侧排查问题.
		metrics.Init()
		// 默认log级别为warn.根据需要自行更改.
		logs.Level = logs.LevelInfo
		return c, e
	}
}

func makeSyncTaskHandler(client byteair.Client, db *sql.DB, c2conf C2Config) func() {
	var f = func() {
		for {
			session := xql.MakeSession(db, "postgres", true)
			var filters = []interface{}{
				xql.Where("sync", 0, "="),
			}
			if n, e := session.Table(ObjectTable).Filter(filters...).Count(); nil != e {
				logrus.Errorln("Count failed:", e)
				session.Close()
				time.Sleep(time.Minute)
			} else if n < 1 {
				time.Sleep(time.Second * 15)
			} else {
				var obj C2Object
				if e := session.Table(ObjectTable).Filter(filters...).OrderBy("created").Limit(1).One().Scan(&obj); nil != e {
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
				//client.WriteData()
				//client.Done()
				// sync ....
				//if adi, e := c2.FTPGetADI(req.CmdFileURL); nil != e {
				//	logrus.Errorln("Get ADI failed:>", e)
				//	up["status"] = -1
				//	up["result"] = e.Error()
				//} else if e := saveADI(session, adi); nil != e {
				//	logrus.Errorln("Save ADI failed:>", e)
				//	up["status"] = -2
				//	up["result"] = e.Error()
				//} else {
				//	up["status"] = 1
				//}
				up["updated"] = time.Now()
				_, _ = session.Table(ObjectTable).Where("id", obj.ID).Update(up)
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
