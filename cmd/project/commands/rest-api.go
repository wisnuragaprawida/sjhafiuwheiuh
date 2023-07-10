package commands

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/wisnuragaprawida/project/bootstrap"
	"github.com/wisnuragaprawida/project/internal/api"
	"github.com/wisnuragaprawida/project/pkg/log"
)

func init() {
	registerCommand(startRestApiService)
}

func startRestApiService(dep *bootstrap.Dependency) *cobra.Command {
	return &cobra.Command{
		Use:   "rest-api",
		Short: "rest api service",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := dep.GetConfig()

			handler := api.NewServer(dep.GetDB(), api.ServerConfig{
				EncKey: cfg.EncKey,
			})

			actx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

			srv := http.Server{Addr: cfg.Host.Address, Handler: handler}

			errChan := make(chan error)

			defer close(errChan)
			go func() {
				log.Infof("Server is running on %s at %s env", cfg.Host.Address, os.Getenv("ENV"))
				err := srv.ListenAndServe()
				if err != nil {
					errChan <- errors.New("Serever error : " + err.Error())
				}

			}()

			select {
			case err := <-errChan:
				log.Error(err)
				return
			case <-actx.Done():
				err := srv.Shutdown(context.Background())
				if err != nil {
					log.Error(err)
					return
				}
			}
			log.Info("Server is shutdown")

		},
	}
}
