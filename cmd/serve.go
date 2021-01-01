package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/enbis/gofinance/internal/app/handler"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var wait time.Duration
		wait = 30 * time.Second
		// Routers declarations
		r := mux.NewRouter()
		api := r.PathPrefix("/api/v1/").Subrouter()

		// Routes handling
		api.HandleFunc("/", handler.Request).Methods("GET")
		api.HandleFunc("/info", handler.GoFinInfo).Methods("GET")

		log.Println("listening on port 3000")

		srv := &http.Server{
			Addr: "0.0.0.0:3000",
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      r, // Pass our instance of gorilla/mux in.
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		// Block until we receive our signal.
		<-c

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()
		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		srv.Shutdown(ctx)
		// Optionally, you could run srv.Shutdown in a goroutine and block on
		// <-ctx.Done() if your application should wait for other services
		// to finalize based on context cancellation.
		log.Println("shutting down")
		os.Exit(0)

	},
}
