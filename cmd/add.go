/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zze326/zpm/core"
	"github.com/zze326/zpm/util"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(2),
	Short: "Add a new record",
	Run: func(cmd *cobra.Command, args []string) {
		if !util.StringIsEmpty(pwd) {
			stmt, err := core.Db.Prepare("INSERT INTO main(addr, pwd, desc) values(?,?,?);")
			util.FatalErr(err)
			encryptedPwd, err := util.EncryptPassword(strings.TrimSpace(pwd), core.Key)
			util.FatalErr(err)
			_, err = stmt.Exec(args[0], encryptedPwd, args[1])
			util.FatalErr(err)
		} else {
			stmt, err := core.Db.Prepare("INSERT INTO main(addr, desc) values(?,?);")
			util.FatalErr(err)
			_, err = stmt.Exec(args[0], args[1])
			util.FatalErr(err)
		}
		fmt.Println("Added successfully!")
	},
}

func init() {
	addCmd.Flags().StringVarP(&pwd, "pwd", "p", "", "password")
	rootCmd.AddCommand(addCmd)
}
