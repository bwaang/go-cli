package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"strings"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "prints all arguments",
	Long: `Creates a comma separated print of arguments:

print a b c
> a,b,c`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, ","))
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
