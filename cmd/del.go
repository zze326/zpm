/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/zze326/zpm/core"
	"github.com/zze326/zpm/util"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a record",
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

		stmt, err := core.Db.Prepare("delete from main where id = ?;")
		util.FatalErr(err)
		_, err = stmt.Exec(addr, id)
		util.FatalErr(err)

		fmt.Println("Deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
