package dummy_data

import (
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-hexa/go-balance/internal/core"
	proto "github.com/go-hexa/proto-shared/generated/go/balance"

 )

 func NewBalance(i int) core.Balance{
	acc := "acc-" + strconv.Itoa(i)
	description := "description-"+ strconv.Itoa(i) + " - UPDATED"
	
	balance := core.Balance{
		Id:    strconv.Itoa(i),
		Account: acc,
		Amount: 1,
		DateBalance: time.Now(),
		Description: description,
	}
	return balance
 }

 
 func NewBalancePB(i int) proto.Balance{
	id :=  strconv.Itoa(i)
	acc := "acc-" + strconv.Itoa(i)
	description := "description-"+ strconv.Itoa(i) + " - UPDATED"
	ts := timestamppb.Now()
	balance := proto.Balance{Id: id, 
		Account: acc, 
		Amount: 1, 
		DateBalance: ts, 
		Description: description,
	}
	return balance
 }