package main

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	LoginCmdDeviceName     string
	WatchCmdPriorityFilter string
)

var RootCmd = &cobra.Command{
	Use: "push-watch",
}

var LoginCmd = &cobra.Command{
	Use:  "login [flags] username password",
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		if err := login(args[0], args[1]); err != nil {
			log.Fatal(err)
		}
	},
}

var WatchCmd = &cobra.Command{
	Use:  "watch [flags] device-id device-secret [...command]",
	Args: cobra.MatchAll(cobra.MinimumNArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		if err := watch(args[0], args[1], args[2:]); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(LoginCmd)
	RootCmd.AddCommand(WatchCmd)
	LoginCmd.Flags().StringVarP(&LoginCmdDeviceName, "device-name", "n", "push-watch", "Device name")
	WatchCmd.Flags().StringVarP(&WatchCmdPriorityFilter, "priority", "p", "-2,-1,0,1,2", "Priority filter")
}

func main() {
	RootCmd.Execute()
}
