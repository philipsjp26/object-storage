package minio

import (
	"io/ioutil"
	"object/storage/managers"
	"object/storage/strategies"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var file_dir = "../../documents/sample.pdf"

func TestStoreObject(t *testing.T) {
	var provider strategies.StorageStrategy
	minioStrategy, err := strategies.NewMinioStorageStrategy()
	if err != nil {
		panic(err)
	}
	provider = minioStrategy

	storageManager := &managers.StrategyManager{}

	storageManager.SetStorageStrategy(provider)

	file, err := os.Open(file_dir)
	if err != nil {
		assert.Error(t, err, "Failed opened file")
	}
	defer file.Close()
	fileBytes, _ := ioutil.ReadAll(file)
	result, err := storageManager.Store("telor.pdf", fileBytes)
	assert.NoError(t, err)
	t.Logf("result : %v", result)
}
