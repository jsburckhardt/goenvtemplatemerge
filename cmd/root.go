package cmd

import (
	"fmt"
	"os"

	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"

	"github.com/spf13/cobra"
)

var cfgFile string
var debug bool

var rootCmd = &cobra.Command{
	Use:   "goenvtemplatemerge",
	Short: "Updates placeholders in template with environment variables",
	Long: `goenvtemplatemerge helps updating config/templates that contain place holders
base on the environment variables in the system.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cLog := console.New(true)
	if debug {
		log.AddHandler(cLog, log.AllLevels...)
	} else {
		log.AddHandler(cLog, []log.Level{
			log.InfoLevel,
			log.NoticeLevel,
			log.WarnLevel,
			log.ErrorLevel,
			log.PanicLevel,
			log.AlertLevel,
			log.FatalLevel,
		}...)
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "whether to show debug info")
}
