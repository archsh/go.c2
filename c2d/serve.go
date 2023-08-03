package main

import (
	"fmt"
	c2 "github.com/archsh/go.c2"
	xql "github.com/archsh/go.xql"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
)

var (
	CmdRequestTable = xql.DeclareTable(&ExecCommandRequest{})
)

// newCmd represents the version command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run application in serve mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To be done...")
		listenAddr, _ := cmd.Flags().GetString("listen")
		//log.Fatalln(service.Listen(listenAddr))
		app := fiber.New()
		app.Post("/soap/exec", c2.MakeRequestCmdHandler(func(CSPID, LSPID, CorrelateID, CmdFileURL string) error {
			fmt.Println("CSPID:>", CSPID)
			fmt.Println("LSPID:>", LSPID)
			fmt.Println("CorrelateID:>", CorrelateID)
			fmt.Println("CmdFileURL:>", CmdFileURL)
			if CmdFileURL != "" {
				if adi, e := c2.FTPGetADI(CmdFileURL); nil != e {
					fmt.Println("Read ADI failed:", e)
					return e
				} else {
					fmt.Println("Read ADI from CmdFileURL:", CmdFileURL)
					fmt.Println("> ", adi.BizDomain, adi.CheckFlag, adi.StaffID)
					for _, obj := range adi.Objects {
						fmt.Println(">> Object:", obj.ID, obj.ElementType, obj.Action)
						for _, prop := range obj.Properties {
							fmt.Println(">>>     ", prop.Name, "=", prop.Value)
						}
					}
					for _, m := range adi.Mappings {
						fmt.Println(">> Mapping:", m.ElementID, m.ElementCode, m.ElementType, m.ParentID, m.ParentCode, m.ParentType, m.Action)
						for _, prop := range m.Properties {
							fmt.Println(">>>     ", prop.Name, "=", prop.Value)
						}
					}
				}
			}
			return nil
		}))
		log.Fatalln(app.Listen(listenAddr))
	},
}

func init() {
	serveCmd.Flags().StringP("listen", "L", "127.0.0.1:8080", "Demo application listen address and port.")
	//serveCmd.Flags().StringP("config", "C", "", "Configuration filename.")
	// rootCmd.AddCommand(newCmd)
}
