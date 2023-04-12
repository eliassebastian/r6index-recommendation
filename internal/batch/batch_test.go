package batch

import (
	"log"
	"testing"
	"time"
)

func TestBatchPipeline(t *testing.T) {
	// Create a BatchPipeline with a max size of 3 and a max wait time of 5 second
	pipeline := NewBatchPipeline(3, 5*time.Second, func(data []interface{}) error {
		// Print the data and return nil (no error)
		if len(data) > 3 {
			t.Errorf("BatchPipeline.data = %v, want %v or less", len(data), 3)
		}

		return nil
	})

	// Add 5 items to the pipeline
	pipeline.Add("item1")
	pipeline.Add("item2")
	pipeline.Add("item3")
	pipeline.Add("item4")
	pipeline.Add("item5")

	// Wait for a second to allow the pipeline to flush
	time.Sleep(11 * time.Second)

	if len(pipeline.data) != 0 {
		t.Errorf("BatchPipeline.data = %v, want %v", len(pipeline.data), 0)
	}
}

func TestBatchPipelineWithErrors(t *testing.T) {
	// Create a BatchPipeline with a max size of 2 and a max wait time of 1 second
	pipeline := NewBatchPipeline(2, 1*time.Second, func(data []interface{}) error {
		// Return an error if the length of data is greater than or equal to 2
		if len(data) > 2 {
			t.Errorf("BatchPipeline.data = %v, want %v or less", len(data), 2)
		}

		return nil
	})

	// Add 3 items to the pipeline
	pipeline.Add("item1")
	pipeline.Add("item2")
	pipeline.Add("item3")
	pipeline.Add("item4")
	pipeline.Add("item5")

	// Wait for a second to allow the pipeline to flush
	time.Sleep(10 * time.Second)
	pipeline.flushChan <- struct{}{}
	log.Println("exiting channel")
	<-pipeline.flushChan
}
