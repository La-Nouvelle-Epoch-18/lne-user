package start

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/auth"
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/driver/postgres"
	server "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/http"
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/store"

	apiv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/api/v1"
	authv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/service/auth/v1"
	userv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/service/user/v1"
)

var (
	port string
	addr string
	pgc  = &postgres.Config{}
)

func init() {
	Cmd.Flags().StringVarP(&port, "port", "p", "9900", "port")
	Cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1", "")

	Cmd.Flags().StringVar(&pgc.Name, "pg-name", "nuitdelinfo", "")
	Cmd.Flags().StringVar(&pgc.Hostname, "pg-hostname", "localhost", "")
	Cmd.Flags().StringVar(&pgc.User, "pg-user", "root", "")
	Cmd.Flags().StringVar(&pgc.Password, "pg-password", "root", "")
	Cmd.Flags().StringVar(&pgc.Port, "pg-port", "5432", "")
}

var Cmd = &cobra.Command{
	Use: "start http server",
	Run: startServer,
}

func startServer(*cobra.Command, []string) {
	engine, err := postgres.New(pgc)
	if err != nil {
		log.Fatalln(err)
	}

	dbStore, err := store.New(engine)
	if err != nil {
		log.Fatalln(err)
	}

	err = dbStore.Sync()
	if err != nil {
		log.Fatalln(err)
	}

	authOp := auth.NewOperator(dbStore)

	userService := userv1.NewService(dbStore, authOp)
	authService := authv1.NewService(dbStore, authOp)

	api := apiv1.New(authService, userService)
	server := server.NewServer(api, fmt.Sprintf("%s:%s", addr, port))

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
