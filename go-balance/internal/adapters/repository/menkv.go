package repository

import (
	"log"
	"encoding/json"
	"sync"

	"github.com/go-hexa/go-balance/pkg"
	"github.com/go-hexa/go-balance/internal/core"
)

var mutex sync.Mutex

type memkv struct {
	kv map[string][]byte
}

func NewMemKV() *memkv {
	return &memkv{kv: map[string][]byte{}}
}

func (repo *memkv) ListBalance() ([]core.Balance, error) {
	log.Printf("ListBalance")

	log.Printf("####################################")
	log.Printf("- DataBase - ListBalance DataBase !!!!")
	log.Printf("####################################")

	var res []core.Balance
	for _, value := range repo.kv {
		balance := core.Balance{}
		err := json.Unmarshal(value, &balance)
		if err != nil {
			log.Printf("Erro : %v \n ", err)
			return nil, pkg.ErrUnmarshal
		}
		res = append(res, balance)
	}
	return res ,nil
}

func (repo *memkv) GetBalance(account string) (core.Balance, error) {
	log.Printf("GetBalance")

	log.Printf("####################################")
	log.Printf("- DataBase - GetBalance DataBase !!!!")
	log.Printf("####################################")

	if value, ok := repo.kv[account]; ok {
		balance := core.Balance{}
		err := json.Unmarshal(value, &balance)
		if err != nil {
			log.Printf("Erro : %v \n ", err)
			return core.Balance{}, pkg.ErrUnmarshal
		}
		return balance, nil
	}
	return core.Balance{}, pkg.ErrNotFound
}

func (repo *memkv) AddBalance(balance core.Balance) error {
	log.Printf("repo-AddBalance")

	log.Printf("####################################")
	log.Printf("- DataBase - AddBalance Data !!!!")
	log.Printf("####################################")
	
	mutex.Lock()
	defer mutex.Unlock()

	bytes, err := json.Marshal(balance)
	if err != nil {
		log.Printf("Erro : %v \n ", err)
		return pkg.ErrInsert
	}
	repo.kv[balance.Id] = bytes
	return nil
}