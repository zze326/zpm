/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/zze326/zpm/core"
	"github.com/zze326/zpm/util"
	"os"
)

var showPwd bool
var pwd string

// lstCmd represents the lst command
var lstCmd = &cobra.Command{
	Use:   "lst",
	Short: "View all records",
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := core.Db.Query("SELECT id, addr, pwd, desc FROM main")
		util.FatalErr(err)

		table := tablewriter.NewWriter(os.Stdout)
		lineHeader := []string{"ID", "Addr", "Desc"}
		if showPwd {
			lineHeader = append(lineHeader, "Password")
		}
		table.SetHeader(lineHeader)
		for rows.Next() {
			var id int
			var addr string
			var desc string
			var encryptedPwd *string
			err = rows.Scan(&id, &addr, &encryptedPwd, &desc)
			util.FatalErr(err)
			line := []string{color.New(color.FgMagenta).Sprintf("%d", id), color.New(color.FgCyan, color.Bold).Sprint(addr), color.New(color.FgBlue).Sprint(desc)}
			if showPwd {
				if encryptedPwd == nil {
					line = append(line, color.New(color.FgHiWhite).Sprint("NULL"))
				} else {
					pwd, err := util.DecryptPassword(*encryptedPwd, core.Key)
					util.FatalErr(err)
					line = append(line, fmt.Sprintf("[%s]", color.New(color.FgRed).Sprint(string(pwd))))
				}
			}
			table.Append(line)
		}
		table.SetAlignment(tablewriter.ALIGN_CENTER)
		table.Render()
	},
}

func init() {
	lstCmd.Flags().BoolVarP(&showPwd, "show-password", "v", false, "Whether to display passwords")
	rootCmd.AddCommand(lstCmd)
}
