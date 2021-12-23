package pkg

import (
	"errors"
)

var (
	ErrGetRate = errors.New("Erro no serviço RATE")
	ErrNotFound = errors.New("Item não encontrado")
	ErrInsert = errors.New("Erro na inserção do dado")
	ErrUnmarshal = errors.New("Erro na conversão do JSON")
)