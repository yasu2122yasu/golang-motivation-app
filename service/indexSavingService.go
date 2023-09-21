package service

import (
	"go_practice/repository"
	"log"
)


func IndexSavingService() ([]repository.Person, error){

	repo, err := repository.GetPersons()

	if err != nil {
		log.Fatal()
	}

	return repo, err
}
