package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/MauricioCortesRo/Proto"

	"github.com/streadway/amqp"
)

type server struct {
	pb.UnimplementedCentralServiceServer
}

func FileReader(filename string) int {

	var values int

	contenido, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return values
	}
	defer contenido.Close()
	lector := bufio.NewScanner(contenido)
	lineas := []string{}
	for lector.Scan() {
		lineas = append(lineas, lector.Text())
	}

	if err := lector.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return values
	}

	values, _ = strconv.Atoi(lineas[0])

	return values
}

func rabbitPublish(requiredKeys string) {

	var message string = ServerName + "-" + requiredKeys

	//channel creation

	err := ch.Publish(
		"",
		"KeysQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(ServerName + " successfully Published needed keys: " + requiredKeys)
}

func (s *server) KeyReleaseAnnouncement(_ context.Context, req *pb.CreateAnnouncementRequest) (*pb.CreateAnnouncementResponse, error) {
	go rabbitPublish(requiredKeys)
	fmt.Println("confirm avalible keys: ", req.GetAvalibleKeys())
	return &pb.CreateAnnouncementResponse{
		Status: true,
	}, nil
}

func (s *server) NonMatchedUsers(_ context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersResponse, error) {
	IntRequiredKeys, _ := strconv.Atoi(requiredKeys)
	fmt.Println("achieved keys: ", strconv.Itoa(IntRequiredKeys-int(req.GetNonMatched())))
	requiredKeys = strconv.Itoa(int(req.GetNonMatched()))
	fmt.Println("actually remaining users: " + requiredKeys)
	return &pb.CreateUsersResponse{
		Status: true,
	}, nil
}

// server name and port (the only code that changes between servers)
var ServerName string = "Asia"
var ServerPort string = "9001"
var requiredKeys string

// rabbitmq connection
var conn, connErr = amqp.Dial("amqp://guest:guest@dist008:5673/")
var ch, chErr = conn.Channel()

func main() {
	//check if the rabbitmq connection failed
	if connErr != nil {
		fmt.Println(connErr)
		panic(connErr)
	}
	defer conn.Close()
	if chErr != nil {
		fmt.Println(chErr)
		panic(chErr)
	}
	defer ch.Close()
	fmt.Println(ServerName + " successfully connected to our rabbitMQ instance!")

	//calculation of the required keys in the iteration
	var value int = FileReader("ServerAsia/parametros_de_inicio.txt")
	var variation int = int(math.Round(float64(value / 5)))
	requiredKeys = strconv.Itoa(int(math.Round(float64(value/2))) + rand.Intn(2*variation) - variation)

	//start up of the regional server: it waits for the global announcement of the keys release

	lis, err := net.Listen("tcp", ":"+ServerPort)
	if err != nil {
		log.Fatalf(ServerName+" failed to listen on port "+ServerPort+": %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCentralServiceServer(grpcServer, &server{})

	log.Println(ServerName + " on port:" + ServerPort + " is UP")
	log.Println("initial keys demand: " + requiredKeys)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf(ServerName+" failed to serve gRPC server over port "+ServerPort+": %v", err)
	}

	//end conection

}

/*


func generateID() int {
	return rand.Int()
}





	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewCentralServiceClient(conn)
	res1, err1 := serviceClient.Create(context.Background(), &pb.CreateRequest{
		Request: &pb.Request{
			A: 1,
			B: 2,
		},
		RequestId: int64(generateID()),
	})

	res2, err2 := serviceClient.Calculate(context.Background(), &pb.Request{
		A: 1,
		B: 2,
	})

	if err1 != nil {
		panic("Reques is not created " + err.Error())
	}
	if err2 != nil {
		panic("Calculation " + err.Error())
	}
	fmt.Println(res1.RequestId)
	fmt.Println(res2.Result)
*/
