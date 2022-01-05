package client

import(
	"log"
	"context"
	"time"

	"github.com/go-hexa/go-balance/internal/core"
	proto "github.com/go-hexa/proto-shared/generated/go/rate"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

const val_timeout = 3
const host =  "svc-go-rate-grpc:60051" 
//const host =  "go_rate_grpc:60051" //uso  docker
//const host = "0.0.0.0:60051" //uso local 

type GrpcAdapterClient struct {
	cliente core.BalanceClientPort
}

func NewGrpcAdapterClient() *GrpcAdapterClient {
	return &GrpcAdapterClient{	}
}

func (g *GrpcAdapterClient) GetRate(account string) (int32, error) {
	log.Printf("GetRate")
	
	log.Printf("--------------------------------------")
	log.Printf("- GetRate Calling another Service !!!!")
	log.Printf("--------------------------------------")

	var opts []grpc.DialOption
	opts = append(opts,grpc.FailOnNonTempDialError(true))
	opts = append(opts, grpc.WithInsecure())
    opts = append(opts, grpc.WithBlock())

	cc, err := grpc.Dial(host, opts...)
	if err != nil{
		log.Printf("Failed to connect : %v", err)
		return 1, err
	}
	defer func() {
		if err := cc.Close(); err != nil {
			log.Printf("Failed to close gPRC connection: %s", err)
		}
	}()

	c := proto.NewRateServiceClient(cc)

	req := &proto.RateRequest {
		Account: account,
	}

	timeout := val_timeout * time.Second
	header := metadata.New(map[string]string{"accept_language": "pt-BR", "jwt":"cookie"})
	ctx, cancel := context.WithTimeout(context.Background(), timeout) 
	ctx = metadata.NewOutgoingContext(ctx, header)
	defer cancel()

	res, err := c.GetRateAccount(ctx, req) 
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok{
			if statusErr.Code() == codes.DeadlineExceeded  {
				log.Printf("---> Timeout, deadline exceeded")
			} else {
				log.Printf("Error unexpected : %v \n", statusErr)
			}
		} else {
			log.Printf("Error to call PRC : %v", err)
		}
	} else {
		log.Println("Taxa do cliente : ", res.Rate.Rate)
	}

	return res.Rate.Rate, nil
}