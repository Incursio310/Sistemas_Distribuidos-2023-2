package main

import (
	"context"
	"fmt"
	"os"

	pb "github.com/MauricioCortesRo/Lab2_Proto"
	"google.golang.org/grpc"
)

var ServerNameNode string = "dist006:9004"

func menu(serviceClient pb.Information_TradesClient) {

	var request string = ""
	for request != "exit" {
		fmt.Print("Enter a request (Infected, Dead or exit to finish): ")
		fmt.Scanln(&request)

		var Status string
		switch request {
		case "Infected":
			Status = "Infected"
		case "Dead":
			Status = "Dead"
		default:
			fmt.Println("End of ONU requests")
			os.Exit(0)
		}

		res, err := serviceClient.ONU(context.Background(), &pb.ONURequest{
			Status: Status,
		})

		fmt.Print("Requesting the " + request + " to the NameNode\n")

		if err != nil {
			panic("Persons not received" + err.Error())
		}
		for i := 0; i < len(res.GetLastName()); i++ {
			fmt.Println(res.GetName()[i], res.GetLastName()[i])
		}
	}
}

func main() {
	conn, err := grpc.Dial(ServerNameNode, grpc.WithInsecure())
	if err != nil {
		panic("Cannot connect to DataNode server" + err.Error())
	}
	serviceClient := pb.NewInformation_TradesClient(conn)
	go menu(serviceClient)

	select {}
}
