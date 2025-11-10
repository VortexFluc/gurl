package cmd

import (
	"fmt"
	"gurl/client"
	"log"

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

		r, err := client.Get(insecure, args[0])

		if err != nil {
			log.Fatal(err)
		}

		res, err := Format(r.Response, format)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", res)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
