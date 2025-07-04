package repositories

import (
	"errors"

	"github.com/Yuta-Matsumoto999/go_gin_tutorial/models"
)

var ErrItemNotFound = errors.New("item not found")

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(itemId uint, updateItem models.Item) (*models.Item, error)
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

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	lastId := uint(len(r.items))
	newItem.ID = lastId + 1
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(itemId uint, updateItem models.Item) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == itemId {
			item.Name = updateItem.Name
			item.Price = updateItem.Price
			item.Description = updateItem.Description
			item.SoldOut = updateItem.SoldOut
			return &item, nil
		}
	}
	return nil, ErrItemNotFound
}
