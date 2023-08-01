package dfstore

import (
	"context"
	"fmt"
	"testing"

	"d7y.io/dragonfly/v2/client/config"
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

	input := &ListObjectMetadatasInput{
		BucketName: "test",
		Prefix:     "",
		Delimiter:  "",
		Marker:     "JuiceFS.png",
		Limit:      2,
	}
	objMetas, err := client.ListObjectMetadatasWithContext(context.Background(), input)
	if err != nil {
		t.Error(err)
	}
	for _, objMeta := range objMetas {
		fmt.Printf("%v", objMeta)
	}

}

func TestGetObjectMetadataWithContext(t *testing.T) {
	client, err := GetDfstoreClient()
	if err != nil {
		t.Error(err)
	}

	input := &GetObjectMetadataInput{
		BucketName: "test",
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
		BucketName:    "test",
		SrcObjectKey:  "JuiceFS.png",
		DestObjectKey: "dst7.png",
	}

	err = client.CopyObjectWithContext(context.Background(), input)
	if err != nil {
		t.Error(err)
	}
}
