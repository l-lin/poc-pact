package cmd

import (
	"fmt"
	"log"

	"github.com/l-lin/kitty-cli/list"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "GET a cat",
		Run:   runGet,
	}
	id int64
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().Int64VarP(&id, "id", "i", 0, "id of the cat to fetch")
	addCmd.MarkFlagRequired("id")
}

func runGet(cmd *cobra.Command, args []string) {
	s := list.Service{URL: providerURL + "/cats"}
	cat, err := s.Get(id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cat)
}
