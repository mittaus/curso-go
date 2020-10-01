package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CobraInitialization() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// viper.AddConfigPath("/config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No configuration file found")
	}
}

func LoggerConfig(rootCmd *cobra.Command) {
	var logLevel, logFormat string
	var isFound, logLine bool = false, false

	if logLevel, isFound = os.LookupEnv("server.host"); !isFound {
		logLevel = "info"
	}

	if logFormat, isFound = os.LookupEnv("server.host"); !isFound {
		logFormat = "text"
	}

	if myLogLine, myIsFound := os.LookupEnv("server.port"); myIsFound {
		line, error := strconv.ParseBool(myLogLine)
		if error == nil {
			logLine = line
		}
	}

	rootCmd.PersistentFlags().String("log.level", logLevel, "one of debug, info, warn, error or fatal")
	rootCmd.PersistentFlags().String("log.format", logFormat, "one of text or json")
	rootCmd.PersistentFlags().Bool("log.line", logLine, "enable filename and line in logs")

	viper.BindPFlags(rootCmd.PersistentFlags())
}

func ServerConfig(cmd *cobra.Command) {
	var serverHost, serverAllowedOrigins, serverToken, jwtSalt string
	var serverPort int = 8080
	var isFound, serverDebug bool = false, false

	if serverHost, isFound = os.LookupEnv("server.host"); !isFound {
		serverHost = "127.0.0.1"
	}

	if serverAllowedOrigins, isFound = os.LookupEnv("server.allowedOrigins"); !isFound {
		serverAllowedOrigins = "*"
	}

	if serverToken, isFound = os.LookupEnv("server.token"); !isFound {
		serverToken = ""
	}

	if jwtSalt, isFound = os.LookupEnv("jwt.salt"); !isFound {
		jwtSalt = ""
	}

	if myServerPort, myIsFound := os.LookupEnv("server.port"); myIsFound {
		port, error := strconv.Atoi(myServerPort)
		if error == nil {
			serverPort = port
		}
	}

	if myServerDebug, myIsFound := os.LookupEnv("server.debug"); !myIsFound {
		debug, error := strconv.ParseBool(myServerDebug)
		if error == nil {
			serverDebug = debug
		}
	}

	cmd.Flags().String("server.host", serverHost, "host on which the server should listen")
	cmd.Flags().Int("server.port", serverPort, "port on which the server should listen")
	cmd.Flags().Bool("server.debug", serverDebug, "debug mode for the server")
	cmd.Flags().String("server.allowedOrigins", serverAllowedOrigins, "allowed origins for the server")
	cmd.Flags().String("server.token", serverToken, "authorization token to use if any")
	cmd.Flags().String("jwt.salt", jwtSalt, "used to sign the JWTs")
	viper.BindPFlags(cmd.Flags())
}

func DatabaseConfig(cmd *cobra.Command) {
	var populate bool
	if myPopulate, myIsFound := os.LookupEnv("populate"); myIsFound {
		value, error := strconv.ParseBool(myPopulate)
		if error == nil {
			populate = value
		}
	}

	cmd.Flags().Bool("populate", populate, "used to populate databases in order to run integration tests")

	viper.BindPFlags(cmd.Flags())
}
