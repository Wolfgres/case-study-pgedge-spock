/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"financial_app/internal/services"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "financial_app",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.financial_app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	getConfigFilePath()
	if cfgFile != "" {
		// Use config file from the flag.
		dir, file := path.Split(cfgFile)
		viper.AddConfigPath(dir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(file)

	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".AsturDB" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//logrus.Debug("Using config file:", viper.ConfigFileUsed())
		logrus.Info("Using config file:", viper.ConfigFileUsed())
	} else {
		logrus.Fatal("Could not read configuration file")
		os.Exit(1)
	}

	//fmt.Println(viper.Get("log.log_format"))
	if viper.Get("log.log_format") == "text" {
		logrus.SetFormatter(&prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		})
	}

	exePath := services.GetExecPath()
	logPath := exePath + services.PathSeparator + filepath.Dir(fmt.Sprintf("%v", viper.Get("log.log_file")))
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		logrus.Error(err)
		logrus.Info("Create log directory in ", logPath)
		err := os.Mkdir(logPath, 0755)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	if viper.Get("log.log_format") == "json" {
		// Log as JSON instead of the default ASCII formatter.
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	logFile := exePath + services.PathSeparator + fmt.Sprintf("%v", viper.Get("log.log_file"))
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, f)
	logrus.SetOutput(mw)

	// Only log the warning severity or above.
	logLevel := services.GetLogLevel(fmt.Sprintf("%v", viper.Get("log.log_level")))
	logrus.SetLevel(logLevel)
	//logrus.SetLevel(logrus.InfoLevel)
	//fmt.Println(viper.Get("logrus.log_format"))
}

func getConfigFilePath() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Current project path not found", err)
	}
	fmt.Println("Current project path:", dir)
	cfgFile = dir + "/config/config.yaml"
}
