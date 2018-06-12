package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

const dateFormat = "2006-01-02"

var releaseDate string

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Bump version on changelog",
	Long:  "Change current Unreleased version into the new version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readChangelog(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		version := chg.Version{
			Name: args[0],
			Date: releaseDate,
		}

		changelog := parser.Parse(input)
		changelog.Release(version)
		changelog.Render(os.Stdout)
	},
}

func init() {
	today := time.Now().Format(dateFormat)
	bumpCmd.Flags().StringVarP(&releaseDate, "release-date", "d", today, "")
	rootCmd.AddCommand(bumpCmd)
}
