package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	pb "github.com/MauricioCortesRo/Lab2_Proto"
	"google.golang.org/grpc"
)

var ID int = 1
var lock sync.Mutex
var DataNode string
var path string = "DATA.txt"
var ServerDataNode1 string = "dist007:9005"
var ServerDataNode2 string = "dist008:9006"

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

func send_names(ID int, name string, lastname string, ip_puerto string) {
	conn, err := grpc.Dial(ip_puerto, grpc.WithInsecure()) //
	if err != nil {
		panic("Cannot connect to DataNode server" + err.Error())
	}

	defer conn.Close()
	fmt.Println("Sending person to DataNode: " + strconv.Itoa(ID) + " " + name + " " + lastname + "\n")
	serviceClient := pb.NewInformation_TradesClient(conn)
	res, errS := serviceClient.Saves_Name(context.Background(), &pb.NameNodeRequest{
		ID:       strconv.Itoa(ID),
		Name:     name,
		LastName: lastname,
	})

	if errS != nil {
		panic("Person not sended" + err.Error())
	}
	fmt.Printf(res.Response)
}

func (s *server) Notificate(_ context.Context, req *pb.Continent) (*pb.NameNodeResponse, error) {
	fmt.Println("Request received from Continent: " + req.GetName() + " " + req.GetLastName() + " " + req.GetStatus())
	firstChar := req.GetLastName()[0]
	if firstChar >= 'A' && firstChar <= 'M' {
		DataNode = "1"
		data := strconv.Itoa(ID) + ";" + DataNode + ";" + req.GetStatus()
		go writeFile(path, data)
		go send_names(ID, req.GetName(), req.GetLastName(), ServerDataNode1)
	} else {
		DataNode = "2"
		data := strconv.Itoa(ID) + ";" + DataNode + ";" + req.GetStatus()
		go writeFile(path, data)
		go send_names(ID, req.GetName(), req.GetLastName(), ServerDataNode2)
	}
	ID = ID + 1

	return &pb.NameNodeResponse{
		Response: "Persona Recibida",
	}, nil
}

func get_names(IDs []string, DataNode string) ([]string, []string) {
	connn, err := grpc.Dial(DataNode, grpc.WithInsecure())
	if err != nil {
		panic("Cannot connect to DataNode server" + err.Error())
	}

	defer connn.Close()

	fmt.Println("Sending IPs to DataNode")

	serviceClientt := pb.NewInformation_TradesClient(connn)
	res, errS := serviceClientt.Get_Name(context.Background(), &pb.NameNodeIDRequest{
		ID: IDs,
	})
	if errS != nil {
		panic("Cannot send IPs" + err.Error())
	}

	fmt.Println("Names received from DataNode")
	return res.GetName(), res.GetLastName()
}

func (s *server) ONU(_ context.Context, req *pb.ONURequest) (*pb.StatusResponse, error) {
	fmt.Print("Request received from ONU\n")
	names := make([]string, 0, 5)
	lastnames := make([]string, 0, 5)

	ID_datanode1 := make([]string, 0, 5)
	ID_datanode2 := make([]string, 0, 5)

	lock.Lock()

	file, err := os.Open(path)
	if err != nil {
		panic("cannot open file" + err.Error())
	}

	scanner := bufio.NewReader(file)
	fmt.Println("Reading file to get Ids")
	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Split(line, ";")
		if err != nil {
			panic("Error converting string to int")
		}

		parts[2] = strings.Replace(parts[2], "\n", "", -1)

		if req.GetStatus() == "Infected" {
			if parts[2] == "Infected" {
				if parts[1] == "1" {
					ID_datanode1 = append(ID_datanode1, parts[0])
				} else {
					ID_datanode2 = append(ID_datanode2, parts[0])
				}
			}
		} else if req.GetStatus() == "Dead" {
			if parts[2] == "Dead" {
				if parts[1] == "1" {
					ID_datanode1 = append(ID_datanode1, parts[0])
				} else {
					ID_datanode2 = append(ID_datanode2, parts[0])
				}
			}
		}
	}
	fmt.Println("Ids getted successfully, sending to DataNodes")
	names1, lastname1 := get_names(ID_datanode1, ServerDataNode1) //Cambiar por IP correspondiente a DataNode1
	names2, lastname2 := get_names(ID_datanode2, ServerDataNode2) //Cambiar por IP correspondiente a DataNode2
	names = append(names, names1...)
	names = append(names, names2...)
	lastnames = append(lastnames, lastname1...)
	lastnames = append(lastnames, lastname2...)

	fmt.Println("Names and lastnames received from DataNodes")
	fmt.Println("Names: ", names)
	fmt.Println("Lastnames: ", lastnames)

	lock.Unlock()
	fmt.Println("Sending names and lastnames to ONU")
	return &pb.StatusResponse{
		Name:     names,
		LastName: lastnames,
	}, nil
}

func handleConnection(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic("cannot create tcp connection on port " + port + ": " + err.Error())
	}
	serv := grpc.NewServer()
	pb.RegisterInformation_TradesServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}

func main() {

	_, err := os.Create(path)
	if err != nil {
		panic("cannot create file" + err.Error())
	}

	go handleConnection(":9000") //Asia
	go handleConnection(":9001") //Australia
	go handleConnection(":9002") //Europa
	go handleConnection(":9003") //Latinoamerica
	go handleConnection(":9004") //ONU

	select {}
}
