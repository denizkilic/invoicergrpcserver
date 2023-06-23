package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"invoicergrpcserver/invoicer"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal("cannot create listener: %s", err)
	}

	serverRegisterer := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegisterer, service)
	err = serverRegisterer.Serve(listen)
	if err != nil {
		log.Fatal("cannot served: %s", err)
	}

}
