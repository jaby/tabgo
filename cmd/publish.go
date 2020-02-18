package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tablDocument string
var tablServerURL string
var tablUsername string
var tablPassword string
var tablSite string
var tablProjectName string

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes a datasource (file-extension: tds or tdsx) or a workbook (twb or twbx) to tableau ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publish called")
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmd.Flags().StringVarP(&tablDocument, "document", "d", "", "tableau document to publish, should have file-extension *.tds(x) for datasource or *twb(x) for workbook")
	publishCmd.MarkFlagRequired("document")

	publishCmd.Flags().StringVarP(&tablServerURL, "url", "u", "", "tableau server URL")
	publishCmd.MarkFlagRequired("url")

	publishCmd.Flags().StringVarP(&tablUsername, "username", "n", "", "tableau username")
	publishCmd.MarkFlagRequired("username")

	publishCmd.Flags().StringVarP(&tablPassword, "password", "x", "", "tableau password")
	publishCmd.MarkFlagRequired("password")

	publishCmd.Flags().StringVarP(&tablSite, "site", "s", "", "tableau site")
	publishCmd.MarkFlagRequired("site")

	publishCmd.Flags().StringVarP(&tablProjectName, "project", "p", "", "tableau project within site")
	publishCmd.MarkFlagRequired("project")
}
