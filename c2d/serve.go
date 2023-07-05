package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
)

// newCmd represents the version command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run application in serve mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To be done...")
		listenAddr, _ := cmd.Flags().GetString("listen")
		//log.Fatalln(service.Listen(listenAddr))
		app := fiber.App{}

		log.Fatalln(app.Listen(listenAddr))
	},
}

func init() {
	serveCmd.Flags().StringP("listen", "L", "127.0.0.1:8080", "Demo application listen address and port.")
	//serveCmd.Flags().StringP("config", "C", "", "Configuration filename.")
	// rootCmd.AddCommand(newCmd)
}
