package data

import (
	"fmt"
)

const MEMORY = "memory"

type ProviderData interface {
	GetUrlList() []string
}

func GetProviderData(source string) (ProviderData, error) {
	if source == "" {
		return nil, fmt.Errorf("источник данных не указан")
	}

	switch source {
	case MEMORY:
		return &MemoryData{}, nil
	default:
		return nil, fmt.Errorf("источник данных не найден %s", source)
	}
}
