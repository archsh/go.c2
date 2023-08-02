package main

import (
	"encoding/xml"
	"fmt"
	c2 "github.com/archsh/go.c2"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// newCmd represents the version command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
}

var testADICmd = &cobra.Command{
	Use:   "adi",
	Short: "ADI file parse",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running tests...")
		for _, s := range args {
			var adi c2.ADI
			fp, e := os.Open(s)
			if nil != e {
				fmt.Println("Open file ("+s+") failed !", e)
				continue
			}
			if bs, e := io.ReadAll(fp); nil != e {
				fmt.Println("Read file ("+s+") failed !", e)
			} else if e := xml.Unmarshal(bs, &adi); nil != e {
				fmt.Println("Unmarshal file ("+s+") failed !", e)
			} else {
				fmt.Println("Read ADI from file ", s)
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
			_ = fp.Close()
		}
	},
}

var testFtpGetCmd = &cobra.Command{
	Use:   "ftp",
	Short: "FTP download file",
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range args {
			if n, e := c2.FtpGet(s, os.Stdout); nil != e {
				fmt.Println("ERROR:>", e)
			} else {
				fmt.Println("Downloaded ", n, " bytes")
			}
		}
	},
}

func init() {
	testCmd.AddCommand(testADICmd, testFtpGetCmd)
}
