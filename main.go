package main

import (
	// "net/http"

	"os"

	"github.com/jordanorelli/blammo"
	"github.com/spf13/cobra"
)

var cmd = cobra.Command{
	Use:   "fetch-gists",
	Short: "fetch-gists gets all of the gists for a user",
	Args:  cobra.ExactArgs(1),
	RunE:  root,
}

var options struct {
	token string
}

const accept = "application/vnd.github.v3+json"

func openLog(name string) *blammo.Log {
	stdout := blammo.NewLineWriter(os.Stdout)
	stderr := blammo.NewLineWriter(os.Stderr)
	return blammo.NewLog(name, blammo.InfoWriter(stdout), blammo.ErrorWriter(stderr))
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func root(cmd *cobra.Command, args []string) error {
	log := openLog("fetch-gists")

	c := client{
		Log:   log,
		token: options.token,
	}

	for page := 0; true; page++ {
		gists, err := c.gists(page)
		if err != nil {
			return err
		}

		if len(gists) == 0 {
			break
		}

		for _, g := range gists {
			log.Info("%s %s", g.Created, g.ID)
			if err := g.clone(); err != nil {
				log.Error(err.Error())
			}
		}
	}

	return nil
}

func init() {
	cmd.PersistentFlags().StringVarP(&options.token, "token", "t", "", "github API oauth token (required)")
	cmd.MarkPersistentFlagRequired("token")
}
