package cmd

import (
	"RepoHound/internal/output"
	"RepoHound/internal/scanner"
	"log"

	"github.com/spf13/cobra"
)

const version = "v0.1"

var rootCmd = &cobra.Command{
	Use:   "RepoHound",
	Short: "Discovery of public repos for list of users",
	Long:  `Discovery of public repos for list of users`,
	Run: func(cmd *cobra.Command, args []string) {
		// cmd.Root().Help()

		// get params
		fu, err := cmd.Flags().GetString("userlist")
		if err != nil {
			log.Fatalln(err)
		}
		proxy, err := cmd.Flags().GetString("proxy")
		if err != nil {
			log.Fatalln(err)
		}

		workers, err := cmd.Flags().GetInt("workers")
		if err != nil {
			log.Fatalln(err)
		}
		json, err := cmd.Flags().GetBool("json")
		if err != nil {
			log.Fatalln(err)
		}

		s, err := scanner.New(workers, fu, proxy)
		if err != nil {
			log.Fatalln(err)
		}

		results := s.Run()
		output.Print(results, json)
	},
}

func init() {
	// rootCmd.PersistentFlags().BoolP("verbose", "V", false, "Verbose")
	rootCmd.PersistentFlags().StringP("userlist", "u", "", "userlist")
	rootCmd.MarkFlagRequired("userlist")
	rootCmd.PersistentFlags().StringP("proxy", "p", "", "HTTP Proxy")
	rootCmd.PersistentFlags().BoolP("json", "", false, "Output in JSON Format")
	rootCmd.PersistentFlags().IntP("workers", "", 5, "Workers")
	// rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}

}
