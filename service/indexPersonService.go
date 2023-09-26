package service

import (
	"go_practice/repository"
	"log"
)

type IIndexSavingService interface {
	IndexSaving()([]repository.Person, error)
}

type IndexSavingService struct {
	repo repository.IIndexPersonRepository
}

// サービスの作成
func NewIndexSavingService(repo repository.IIndexPersonRepository) *IndexSavingService {
	return &IndexSavingService{
		repo,
	}
}

func (iss IndexSavingService) IndexSaving() ([]repository.Person, error){
	repo, err := iss.repo.GetPersons()

	if err != nil {
		log.Fatal()
	}

	return repo, err
}
