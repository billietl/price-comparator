package main

import (
	"fmt"
	"os"
	"price-comparator/dao"
	"price-comparator/logger"
	"price-comparator/web"

	cli "github.com/urfave/cli/v2"
)

var (
	daoType      = "firestore"
	gcpProjectID = ""
)

func run(c *cli.Context) error {
	if gcpProjectID != "" {
		os.Setenv("GOOGLE_PROJECT_ID", gcpProjectID)
	}

	usedDao, err := dao.GetBundle(c.Context, daoType)
	if err != nil {
		logger.Error(err, fmt.Sprintf("Error creating DAO of type \"%s\"", daoType))
		return err
	}

	logger.InitLoggers(os.Stdout)

	server := web.MakeServer(8080, usedDao)
	err = server.Run()
	if err != nil {
		logger.Error(err, "Server crash")
	}

	usedDao.Shutdown()
	return nil
}

func main() {
	app := cli.App{}
	app.Name = "Price comparator"
	app.Usage = "price-comparator service launcher"
	app.Action = run
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "dao-type",
			Value:       daoType,
			Usage:       "Set the storage backend type to use for the service (only firestore available for now)",
			Destination: &daoType,
		},
		&cli.StringFlag{
			Name:        "gcp-project-id",
			Value:       gcpProjectID,
			Usage:       "Set the storage backend type to use for the service (only firestore available for now)",
			Destination: &gcpProjectID,
		},
	}
	app.Run(os.Args)
}
