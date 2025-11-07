package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	insecure bool
)

var rootCmd = &cobra.Command{
	Use:   "gurl",
	Short: "Go analog of cURL",
	Long: `

gURL was created for including it in docker image. Basically I just don't want to
install garbage in my alpine image. So we've decided to write cURL analog.`,
	Example: `
gurl get <url> // perform GET request
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "k", false, "Skip certificate verification")
}
