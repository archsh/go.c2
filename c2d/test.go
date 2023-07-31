package main

import (
	"encoding/xml"
	"fmt"
	c2 "github.com/archsh/go.c2"
	"github.com/spf13/cobra"
	"os"
)

// newCmd represents the version command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running tests...")
		var adi = c2.ADI{
			Objects: []c2.Object{
				c2.NewObject(c2.PROGRAM, "123", "REGIST"),
				c2.NewObject(c2.MOVIE, "123", "REGIST"),
				c2.NewObject(c2.PICTURE, "123", "REGIST"),
			},
			Mappings: []c2.Mapping{
				{ID: "111", ParentType: "A", ParentID: "1", ElementType: "B", ElementID: "2", Action: "Register", Properties: []c2.Property{{Name: "AAA", Value: "BBB"}, {Name: "XXX", Value: "YYY"}}},
				{ID: "111", ParentType: "A", ParentID: "1", ElementType: "B", ElementID: "2", Action: "Register", Properties: []c2.Property{{Name: "AAA", Value: "BBB"}, {Name: "XXX", Value: "YYY"}}},
				{ID: "111", ParentType: "A", ParentID: "1", ElementType: "B", ElementID: "2", Action: "Register", Properties: []c2.Property{{Name: "AAA", Value: "BBB"}, {Name: "XXX", Value: "YYY"}}},
			},
		}
		if bs, e := xml.MarshalIndent(adi, "", "  "); nil != e {
			fmt.Println(e)
		} else {
			_, _ = os.Stdout.WriteString(xml.Header)
			_, _ = os.Stdout.Write(bs)
		}
	},
}

func init() {

}
