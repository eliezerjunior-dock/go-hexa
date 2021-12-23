package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"context"
	"math/rand"
	"time"
	"strings"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	proto "github.com/go-hexa/proto-shared/generated/go/rate"

)

type server struct{}

var host = "0.0.0.0:60051" 
//var host = "0.0.0.0:60052" 

func initSetup(){
	addrs, err := net.InterfaceAddrs()

	var pid = ""
	var ip_adress = ""

	if err != nil {
		log.Println("Erro Fatal")
		panic(err)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip_adress = ipnet.IP.String()
			}
		}
	}

	pid = strconv.Itoa(os.Getpid())

	log.Printf("Setup : PID(%v) IP (%v) host (%v) \n", pid, ip_adress, host)
}

func main(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Rate gRPC server")

	initSetup()

	lis, err := net.Listen("tcp",host)
	if err != nil{
		log.Fatalf("Failed to listen %v", err)
	}

	var opts []grpc.ServerOption
 	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: time.Minute * 1}))

	s := grpc.NewServer(opts...)
	proto.RegisterRateServiceServer(s , &server{})

	go func(){
		log.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to server %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM )
	<-ch
	log.Println("Stopping server")
	s.Stop()
	log.Println("Stopping listener")
	lis.Close()
	log.Println("Stopped Done !!!")
}

func(*server) GetRateAccount(ctx context.Context, req *proto.RateRequest) (*proto.RateResponse, error) {
	log.Println("GetRateAccount")
	log.Printf("Incoming request data : %v", req)

	rand.Seed(time.Now().UnixNano())
	randon_i := rand.Intn(30)
	var stringBuilder strings.Builder

	rate := &proto.Rate{}
	rate.Rate = int32(randon_i)

	stringBuilder.WriteString("Taxa da Cliente -- ")
	stringBuilder.WriteString(req.GetAccount())

	rate.Description = stringBuilder.String()

	res := &proto.RateResponse {
		Rate: rate,
	}

	return res, nil
}