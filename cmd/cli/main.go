package main

import (
	"fmt"
	"os"

	"github.com/mycroft/go-cli-template/commands"
	"github.com/mycroft/go-cli-template/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello-world",
	Short: "This tool prints hello-world",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.HelloWorld())
	},
}

func init() {
	rootCmd.AddCommand(commands.EchoCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
