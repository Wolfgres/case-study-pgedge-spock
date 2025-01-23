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

var (
	numGoroutines   int
	testDuration    int
	maxConns        int
	numTransactions int
	operation       int
	milisecondValue int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "financial_app",
	Short: "Stress Test Application",
	Long: `
	This database stress testing application allows you to run insert, delete, update, and query scripts.
	
	Some of the application's features include specifying duration seconds, transactions, goroutines (threads), 
	and maximum client connection capacity to the database stress test.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
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

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().IntVarP(&numGoroutines, "goroutines", "g", 0, "Number of concurrent goroutines")
	rootCmd.Flags().IntVarP(&testDuration, "duration", "d", 0, "Test duration in seconds")
	rootCmd.Flags().IntVarP(&maxConns, "max-conns", "c", 0, "Maximum number of connections in the pool")
	rootCmd.Flags().IntVarP(&numTransactions, "transactions", "t", 0, "Number of transactions. Must be a number that are multiples of goroutine")
	rootCmd.Flags().IntVarP(&operation, "operation", "o", 0, "Choose a stress test transactions: INSERT=1, SELECT=2, UPDATE=3")
	rootCmd.Flags().IntVarP(&milisecondValue, "miliseconds", "m", 0, "Milliseconds that a transaction takes to execute")
}

func start() {
	logrus.Info("************************************************")
	logrus.Info("*** Financial App by Wolfgres - Postgres Enterprise ****")
	logrus.Info("************************************************")
	logrus.Infof("Number of Goroutines: %d", numGoroutines)
	logrus.Infof("Test Duration: %d seconds", testDuration)
	logrus.Infof("Max Connections in Pool: %d", maxConns)
	logrus.Infof("Transacions per table: %d", numTransactions)
	logrus.Infof("Milliseconds per transaction: %d", milisecondValue)
	logrus.Info("************************************************")
	services.StressTestNodes(numGoroutines, testDuration, numTransactions, operation, maxConns, milisecondValue)
}

func getConfigFilePath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Current project path not found", err)
	}
	fmt.Println("Current project path:", dir)
	return dir + "/config/config.yaml"
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfgFile := getConfigFilePath()
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
