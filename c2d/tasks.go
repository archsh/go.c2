package main

import (
	"database/sql"
	c2 "github.com/archsh/go.c2"
	xql "github.com/archsh/go.xql"
	log "github.com/sirupsen/logrus"
	"time"
)

func makeADITaskHandler(db *sql.DB, c2conf C2Config) func() {
	var f = func() {
		log.Infoln("Starting ADI task ...")
		for {
			session := xql.MakeSession(db, "postgres", true)
			var filters = []interface{}{
				xql.Where("status", 0, "="),
			}
			if n, e := session.Table(CmdRequestTable).Filter(filters...).Count(); nil != e {
				log.Errorln("Count failed:", e)
				session.Close()
				time.Sleep(time.Minute)
			} else if n < 1 {
				time.Sleep(time.Second * 15)
			} else {
				var req ExecRequest
				if e := session.Table(CmdRequestTable).Filter(filters...).OrderBy("created").Limit(1).One().Scan(&req); nil != e {
					log.Errorln("Select failed:", e)
					session.Close()
					continue
				}
				if e := session.Begin(); nil != e {
					log.Errorln("Session begin:", e)
					session.Close()
					continue
				}
				var up = make(map[string]interface{}, 3)
				if adi, e := c2.FTPGetADI(req.CmdFileURL); nil != e {
					log.Errorln("Get ADI failed:>", e)
					up["status"] = -1
					up["result"] = e.Error()
				} else if e := saveADI(session, adi); nil != e {
					log.Errorln("Save ADI failed:>", e)
					up["status"] = -2
					up["result"] = e.Error()
				} else {
					up["status"] = 1
				}
				up["updated"] = time.Now()
				_, _ = session.Table(CmdRequestTable).Where("correlate_id", req.CorrelateID).Update(up)
				if e := session.Commit(); nil != e {
					log.Errorln("Session begin:", e)
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
	log.Infoln("Saving ADI:>", adi.BizDomain, adi.CheckFlag, adi.StaffID)
	for _, obj := range adi.Objects {
		log.Infoln("Processing Object:>", obj.Action, obj.ElementType, obj.ID, obj.Code)
		switch obj.Action {
		case c2.REGIST: // Register
			var c2obj = C2Object{
				ID:          obj.ID,
				Code:        obj.Code,
				ElementType: obj.ElementType,
				Properties:  c2.L2M(obj.Properties),
			}
			if n, e := session.Table(ObjectTable).Where("id", c2obj.ID).Count("id"); nil != e {
				log.Errorln("Count Object failed:>", e)
				return e
			} else if n < 1 {
				if n, e := session.Table(ObjectTable).Insert(&c2obj); nil != e {
					log.Errorln("Insert Object failed:>", e)
					return e
				} else {
					log.Infoln("Inserted Object:>", c2obj.ElementType, c2obj.ID, c2obj.Code, n)
				}
			}
		case c2.UPDATE: // Update
			var c2obj C2Object
			if e := session.Table(ObjectTable).Get(obj.ID).Scan(&c2obj); nil != e {
				log.Errorln("Get Object failed:>", e)
				return e
			} else {
				var up = make(map[string]interface{}, 3)
				up["updated"] = time.Now()
				up["props"] = c2.L2M(obj.Properties)
				up["sync"] = 0
				if n, e := session.Table(ObjectTable).Where("id", obj.ID).Where("type", obj.ElementType).Update(up); nil != e {
					log.Errorln("Update Object failed:>", e)
					return e
				} else {
					log.Infoln("Updated Object:>", c2obj.ElementType, c2obj.ID, c2obj.Code, n)
				}
			}
		case c2.DELETE: // Delete
			var c2obj C2Object
			if e := session.Table(ObjectTable).Get(obj.ID).Scan(&c2obj); nil != e {
				log.Errorln("Get Object failed:>", e)
				return e
			} else {
				var up = make(map[string]interface{}, 3)
				up["updated"] = time.Now()
				up["status"] = -1
				up["sync"] = 0
				if n, e := session.Table(ObjectTable).Where("id", obj.ID).Where("type", obj.ElementType).Update(up); nil != e {
					log.Errorln("Delete Object failed:>", e)
					return e
				} else {
					log.Infoln("Delete Object:>", c2obj.ElementType, c2obj.ID, c2obj.Code, n)
				}
			}
		}
	}

	for _, m := range adi.Mappings {
		log.Infoln("Processing Mapping:>", m.Action, m.ElementType, m.ElementID, m.ElementCode)
		switch m.Action {
		case c2.REGIST, c2.UPDATE:
			var c2obj C2Object
			if e := session.Table(ObjectTable).Get(m.ElementID).Scan(&c2obj); nil != e {
				log.Errorln("Get Object failed:>", e)
				return e
			}
			var up = make(map[string]interface{})
			up["updated"] = time.Now()
			up["parent_type"] = m.ParentType
			up["parent_id"] = m.ParentID
			up["parent_code"] = m.ParentCode
			for _, p := range m.Properties {
				switch p.Name {
				case "Type":
					up["map_type"] = p.Value
				case "Sequence":
					up["sequence"] = p.Value
				case "ValidStart":
					up["valid_start"] = p.Value
				case "ValidEnd":
					up["valid_end"] = p.Value
				}
			}
			if n, e := session.Table(ObjectTable).Where("id", m.ElementID).Where("type", m.ElementType).Update(up); nil != e {
				log.Errorln("Mapping Object failed:>", e)
				return e
			} else {
				log.Infoln("Mapping Object:>", c2obj.ElementType, c2obj.ID, c2obj.Code, n)
			}
		case c2.DELETE:
			var c2obj C2Object
			if e := session.Table(ObjectTable).Get(m.ElementID).Scan(&c2obj); nil != e {
				log.Errorln("Get Object failed:>", e)
				return e
			}
			var up = make(map[string]interface{})
			up["updated"] = time.Now()
			up["parent_type"] = ""
			up["parent_id"] = ""
			up["parent_code"] = ""
			up["map_type"] = -1
			up["sequence"] = -1
			up["valid_start"] = nil
			up["valid_end"] = nil
			if n, e := session.Table(ObjectTable).Where("id", m.ElementID).Where("type", m.ElementType).Update(up); nil != e {
				log.Errorln("Un Mapping Object failed:>", e)
				return e
			} else {
				log.Infoln("Un Mapping Object:>", c2obj.ElementType, c2obj.ID, c2obj.Code, n)
			}
		}
	}

	return nil
}
