package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	pb "github.com/MauricioCortesRo/Lab2_Proto"
	"google.golang.org/grpc"
)

var lock sync.Mutex
var port string = ":9005"
var path string = "DATA.txt"

type server struct {
	pb.UnimplementedInformation_TradesServer
}

func writeFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	lock.Lock()

	_, err = fmt.Fprintln(file, data)
	if err != nil {
		return err
	}
	lock.Unlock()
	return nil
}

func (s *server) Saves_Name(_ context.Context, req *pb.NameNodeRequest) (*pb.DataNodeResponse, error) {
	fmt.Println("Saving name: " + req.GetName() + " " + req.GetLastName() + " in DataNode")
	go writeFile(path, req.GetName()+" "+req.GetLastName()+" "+req.ID)
	fmt.Println("Name saved in DataNode")
	fmt.Println("")
	return &pb.DataNodeResponse{}, nil
}

func fileReader(filename string, data []string) ([]string, []string) {
	lock.Lock()
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		panic("cannot open file" + err.Error())
	}
	defer file.Close()

	var names []string
	var lastnames []string
	var id string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Split(line, " ")
		id = strings.Replace(parts[2], "\n", "", -1)
		for _, id_array := range data {
			if id == id_array {
				names = append(names, parts[0])
				lastnames = append(lastnames, parts[1])
			}
		}
	}

	lock.Unlock()
	return names, lastnames
}

func (s *server) Get_Name(_ context.Context, req *pb.NameNodeIDRequest) (*pb.DataNodeNamesResponse, error) {
	fmt.Println("Request received from NameNode, getting names from DataNode")
	fmt.Println("IDs: ", req.GetID())
	fmt.Println("")
	names, lastnames := fileReader(path, req.GetID())
	fmt.Println("Names and lastnames sent to NameNode")
	fmt.Println("Names: ", names)
	fmt.Println("Lastnames: ", lastnames)
	fmt.Println("")
	return &pb.DataNodeNamesResponse{
		Name:     names,
		LastName: lastnames,
	}, nil
}
func main() {
	_, err := os.Create(path)
	if err != nil {
		panic("cannot create file" + err.Error())
	}

	listener, err := net.Listen("tcp", port)

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterInformation_TradesServer(serv, &server{})
	if err := serv.Serve(listener); err != nil {
		panic("cannot initialize the DataNode server" + err.Error())
	}
	fmt.Println("DataNode server initialized")
}
