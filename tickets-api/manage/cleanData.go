package manage

import (
	"context"
	"encoding/json"
	"os"
	"ticket-api/models"
	"time"
)

var filename = "tickets.json"

func deleteData() {
	_, err := os.Stat(filename)
	var allData models.TicketAvailable

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			return
		}
		_ = json.Unmarshal(data, &allData)

		allData.Concert = allData.Concert[:0]
		allData.Movie = allData.Concert[:0]
		allData.Train = allData.Train[:0]

		newJson, err := json.MarshalIndent(allData, " ", " ")

		if err != nil {
			return
		}

		err = os.WriteFile(filename, newJson, 0644)

		return
	}

}

func CleanDataFile(ctx context.Context) {
	timer := time.NewTicker(5 * time.Minute)

	go func() {
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				deleteData()
			case <-ctx.Done():
				return
			}
		}
	}()
}
