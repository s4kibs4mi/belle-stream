package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"github.com/s4kibs4mi/belle-stream/api"
	"fmt"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start streaming server",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Println("Serve has been started")
		if err := api.StartServer(); err != nil {
			fmt.Println("Failed to start server", err)
		}
	}()

	<-stop
	fmt.Println("Server shutting down gracefully...")
	api.StopServer()
	fmt.Println("Server shutted down gracefully...")
}
