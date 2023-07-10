package commands

import (
	"io"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wisnuragaprawida/project/bootstrap"
	"github.com/wisnuragaprawida/project/pkg/log"
	"github.com/wisnuragaprawida/project/pkg/log/logrusw"
)

type commandFn func(dep *bootstrap.Dependency) *cobra.Command

var subCommands []commandFn

func registerCommand(fn commandFn) {
	subCommands = append(subCommands, fn)
}

func Run(dep *bootstrap.Dependency) error {
	var (
		cpu          int
		config       string
		verbose      bool
		tracerCloser io.Closer
	)
	rootCmd := &cobra.Command{
		Use: "project",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			//set logger setting
			lr := logrus.New()
			lr.SetFormatter(&logrus.JSONFormatter{})
			log.SetLogger(&logrusw.Logger{Logger: lr})
			if verbose {
				lr.Level = logrus.DebugLevel
			}

			//set number of cpu

			x := runtime.GOMAXPROCS(cpu)
			if cpu == 0 {
				cpu = x
			}
			log.Debugf("using %v cpu", cpu)

			//set configuration file
			log.Debugf("load config from %v", config)
			cfg, err := bootstrap.LoadConfig(config)
			if err != nil {
				log.Errorf("unable to load config file: %s : %s", config, err)
				os.Exit(1)
			}

			dep.SetConfig(cfg)
			if err := dep.Initialize(); err != nil {
				log.Errorf("unable to initialize dependency: %v", err)

			}

		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {

			if tracerCloser != nil {
				tracerCloser.Close()
			}

		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().IntVar(&cpu, "cpu", 0, "set number of cpu to use, default is number of cpu available")
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "config.yaml", "set configuration file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "set verbose mode")

	for _, fn := range subCommands {
		rootCmd.AddCommand(fn(dep))
	}

	return rootCmd.Execute()
}
