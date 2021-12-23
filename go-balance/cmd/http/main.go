package main

import (
	"log"
	"time"
	"context"
	"syscall"
	"os"
	"os/signal"
	"net"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/go-hexa/go-balance/dummy_data"
	"github.com/go-hexa/go-balance/internal/adapters/repository"
	"github.com/go-hexa/go-balance/internal/core"
	"github.com/go-hexa/go-balance/internal/handlers/hdl_http"
	"github.com/go-hexa/go-balance/internal/adapters/client"
)

var my_pod core.Pod
var my_setup core.Setup 

func envVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("*** WARNING .env file NOT FOUND, using the os.env")
	}
	return os.Getenv(key)
}

func initSetup(){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Erro Fatal")
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				my_pod.Ip = ipnet.IP.String()
			}
		}
	}

	my_pod.Port = envVariable("PORT")
	if my_pod.Port == ""{
		my_pod.Port = "8900"
	}
	my_pod.Name = envVariable("NAME_POD")
	if my_pod.Name == ""{
		my_pod.Name = "no-name"
	}

	my_pod.PID = strconv.Itoa(os.Getpid())

	my_setup.ResponseTime = 0
	my_setup.ResponseStatusCode = 200
	my_setup.IsRandomTime = false
	my_setup.Count = 100
	
	my_pod.Setup = my_setup
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Starting Http Server 1.0")

	initSetup()

	handler_cliente	:= client.NewGrpcAdapterClient()
	repo 			:= repository.NewMemKV()
	service 		:= core.NewService(repo, handler_cliente)
	handler			:= hdl_http.NewHttpAdapter(service)
	
	//------------------------------------
	//Load dummy data
	for i:=0 ; i < 50; i++ {
		b := dummy_data.NewBalance(i)
		service.AddBalance(b)
	}
	//------------------------------------
	
	handleRequests(handler)
}

func handleRequests(handler *hdl_http.HttpAdapter) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(my_pod)
	})

	add_balance := myRouter.Methods(http.MethodPost).Subrouter()
    add_balance.HandleFunc("/add_balance", handler.AddBalance)

	list_balance := myRouter.Methods(http.MethodGet).Subrouter()
    list_balance.HandleFunc("/list_balance", handler.ListBalance)

	get_balance := myRouter.Methods(http.MethodGet).Subrouter()
    get_balance.HandleFunc("/get_balance", handler.GetBalance).Queries("id", "{id}")

	//list_balance.Use(MiddleWareHandler)

	s := http.Server{
		Addr:         ":" + my_pod.Port,      	
		Handler:      myRouter,                	          
		ReadTimeout:  time.Duration(60) * time.Second,   
		WriteTimeout: time.Duration(60) * time.Second,  
		IdleTimeout:  time.Duration(60) * time.Second, 
	}

	go func() {
		log.Printf("Server Running -> name: %v | pid: %v | ip: %v | port: %v) \n", my_pod.Name , my_pod.PID ,my_pod.Ip, my_pod.Port)
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Failed to connect server : %s\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	log.Println("-> Stopping initialized")

	ctx , cancel := context.WithTimeout(context.Background(), time.Duration(60) * time.Second)
	defer cancel()

	log.Println("--> Stopping gRPC server")
	if err := s.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Println("WARNING Dirty Shutdown : %s \n", err)
		return
	}
	log.Println("---> Stop DONE !!!")
}