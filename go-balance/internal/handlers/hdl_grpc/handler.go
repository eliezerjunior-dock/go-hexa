package hdl_grpc

import(
	"log"
	"context"

	"github.com/go-hexa/go-balance/internal/core"
	
	proto "github.com/go-hexa/proto-shared/generated/go/balance"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type GrpcAdapter struct {
	service core.BalanceServicePort
}

func NewGrpcAdapter(serv core.BalanceServicePort) *GrpcAdapter {
	return &GrpcAdapter{
		service: serv,
	}
}

func (g *GrpcAdapter) AddBalance(ctx context.Context, req *proto.AddBalanceRequest) (*proto.AddBalanceResponse, error) {
	log.Printf("/AddBalance")
	log.Printf("Incoming request data : %v", req)

	_balance := core.Balance{}
	_balance.Id 		= req.GetBalance().GetId()
	_balance.Account 	=  req.GetBalance().GetAccount()
	_balance.Amount 	=  req.GetBalance().GetAmount()
	_balance.DateBalance = req.GetBalance().GetDateBalance().AsTime()
	_balance.Description =  req.GetBalance().GetDescription()

	err := g.service.AddBalance(_balance)
	if err != nil{
		log.Print(err)
		return nil, status.Errorf(codes.InvalidArgument, "Erro na inclusão do item")
	}

	res := &proto.AddBalanceResponse {
		Result: true,
	}
	return res, nil
}

func (g *GrpcAdapter) GetBalance(ctx context.Context, req *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	log.Printf("/GetBalance")
	log.Printf("Incoming request data : %v", req)

	result, err := g.service.GetBalance(req.GetId())
	if err != nil{
		log.Print(err)
		return nil, status.Errorf(codes.InvalidArgument, "Item não encontrado")
	}
	
	_balance :=  &proto.Balance{}
	_balance.Id 		= result.Id
	_balance.Account 	= result.Account
	_balance.Amount 	= result.Amount
	_balance.DateBalance = timestamppb.New(result.DateBalance)
	_balance.Description = result.Description

	res := &proto.GetBalanceResponse {
		Balance: _balance,
	}
	return res, nil
}

func (g *GrpcAdapter) ListBalance(ctx context.Context, req *proto.ListBalanceRequest) (*proto.ListBalanceResponse, error) {
	log.Printf("/ListBalance")
	log.Printf("Incoming request data : %v", req)

	result, err := g.service.ListBalance()
	if err != nil{
		log.Print(err)
		return nil, status.Errorf(codes.InvalidArgument, "Erro na listagem dos itens")
	}

	var array_balance []*proto.Balance
	for _, value := range result {
		_balance :=  &proto.Balance{}
		_balance.Id 		= value.Id
		_balance.Account 	= value.Account
		_balance.Amount 	= value.Amount
		_balance.DateBalance =  timestamppb.New(value.DateBalance)
		_balance.Description = value.Description

		array_balance = append(array_balance, _balance)
	}

	res := &proto.ListBalanceResponse {
	 	Balance: array_balance,
	}

	return res, nil
}
