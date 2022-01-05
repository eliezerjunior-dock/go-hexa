package main

import (
	"fmt"
	"log"
	"context"
	"time"
	"flag"
	"os"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	//"google.golang.org/grpc/backoff"

	"github.com/go-hexa/go-balance/pkg"
	"github.com/go-hexa/go-balance/dummy_data"
	proto "github.com/go-hexa/proto-shared/generated/go/balance"
)

func get(c proto.BalanceServiceClient, done chan string){
	for i:=0; i < 3600; i++ {
		ListBalance(c , 3 * time.Second)
		time.Sleep(time.Millisecond * time.Duration(1000))
	}
	done <- "END"
}	

func post(c proto.BalanceServiceClient, done chan string){
	for a:=0; a < 3600; a++ {
		for i:=0; i < 50; i++ {
			b := dummy_data.NewBalancePB(i)
			req := &proto.AddBalanceRequest {
				Balance: &b,
			}
			AddBalance(c , req , 5 * time.Second)
		}
		time.Sleep(time.Millisecond * time.Duration(1000))
	}
	done <- "END"
}

func GetBalance(c proto.BalanceServiceClient, timeout time.Duration){
	fmt.Println("GetBalance")

	req := &proto.GetBalanceRequest {
		Id: "3",
	}

	header := metadata.New(map[string]string{"accept_language": "pt-BR", "jwt":"cookie"})
	ctx, cancel := context.WithTimeout(context.Background(), timeout) 
	ctx = metadata.NewOutgoingContext(ctx, header)
	defer cancel()

	res, err := c.GetBalance(ctx, req) 
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
		result, err := pkg.ProtoToJSON(res)
		if err != nil {
			log.Printf(" **** Failed to connect %v", err)
			panic(err)
		}
		fmt.Println(result)
		//fmt.Printf("Timestamppb.AsTime() : %s\n", res.Balance.DateBalance.AsTime().String())
	}
}

func ListBalance(c proto.BalanceServiceClient, timeout time.Duration){
	fmt.Println("#### ListBalance")

	req := &proto.ListBalanceRequest {}

	header := metadata.New(map[string]string{"accept_language": "pt-BR", "jwt":"cookie"})

	ctx, cancel := context.WithTimeout(context.Background(), timeout) 
	ctx = metadata.NewOutgoingContext(ctx, header)
	defer cancel()

	res, err := c.ListBalance(ctx, req) 
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
		result, err := pkg.ProtoToJSON(res)
		if err != nil {
			log.Printf(" **** Failed to connect %v", err)
			panic(err)
		}
		fmt.Println(result)
		//fmt.Printf("Timestamppb.AsTime() : %s\n", res.Balance.DateBalance.AsTime().String())
	}
}

func AddBalance(c proto.BalanceServiceClient, req *proto.AddBalanceRequest, timeout time.Duration){
	fmt.Println("#### AddBalance")
	
	header := metadata.New(map[string]string{"accept_language": "pt-BR", "jwt":"cookie"})
	ctx, cancel := context.WithTimeout(context.Background(), timeout) 
	ctx = metadata.NewOutgoingContext(ctx, header)
	defer cancel()

	res, err := c.AddBalance(ctx, req) 
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
		fmt.Println(res)
		//fmt.Printf("Timestamppb.AsTime() : %s\n", res.Balance.DateBalance.AsTime().String())
	}
}

func main(){
	fmt.Println("Balance gRPC client")

	port := flag.String("port","","")
	flag.Parse()
	flag.VisitAll(func (f *flag.Flag) {
		if f.Value.String()=="" {
			fmt.Printf("A flag -%v não foi informado \n", f.Name )
			os.Exit(1)
		}
	})

	var host = "127.0.0.1:" + *port
	
	// var DefaultConfig = backoff.Config{
	// 	BaseDelay:  1.0 * time.Second,
	// 	Multiplier: 2.0,
	// 	Jitter:     0.2,
	// 	MaxDelay:   120 * time.Second,
	// }

	// var opts2 := []grpc_retry.CallOption{
	// 	grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100 * time.Millisecond)),
	// 	grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
	// }

	var opts []grpc.DialOption
	opts = append(opts, grpc.FailOnNonTempDialError(true)) // Wait for ready
	opts = append(opts, grpc.WithBlock()) // Wait for ready
	opts = append(opts, grpc.WithInsecure()) // no TLS

	// opts = append(opts, grpc.WithConnectParams(grpc.ConnectParams{
    //  											Backoff: DefaultConfig, })) 

	cc, err := grpc.Dial(host, opts...)
	if err != nil{
		log.Printf(" **** Failed to connect %v", err)
		panic(err)
	}
	defer func() {
		if err := cc.Close(); err != nil {
			log.Printf("Failed to close gPRC connection: %s", err)
		}
	}()

	c := proto.NewBalanceServiceClient(cc)

	done := make(chan string)
	
	log.Println("-----------------------------")
	log.Println("Goroutine - Post Data")
	go post(c, done)
	log.Println("End Post Data")
	log.Println("-----------------------------")

	log.Println("-----------------------------")
	log.Println("Goroutine - Get Data")
	//go get(c, done)
	log.Println("End Reading Data")
	log.Println("-----------------------------")

	log.Println(<-done)
}
