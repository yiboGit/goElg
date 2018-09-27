package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/olivere/elastic"
)

var (
	esClient, esErr = elastic.NewClient(elastic.SetURL("http://elastic:changeme@101.132.73.244:9200"),
		elastic.SetSniff(false))
	ctx = context.Background()
)

func main() {
	e := echo.New()
	e.POST("/remove/log/record", func(c echo.Context) error {
		var dateRange map[string]string
		c.Bind(&dateRange)
		err := removeRecord(dateRange["start"], dateRange["end"])
		if err != nil {
			return c.JSON(500, (fmt.Sprintf("%s \n", err.Error())))
		}
		return c.JSON(200, "delete ok")
	})
	e.Logger.Panic(e.Start(":23030"))
}

func removeRecord(start, end string) error {
	sDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		log.Printf("string time format err: %v\n", err)
		return err
	}
	eDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		log.Printf("string time format err: %v\n", err)
		return err
	}

	date := eDate
	for !date.Before(sDate) {
		index := fmt.Sprintf("docker-%v", date.Format("2006-01-02"))
		exits, err := esClient.IndexExists(index).Do(ctx)
		if err != nil {
			return err
		}
		if exits {
			_, err := esClient.DeleteIndex(index).Do(ctx)
			if err != nil {
				return err
			}
		}
		date = date.AddDate(0, 0, -1)
	}
	return nil
}
