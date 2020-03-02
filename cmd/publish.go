package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/jaby/tabgo/tableau"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var tablDocument string
var tablServerURL string
var tablUsername string
var tablPassword string
var tablSite string
var tablProjectName string

var tablTargetConnections string

type ExampleConnectionFinder struct {
	connections map[string]tableau.Connection
}

func (cf ExampleConnectionFinder) FindConnection(caption string) (tableau.Connection, error) {
	if connection, found := cf.connections[caption]; found {
		return connection, nil
	}
	return tableau.Connection{}, fmt.Errorf("no target connection found for caption '%s'", caption)
}

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes a datasource (file-extension: tds or tdsx) or a workbook (twb or twbx) to tableau ",
	Run: func(cmd *cobra.Command, args []string) {
		tabl := tableau.TabGo{ServerURL: tablServerURL, ApiVersion: "3.6"}
		err := tabl.Signin(tablUsername, tablPassword, tablSite)
		if err != nil {
			log.Fatalf("unable to signin, error: %+v", err)
		}

		// create a ConnectionFinder cfr tableau.ConnectionFinder interface
		var connections map[string]tableau.Connection
		targetConnectionsContent, err := ioutil.ReadFile(tablTargetConnections)

		err = json.Unmarshal(targetConnectionsContent, &connections)
		if err != nil {
			log.Fatalf("can not read connections from %s, error: %+v", tablTargetConnections, err)
		}
		myConnectionFinder := ExampleConnectionFinder{connections: connections}

		startUpload := time.Now()
		log.Printf(">>>>  start upload %s ", tablDocument)
		_, err = tabl.PublishDocument(tablDocument, tablProjectName, myConnectionFinder)
		if err != nil {
			log.Fatalf("can not publish '%s' to project '%s' on site '%s',\nError: %+v ", tablDocument, tablProjectName, tablSite, err)
		}
		log.Printf(">>>>  upload of %s took: %s", tablDocument, time.Now().Sub(startUpload))

		err = tabl.Signout()
		if err != nil {
			log.Fatalf("unable to signout")
		}
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

	publishCmd.Flags().StringVarP(&tablTargetConnections, "targetConnections", "t", "", "reference to target connections json file")
	publishCmd.MarkFlagRequired("targetConnections")
}
