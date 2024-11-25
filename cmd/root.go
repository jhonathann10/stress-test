/*
Copyright © 2024 NAME HERE @jhonathann10
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jhonathann10/stress-test/internal/infra/client"
	"github.com/jhonathann10/stress-test/internal/usecase"
	"github.com/spf13/cobra"
)

type Variables struct {
	URL         string
	Requests    int
	Concurrency int
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Stress-Test para a pós da Full Cycle.",
	Long:  `Serviço para testar a capacidade de carga de um servidor.`,
	Run: func(cmd *cobra.Command, args []string) {
		var v Variables
		v.URL, _ = cmd.Flags().GetString("url")
		v.Requests, _ = cmd.Flags().GetInt("requests")
		v.Concurrency, _ = cmd.Flags().GetInt("concurrency")

		if !v.containsVariables() {
			cmd.Help()
			return
		}

		newClient := client.NewClient(v.URL)
		startRequests := usecase.NewStartRequests(v.Requests, v.Concurrency, newClient)
		err := startRequests.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("url", "u", "", "URL to send the request")
	rootCmd.PersistentFlags().IntP("requests", "r", 0, "Number of requests to send")
	rootCmd.PersistentFlags().IntP("concurrency", "c", 0, "Number of concurrent requests to send")
}

func (v *Variables) containsVariables() bool {
	if v.URL == "" {
		fmt.Println("URL is required")
		return false
	}
	if !strings.Contains(v.URL, "http") {
		fmt.Println("URL must be a valid URL, starting with http or https")
		return false
	}
	if v.Requests == 0 {
		fmt.Println("Requests is required")
		return false
	}
	if v.Concurrency == 0 {
		fmt.Println("Concurrency is required")
		return false
	}

	return true
}
