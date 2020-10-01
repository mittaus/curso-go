package main

import (
	"fmt"
	"os"

	"example.com/mittaus/ddd-example/application"
	"example.com/mittaus/ddd-example/config"
	articleValidator "example.com/mittaus/ddd-example/infraestructure/dummy.articleValidator"
	server "example.com/mittaus/ddd-example/infraestructure/gin.server"
	slugger "example.com/mittaus/ddd-example/infraestructure/gosimple.slugger"
	jwt "example.com/mittaus/ddd-example/infraestructure/jwt.authHandler"
	logger "example.com/mittaus/ddd-example/infraestructure/logrus.logger"
	articleRW "example.com/mittaus/ddd-example/infraestructure/memory.articleRW"
	commentRW "example.com/mittaus/ddd-example/infraestructure/memory.commentRW"
	tagsRW "example.com/mittaus/ddd-example/infraestructure/memory.tagsRW"
	userRW "example.com/mittaus/ddd-example/infraestructure/memory.userRW"
	validator "example.com/mittaus/ddd-example/infraestructure/user.validator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// the command to run the server
var rootCmd = &cobra.Command{
	Use:   "go-realworld-clean",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\nargs:%s\n", Build, Version, args)
		run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\n", Build, Version)
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	cobra.OnInitialize(config.CobraInitialization)

	config.LoggerConfig(rootCmd)
	config.ServerConfig(rootCmd)
	config.DatabaseConfig(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}
}

func run() {
	puerto := os.Getenv("server.port")
	fmt.Println("puerto", puerto)
	ginServer := config.NewServer(
		viper.GetInt("server.port"),
		config.DebugMode,
	)

	authHandler := jwt.New(viper.GetString("jwt.Salt"))

	routerLogger := logger.NewLogger("TEST",
		viper.GetString("log.level"),
		viper.GetString("log.format"),
	)

	server.NewRouterWithLogger(
		application.HandlerConstructor{
			Logger:           routerLogger,
			UserRW:           userRW.New(),
			ArticleRW:        articleRW.New(),
			UserValidator:    validator.New(),
			AuthHandler:      authHandler,
			Slugger:          slugger.New(),
			ArticleValidator: articleValidator.New(),
			TagsRW:           tagsRW.New(),
			CommentRW:        commentRW.New(),
		}.New(),
		authHandler,
		routerLogger,
	).SetRoutes(ginServer.Router)

	ginServer.Start()
}
