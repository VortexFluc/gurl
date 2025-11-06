/*
Copyright © 2025 VFlux
*/

package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Perform a GET request",
	Long: `
Perform a GET request.
Returns non-zero status code if the request is failed.
`,
	Example: "gurl get https://google.com",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) == 0 {
			log.Fatal("No url passed")
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		}

		client := &http.Client{Transport: tr}

		r, err := client.Get(args[0])

		if err != nil {
			log.Fatal(err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(r.Body)

		body, err := io.ReadAll(r.Body)

		if r.StatusCode != http.StatusOK {
			log.Fatalf("Bad status code: %d.\nBody: %s", r.StatusCode, body)
		}

		fmt.Printf("%s\n", body)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
