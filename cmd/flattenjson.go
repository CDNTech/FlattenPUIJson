/*
Copyright © 2026 Patrick Buick <patrick.buick@powerfleet.com>
This file is part of a CLI application for PowerFleet.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// flattenjsonCmd represents the flattenjson command
var flattenjsonCmd = &cobra.Command{
	Use:   "flattenjson",
	Short: "A program to flatten the JSON that comes out of the queries to the PUI Posiverse Portal",
	Long: `See the short description, but
	Usage: flattenjson -filename <filename>
	It expects the JSON in the format coming out of the query to the PosiversePortal`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: I have decided this should be the generic version that flattens *ANY* JSON following the "standard"
		// TODO: So I have to track depth, starting tokens to know if I am in an object or an array and how many etc etc.
		if filename, err := cmd.Flags().GetString("filename"); err == nil {
			fmt.Printf("flattenjson called with filename: %s\n", filename)
		} else {
			log.Fatal(err)
		}
		// This needs to be a routine to have structures representing the file format coming out of the query I am using.
		// Then the job is to read in the file and print out the objects as flattened JSON for import to Excel etc.
	},
}

func init() {
	rootCmd.AddCommand(flattenjsonCmd)

	flattenjsonCmd.Flags().String("filename", "", "Provide a filename with JSON to flatten.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flattenjsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flattenjsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
