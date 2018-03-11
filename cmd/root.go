package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{

}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Coudn't execute command", err)
		os.Exit(-1)
	}
}
