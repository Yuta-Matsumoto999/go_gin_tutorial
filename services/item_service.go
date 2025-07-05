package services

import (
	"github.com/Yuta-Matsumoto999/go_gin_tutorial/dto"
	"github.com/Yuta-Matsumoto999/go_gin_tutorial/models"
	"github.com/Yuta-Matsumoto999/go_gin_tutorial/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint) error
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{
		repository: repository,
	}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	existingItem, err := s.repository.FindById(itemId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != "" {
		existingItem.Name = updateItemInput.Name
	}

	if updateItemInput.Price != 0 {
		existingItem.Price = updateItemInput.Price
	}

	if updateItemInput.Description != "" {
		existingItem.Description = updateItemInput.Description
	}

	if updateItemInput.SoldOut {
		existingItem.SoldOut = updateItemInput.SoldOut
	}

	return s.repository.Update(itemId, *existingItem)
}

func (s *ItemService) Delete(itemId uint) error {
	return s.repository.Delete(itemId)
}
