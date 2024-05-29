package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var EchoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Print the given string passed in arguments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, " "))
	},
}
