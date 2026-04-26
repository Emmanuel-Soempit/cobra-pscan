/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github/Emmanuel-Soempit/cobra-pscan/scan"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:          "delete <host1>....<hostn>",
	Short:        "Delete host(s) from the host list",
	Aliases:      []string{"d"},
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile := viper.GetString("hosts-file")
		return deleteAction(os.Stdout, hostsFile, args)
	},
}

func deleteAction(out io.Writer, hostsFile string, args []string) error {

	hl := &scan.HostsList{}

	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	for _, host := range args {

		if err := hl.Remove(host); err != nil {
			return err
		}

		fmt.Fprintf(out, "Deleted host: %s\n\n", host)
	}

	return hl.Save(hostsFile)
}

func init() {
	hostsCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
