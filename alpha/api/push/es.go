package push

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

type Context struct {
	ES          *elastic.Client
	BulkService *elastic.BulkService
	BulkOpCount int
}

// connect to ES instance and pass to context
func New(URL, indexName, ESType string, bulkOpCount int) (*Context, error) {

	client, err := elastic.NewClient(elastic.SetURL(URL), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	return &Context{
		ES:          client,
		BulkService: client.Bulk().Index(indexName).Type(ESType),
		BulkOpCount: bulkOpCount,
	}, nil
}

func (c *Context) PushToES(index, docId, ESType string, payload interface{}) error {
	ctx := context.Background()
	_, err := c.ES.Index().
		Index(index).
		Type(ESType).
		Id(docId).
		BodyJson(payload).
		Do(ctx)
	if err != nil {
		fmt.Printf("Failed to index, ID: %v", docId)
		return err
	}
	return nil
}
