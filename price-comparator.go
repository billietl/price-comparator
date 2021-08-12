package main

import (
	"log"
	"os"
	"price-comparator/dao"
	"price-comparator/model"

	cli "github.com/urfave/cli/v2"
)

var (
	Version string
	daoType = "firestore"
)

func run(c *cli.Context) error {
	ctx := c.Context
	usedDao, err := dao.GetDAOBundle(ctx, daoType)
	if err != nil {
		log.Fatal("Error creating DAO of type \"" + daoType + "\" : " + err.Error())
		return err
	}
	product := model.NewProduct("test", false)
	product, err = usedDao.ProductDAO.Upsert(ctx, product)
	if err != nil {
		log.Fatal("Error creating product " + err.Error())
		return err
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
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
