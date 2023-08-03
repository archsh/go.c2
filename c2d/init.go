package main

import (
	"database/sql"
	xql "github.com/archsh/go.xql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize database etc...",
	Run: func(cmd *cobra.Command, args []string) {
		var pgconf = DefaultPgConfig()
		var db *sql.DB
		if viper.InConfig("database") {
			if e := viper.UnmarshalKey("database", &pgconf); nil != e {
				log.Fatalln("Read database config failed:", e)
			}
		}
		if d, e := ConnectSQL(pgconf.Host, pgconf.Port, pgconf.Username, pgconf.Password, pgconf.DBName, pgconf.SSLMode); nil != e {
			log.Fatalln("Connect database config failed:", e)
		} else {
			log.Println("Connected Database: ", pgconf)
			db = d
		}
		session := xql.MakeSession(db, "postgres")
		for _, t := range tables {
			if e := session.Create(t); nil != e {
				log.Fatalln("Create table '"+t.TableName()+"' failed:", e)
			} else {
				log.Println("Created table:", t.TableName())
			}
		}

	},
}

func init() {
}
