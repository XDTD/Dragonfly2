package dfstore

import (
	"context"
	"fmt"
	"testing"

	"github.com/XDTD/Dragonfly2/client/config"
)

func GetDfstoreClient() (Dfstore, error) {

	cfg := config.NewDfstore()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	client := New(cfg.Endpoint)
	return client, nil
}

func TestCreateBucketWithContext(t *testing.T) {
	client, err := GetDfstoreClient()
	if err != nil {
		t.Error(err)
	}

	input := &CreateBucketInput{
		BucketName: "newbucketdragonflytest4",
	}

	err = client.CreateBucketWithContext(context.Background(), input)
	if err != nil {
		t.Error(err)
	}
}

func TestGetObjectListWithContext(t *testing.T) {
	client, err := GetDfstoreClient()
	if err != nil {
		t.Error(err)
	}

	input := &GetObjectMetadatasInput{
		BucketName: "test",
		Prefix:     "",
		Delimiter:  "/",
		Marker:     "",
		Limit:      20,
	}
	objMetas, err := client.GetObjectMetadatasWithContext(context.Background(), input)
	if err != nil {
		t.Error(err)
	}
	for _, objMeta := range objMetas {
		fmt.Printf("%v\n", objMeta)
	}

}

func TestGetObjectMetadataWithContext(t *testing.T) {
	client, err := GetDfstoreClient()
	if err != nil {
		t.Error(err)
	}

	input := &GetObjectMetadataInput{
		BucketName: "newbucketdragonflytest",
		ObjectKey:  "JuiceFS.png",
	}

	objMeta, err := client.GetObjectMetadataWithContext(context.Background(), input)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v", objMeta)

}

func TestCopyObjectWithContext(t *testing.T) {
	client, err := GetDfstoreClient()
	if err != nil {
		t.Error(err)
	}

	input := &CopyObjectInput{
		BucketName:           "newbucketdragonflytest",
		SourceObjectKey:      "JuiceFS.png",
		DestinationObjectKey: "dst6.png",
	}

	err = client.CopyObjectWithContext(context.Background(), input)
	if err != nil {
		t.Error(err)
	}
}
