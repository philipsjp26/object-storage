package managers

import (
	"context"
	"object/storage/strategies"
)

type StrategyManager struct {
	storageStrategy strategies.StorageStrategy
}

func (m *StrategyManager) SetStorageStrategy(strategy strategies.StorageStrategy) {
	m.storageStrategy = strategy
}

func (m *StrategyManager) Store(objectName string, file []byte) (string, error) {
	result, err := m.storageStrategy.StoreObject(context.Background(), objectName, file)
	if err != nil {

		return "", err
	}
	return result, nil
}

func (m *StrategyManager) Get(objectName string) (string, error) {
	result, err := m.storageStrategy.GetObject(context.Background(), objectName)
	if err != nil {
		return "", err
	}
	return result, nil
}
