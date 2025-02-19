package main

import (
	"context"
	"log"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// CustomDatasource struct
type CustomDatasource struct{}

// QueryData processes Grafana queries
func (d *CustomDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	response := backend.NewQueryDataResponse()

	for _, q := range req.Queries {
		res := d.handleQuery(q)
		response.Responses[q.RefID] = res
	}

	return response, nil
}

func (d *CustomDatasource) handleQuery(q backend.DataQuery) backend.DataResponse {
	// Create response structure
	response := backend.DataResponse{}

	// Example values for the query
	time := q.TimeRange.From.Unix() // Use the timestamp from the query's time range
	value := 42                     // Simulated value for the query
	status := "ok"                  // Simulated status for the query

	// Create fields with the actual data
	timeField := data.NewField("time", nil, []int64{time})
	valueField := data.NewField("value", nil, []float64{float64(value)})
	statusField := data.NewField("status", nil, []string{status})

	// Create a data frame to hold the fields
	frame := data.NewFrame("response", timeField, valueField, statusField)

	// Add the frame to the response
	response.Frames = append(response.Frames, frame)

	return response
}

func main() {
	// Serve the plugin
	err := datasource.Serve(datasource.ServeOpts{
		QueryDataHandler: &CustomDatasource{},
	})
	if err != nil {
		log.Fatalf("Failed to start plugin: %v", err)
	}
}
