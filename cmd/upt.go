/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zze326/zpm/core"
	"github.com/zze326/zpm/util"
	"log"
	"strconv"
	"strings"
)

var addr string
var desc string

// uptCmd represents the upt command
var uptCmd = &cobra.Command{
	Use:   "upt",
	Args:  cobra.ExactArgs(1),
	Short: "Update a record",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("The id must be a number.")
		}
		var count int
		err = core.Db.QueryRow("SELECT count(1) FROM main WHERE id=?", id).Scan(&count)
		util.FatalErr(err)

		if count == 0 {
			log.Fatalln("The record does not exist")
		}

		if !util.StringIsEmpty(addr) {
			stmt, err := core.Db.Prepare("update main set addr = ? where id = ?;")
			_, err = stmt.Exec(addr, id)
			util.FatalErr(err)
		}

		if !util.StringIsEmpty(desc) {
			stmt, err := core.Db.Prepare("update main set desc = ? where id = ?;")
			_, err = stmt.Exec(desc, id)
			util.FatalErr(err)
		}

		if !util.StringIsEmpty(pwd) {
			stmt, err := core.Db.Prepare("update main set pwd = ? where id = ?;")
			encryptedPwd, err := util.EncryptPassword(strings.TrimSpace(pwd), core.Key)
			util.FatalErr(err)
			_, err = stmt.Exec(encryptedPwd, id)
			util.FatalErr(err)
		}

		fmt.Println("Updated successfully!")
	},
}

func init() {
	uptCmd.Flags().StringVarP(&addr, "addr", "a", "", "addr")
	uptCmd.Flags().StringVarP(&desc, "desc", "d", "", "desc")
	uptCmd.Flags().StringVarP(&pwd, "pwd", "p", "", "password")
	uptCmd.MarkFlagsOneRequired("addr", "desc", "pwd")
	rootCmd.AddCommand(uptCmd)
}
