package repositories

import (
	"errors"

	"github.com/Yuta-Matsumoto999/go_gin_tutorial/models"
)

var ErrItemNotFound = errors.New("item not found")

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{
		items: items,
	}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, ErrItemNotFound
}
