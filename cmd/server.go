//nolint:gochecknoinits,gochecknoglobals //lint issue ignored
package cmd

import (
	"net/http"
	"fmt"

	"github.com/JpUnique/Verbatim/pkg/config"
	"github.com/JpUnique/Verbatim/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Bookstore Api",
	Long:  `An API used to Delete, Put, Get, Post books into the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, err := config.SetupLogger()
		if err != nil {
			return err
		}
		return StartServer(logger)
	},
}

func StartServer(logger *zap.Logger) error {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	addr := "localhost:9010"
	logger.Info("Starting server", zap.String("address", addr))
	err := http.ListenAndServe(addr, r)
	if err != nil {
		logger.Error("failed to start the server", zap.Error(err))
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
