package data

import (
	"context"
	"csv2mongo/configs"
	"csv2mongo/models"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddData() (string, error) {
	clientOptions := options.Client().ApplyURI(configs.EnvMongoURI())
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return "", err
	}
	defer client.Disconnect(context.Background())

	db := client.Database("local")
	collection := db.Collection("weconnect")

	csvFile, err := os.Open("./data/data.csv")
	if err != nil {
		return "", err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	dataChan := make(chan interface{})
	errChan := make(chan error)
	done := make(chan bool)

	go func() {
		for entry := range dataChan {
			_, err := collection.InsertOne(context.Background(), entry)
			if err != nil {
				errChan <- err
				return
			}
		}
		done <- true
	}()

	go func() {
		defer close(dataChan)
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				errChan <- err
				return
			}

			dataValue, _ := strconv.ParseFloat(line[2], 64)
			magnitude, _ := strconv.Atoi(line[6])

			entry := models.Data{
				SeriesReference: line[0],
				Period:          line[1],
				DataValue:       dataValue,
				Suppressed:      line[3],
				Status:          line[4],
				Units:           line[5],
				Magnitude:       magnitude,
				Subject:         line[7],
				Group:           line[8],
				SeriesTitle1:    line[9],
				SeriesTitle2:    line[10],
				SeriesTitle3:    line[11],
				SeriesTitle4:    line[12],
				SeriesTitle5:    line[13],
			}

			dataChan <- entry
		}
	}()

	go func() {
		for err := range errChan {
			log.Println("Error:", err)
		}
	}()

	<-done

	return "CSV data successfully imported to MongoDB.", nil
}
