package main

import (
	"database/sql"
	"fmt"
	c2 "github.com/archsh/go.c2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// newCmd represents the version command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run application in serve mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To be done...")
		var listenAddr = "127.0.0.1:8080"
		if s, b := cmd.Flags().GetString("listen"); nil != b && s != "" {
			listenAddr = s
		} else if ss := viper.GetString("listen"); ss != "" {
			listenAddr = ss
		}
		var c2conf C2Config
		if viper.InConfig("c2") {
			if e := viper.UnmarshalKey("c2", &c2conf); nil != e {
				log.Fatalln("Read c2 config failed:", e)
			}
		}
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
		go makeADITaskHandler(db, c2conf)()
		if viper.InConfig("volcengine") {
			var vc VolEngineConfig
			if e := viper.UnmarshalKey("", &vc); nil != e {
				log.Fatalln("Read volcengine config failed:", e)
			}
			if c, e := setupVolClient(vc); nil != e {
				log.Fatalln("Setup volcengine failed:", e)
			} else {
				go makeSyncTaskHandler(c, db, c2conf)()
			}
		}

		app := fiber.New()
		app.Use(func(ctx *fiber.Ctx) error {
			log.Println(ctx.Method(), "", ctx.Path())
			return ctx.Next()
		})
		app.Post("/", c2.MakeRequestCmdHandler(makeRequestProcessHandler(db, c2conf)))
		log.Fatalln(app.Listen(listenAddr))
	},
}

func init() {
	serveCmd.Flags().StringP("listen", "L", "", "Demo application listen address and port.")
	//serveCmd.Flags().StringP("config", "C", "", "Configuration filename.")
	// rootCmd.AddCommand(newCmd)
}
