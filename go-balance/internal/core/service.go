package core

import (
	"log"
	"context"

	"github.com/go-hexa/go-balance/pkg"
)

type service struct {
	repo 	BalanceRepositoryPort
	cliente BalanceClientPort
	ctx 	context.Context
}

func NewService(repo BalanceRepositoryPort, clie BalanceClientPort) *service {
	return &service{
		cliente: clie,
		repo: repo,
		ctx: context.Background(),
	}
}

func (p *service) AddBalance(balance Balance) error {
	log.Printf("AddBalance")
	
	log.Printf("--------------------------------------")
	log.Printf("- AddBalance Doing Business Rules !!!!")
	log.Printf("--------------------------------------")

	err := p.repo.AddBalance(balance)
	if err != nil {
		return err
	}
	return nil
}

func (p *service) GetBalance(account string) (Balance, error) {
	log.Printf("GetBalance")
	
	log.Printf("--------------------------------------")
	log.Printf("- GetBalance Doing Business Rules !!!!")
	log.Printf("--------------------------------------")

	res, err := p.repo.GetBalance(account)
	if err != nil {
		return Balance{}, err
	}
	return res, nil
}

func (p *service) ListBalance() ([]Balance, error) {
	log.Printf("ListBalance")
	
	log.Printf("--------------------------------------")
	log.Printf("- ListBalance Doing Business Rules !!!!")
	log.Printf("--------------------------------------")

	res, err := p.repo.ListBalance()
	if err != nil {
		return nil, err
	}

	for idx, value := range res {
		rate, err := p.cliente.GetRate(value.Account)
		if err != nil {
			return nil, pkg.ErrGetRate
		}
		res[idx].Amount = value.Amount * rate
		log.Printf("--------------------------------------")
		log.Printf("- ListBalance Doing Business Rules !!!!")
		log.Println("")
		log.Printf("Account ( %v ) com Valor ( %v ) \n", value.Account , res[idx].Amount)
		log.Println("")
		log.Printf("--------------------------------------")
	}

	return res, nil
}
