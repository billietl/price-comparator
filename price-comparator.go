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

func main() {
	app := cli.App{}
	app.Name = "Price comparator"
	app.Usage = "price-comparator service launcher"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "dao-type",
			Value:       daoType,
			Usage:       "Set the storage backend type to use for the service (only firestore available for now)",
			Destination: &daoType,
		},
	}

	app.Action = func(c *cli.Context) error {
		ctx := c.Context
		usedDao, err := dao.GetDAOBundle(daoType)
		if err != nil {
			log.Fatal("Error creating DAO of type \"" + daoType + "\" : " + err.Error())
		}
		product := &model.Product{
			Name: "test",
			Bio:  false,
		}
		product, err = usedDao.ProductDAO.Upsert(ctx, product)
		if err != nil {
			log.Fatal("Error creating product " + err.Error())
		}
		_, err = usedDao.ProductDAO.Load(ctx, product.ID)
		if err != nil {
			log.Fatal("Error fetching product " + err.Error())
		}
		err = usedDao.ProductDAO.Delete(ctx, product.ID)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
