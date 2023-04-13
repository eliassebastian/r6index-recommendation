package weaviate

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

func createSimpleTestClient() *weaviate.Client {
	cfg := weaviate.Config{
		Host:   "localhost:6464",
		Scheme: "http",
	}

	return weaviate.New(cfg)
}

func cleanupSimpleTestClient(t *testing.T, client *weaviate.Client) {
	// Clean up test class and by that also all data
	err := client.Schema().ClassDeleter().WithClassName("TestR6Index").Do(context.Background())
	if err != nil {
		t.Errorf("weaviate class delete error: got %v want nil", err)
	}
}

func createWeaviateTestSchemaWithVectorizorlessClass(t *testing.T, client *weaviate.Client) {
	vectorizorlessClass := &models.Class{
		Class:       "TestR6Index",
		Description: "Test R6Index Class",
		Vectorizer:  "none",
	}

	err := client.Schema().ClassCreator().WithClass(vectorizorlessClass).Do(context.Background())
	if err != nil {
		t.Errorf("weaviate class creator error: got %v want nil", err)
	}
}

func TestWeaviateData(t *testing.T) {

	t.Run("Test Single Vector Object", func(t *testing.T) {
		client := createSimpleTestClient()
		createWeaviateTestSchemaWithVectorizorlessClass(t, client)

		vec := []float32{211.0, 0.76, 35.0, 3424.0}

		wrapper, errCreate := client.Data().Creator().
			WithClassName("TestR6Index").
			WithID("6844b415-aa94-43c9-8823-9389e4816902").
			WithVector(vec).
			Do(context.Background())

		if errCreate != nil {
			t.Errorf("weaviate data creator error: got %v want nil", errCreate)
		}

		if wrapper == nil {
			t.Errorf("weaviate data creator error: got nil want not nil")
		}

		object, objErr := client.Data().ObjectsGetter().
			WithClassName("TestR6Index").
			WithID("6844b415-aa94-43c9-8823-9389e4816902").
			WithAdditional("vector").
			Do(context.Background())

		if objErr != nil {
			t.Errorf("weaviate data getter error: got %v want nil", objErr)
		}

		if len(object) == 0 {
			t.Errorf("weaviate data getter error: got empty object want not empty")
		}

		arr := []float32(object[0].Vector)
		if !reflect.DeepEqual(arr, vec) {
			t.Errorf("weaviate data getter error: got %v want %v", arr, vec)
		}

		cleanupSimpleTestClient(t, client)
	})

	t.Run("Test Batch Import Vector Object", func(t *testing.T) {
		client := createSimpleTestClient()
		createWeaviateTestSchemaWithVectorizorlessClass(t, client)

		// sample batch data
		data := []*models.Object{
			{
				ID:     "6844b415-aa94-43c9-8823-9389e4816910",
				Vector: []float32{211.0, 0.76, 35.0, 3424.0},
				Class:  "TestR6Index",
			},
			{
				ID:     "6844b415-aa94-43c9-8823-9389e4816914",
				Vector: []float32{250.0, 0.80, 35.0, 5000.0},
				Class:  "TestR6Index",
			},
			{
				ID:     "6844b415-aa94-43c9-8823-9389e4816923",
				Vector: []float32{110.0, 0.43, 15.0, 1000.0},
				Class:  "TestR6Index",
			},
			{
				ID:     "6844b415-aa94-43c9-8823-9389e4816905",
				Vector: []float32{300.0, 0.54, 17.0, 1233.0},
				Class:  "TestR6Index",
			},
		}

		batchR, err := client.Batch().ObjectsBatcher().WithObjects(data...).Do(context.Background())

		if err != nil {
			t.Errorf("weaviate batch creator error: got %v want nil", err)
		}

		if batchR == nil {
			t.Errorf("weaviate batch creator error: got nil want not nil")
		}

		if len(batchR) != 4 {
			t.Errorf("weaviate batch creator error: got %d want %d", len(batchR), 4)
		}

		log.Println(batchR)

		cleanupSimpleTestClient(t, client)
	})

}
