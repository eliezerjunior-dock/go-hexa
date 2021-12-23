package hdl_http

import(
	"log"
	"net/http"
	"encoding/json"

	"github.com/go-hexa/go-balance/internal/core"
)

type HttpAdapter struct {
	service core.BalanceServicePort
}

func NewHttpAdapter(serv core.BalanceServicePort) *HttpAdapter {
	return &HttpAdapter{
		service: serv,
	}
}

func (h *HttpAdapter) AddBalance(rw http.ResponseWriter, req *http.Request) {
	log.Printf("/AddBalance")
	rw.Header().Set("Content-Type", "application/json")

	balance := core.Balance{}
	err := json.NewDecoder(req.Body).Decode(&balance)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusBadRequest)
        return
    }

	err = h.service.AddBalance(balance)
	if err != nil{
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(err.Error())
		return
	}

	json.NewEncoder(rw).Encode("sucesso")
	return
}

func (h *HttpAdapter) GetBalance(rw http.ResponseWriter, req *http.Request) {
	log.Printf("/GetBalance")
	rw.Header().Set("Content-Type", "application/json")

	result, err := h.service.GetBalance(req.FormValue("id"))
	if err != nil{
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(err.Error())
		return
	}

	json.NewEncoder(rw).Encode(result)
	return
}

func (h *HttpAdapter) ListBalance(rw http.ResponseWriter, req *http.Request) {
	log.Printf("/ListBalance")
	rw.Header().Set("Content-Type", "application/json")

	result, err := h.service.ListBalance()
	if err != nil{
		json.NewEncoder(rw).Encode(err.Error())
		return
	}

	json.NewEncoder(rw).Encode(result)
	return
}
