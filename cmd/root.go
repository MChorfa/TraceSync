/*
Copyright © 2024 TraceSync
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tracesync",
	Short: "TraceSync is a powerful CLI tool for managing AI artifacts, datasets, and SBOMs.",
	Long: `TraceSync streamlines the process of securely managing, validating, and tracking 
the transfer of AI models, datasets, and other critical outputs like SBOMs (Software Bill of Materials). 
It integrates with CI/CD pipelines, enforces metadata management, and maintains compliance with standards 
like SLSA and NIST. 

Example Usage:

  tracesync upload <artifact>      # Upload an artifact
  tracesync validate <dataset>     # Validate a dataset
  tracesync status <artifact>      # Check the status of a transfer

For more information, use 'tracesync [command] --help' to see detailed instructions on each command.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error executing command:", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Define persistent flags and configuration settings
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tracesync.yaml)")
	rootCmd.PersistentFlags().String("env", "staging", "Environment context (production, staging, etc.)")

	// Bind environment flag to Viper
	viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))

	// Define local flags specific to the root command
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in the config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the --config flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Locate the user's home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search for the config file in the home directory.
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".tracesync")
	}

	viper.AutomaticEnv() // Automatically read environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(os.Stderr, "Warning: No configuration file found.")
	}
}
