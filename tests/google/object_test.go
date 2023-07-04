package google

import (
	"fmt"
	"io/ioutil"
	"object/storage/managers"
	"object/storage/strategies"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var file_dir = "../../documents/sample.pdf"
var creds = "../../gcs_accel.json"

func TestGetObject(t *testing.T) {
	var provider strategies.StorageStrategy

	gcsStrategy, err := strategies.NewGCSStorageStrategy(creds)
	if err != nil {
		assert.Error(t, err, "Error initiate")
	}
	provider = gcsStrategy

	storageManeger := &managers.StrategyManager{}
	storageManeger.SetStorageStrategy(provider)

	result, err := storageManeger.Get("sample.pdf")
	assert.NoError(t, err)
	t.Logf("result : %v", result)

}

func TestPutObject(t *testing.T) {
	var provider strategies.StorageStrategy

	gcsStrategy, err := strategies.NewGCSStorageStrategy(creds)
	if err != nil {
		assert.Error(t, err, "Error initiate")
	}
	provider = gcsStrategy

	storageManeger := &managers.StrategyManager{}
	storageManeger.SetStorageStrategy(provider)

	file, err := os.Open(file_dir)
	if err != nil {
		assert.Error(t, err, "Failed opened file")
	}
	defer file.Close()
	fileBytes, _ := ioutil.ReadAll(file)

	result, err := storageManeger.Store("sample.pdf", fileBytes)
	assert.NoError(t, err)
	t.Logf("result : %v", result)
}

func TestFileRoot(t *testing.T) {
	path := "../../documents/sample.pdf"
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("File exits")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	}
	dir, _ := os.Getwd()
	t.Logf("result : %v", dir)
	assert.NoError(t, err)
}
