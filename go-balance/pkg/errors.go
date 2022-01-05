package pkg

import (
	"errors"
	"net/http"
)

var (
	ErrGetRate = errors.New("Erro no serviço RATE")
	ErrNotFound = errors.New("Item não encontrado")
	ErrInsert = errors.New("Erro na inserção do dado")
	ErrUnmarshal = errors.New("Erro na conversão do JSON")
	ErrUnauthorized = errors.New("Erro de autorização")
)

func HandlerHttpError(w http.ResponseWriter, err error) { 
	switch err {
		case ErrUnauthorized:
			w.WriteHeader(http.StatusUnauthorized)	
		default:
			w.WriteHeader(http.StatusInternalServerError)
	}
}