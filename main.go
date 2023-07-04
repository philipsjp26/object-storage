package main

import (
	"object/storage/managers"
	"object/storage/strategies"
)

const STORAGE_PROVIDER = "GOOGLE"

func main() {
	var provider strategies.StorageStrategy
	switch STORAGE_PROVIDER {
	case "GOOGLE":
		gcsStrategy, err := strategies.NewGCSStorageStrategy("")
		if err != nil {
			panic(err)
		}
		provider = gcsStrategy
	case "MINIO":
		minioStrategy, err := strategies.NewMinioStorageStrategy()
		if err != nil {
			panic(err)
		}
		provider = minioStrategy
	default:
		minioStrategy, err := strategies.NewMinioStorageStrategy()
		if err != nil {
			panic(err)
		}
		provider = minioStrategy
	}

	storageManager := &managers.StrategyManager{}

	storageManager.SetStorageStrategy(provider)

}
