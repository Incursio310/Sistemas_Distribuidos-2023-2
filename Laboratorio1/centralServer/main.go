package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	pb "github.com/MauricioCortesRo/Proto"
	"github.com/streadway/amqp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FileReader(filename string) [3]int {

	var values [3]int

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

	for index, number := range strings.Split(lineas[0], "-") {
		values[index], _ = strconv.Atoi(number)
	}

	values[2], _ = strconv.Atoi(lineas[1])

	return values
}

func RegionToIndex(region string) int {
	if region == "America" {
		return 0
	} else if region == "Asia" {
		return 1
	} else if region == "Europa" {
		return 2
	} else if region == "Oceania" {
		return 3
	} else {
		return 4
	}
}

func main() {

	//Queue setting
	conn, err := amqp.Dial("amqp://guest:guest@dist008:5673/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	//queue creation
	q, err := ch.QueueDeclare(
		"KeysQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	msgs, _ := ch.Consume(
		"KeysQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	fmt.Println("Keys queue created. Initial status: ", q)
	//queue is seted

	//consume the initial parameters
	var values [3]int = FileReader("centralServer/parametros_de_inicio.txt")

	//servers connection

	ports := [4]string{"9000", "9001", "9002", "9003"}
	IPs := [4]string{"dist005", "dist006", "dist007", "dist008"}

	var serviceClient [4]pb.CentralServiceClient

	var ServersRemaining [4]int
	for index, port := range ports {
		conn, err := grpc.Dial(IPs[index]+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic("cannot connect with server on port: " + port + " with error: " + err.Error())
		}

		serviceClient[index] = pb.NewCentralServiceClient(conn)
		ServersRemaining[index] = 1
	}
	//end of the connections
	var flag bool = false
	if values[2] == -1 {
		flag = true
		values[2] = 10
	}

	for iteration := 0; iteration < values[2]; iteration++ {
		if flag {
			log.Println("iteration " + strconv.Itoa(iteration+1) + " of inf")
		} else {
			log.Println("iteration " + strconv.Itoa(iteration+1) + " of " + strconv.Itoa(values[2]))
		}
		//keys release announcement
		var SetedAvalibleKeys int = rand.Intn(values[1]+1-values[0]) + values[0]

		for index := range ports {
			_, KeysError := serviceClient[index].KeyReleaseAnnouncement(context.Background(), &pb.CreateAnnouncementRequest{
				AvalibleKeys: int64(SetedAvalibleKeys),
			})

			if KeysError != nil {
				panic("The keys release announcement had troubles " + KeysError.Error())
			}
		}
		log.Println("Announcement of " + strconv.Itoa(SetedAvalibleKeys) + " keys was successfully!")
		//end of the announcement

		//rabbit consume code

		fmt.Print("\n----- Receiving messages -----\n\n")
		var counter int = 1
		for d := range msgs {
			fmt.Printf("Recieved message from: %s\n keys", d.Body)
			result := strings.Split(string(d.Body), "-")
			region := result[0]
			keys, _ := strconv.Atoi(result[1])
			var nonMatched int
			if keys <= SetedAvalibleKeys {
				SetedAvalibleKeys = SetedAvalibleKeys - keys
				nonMatched = 0

			} else {
				nonMatched = keys - SetedAvalibleKeys
				keys = SetedAvalibleKeys
				SetedAvalibleKeys = 0

			}
			_, KeysError := serviceClient[RegionToIndex(region)].NonMatchedUsers(context.Background(), &pb.CreateUsersRequest{
				NonMatched: int64(nonMatched),
			})
			ServersRemaining[RegionToIndex(region)] = nonMatched

			if KeysError != nil {
				log.Fatalf("Failed to sending remaining users to: " + region)
			} else {
				fmt.Print(region + " just recieved " + strconv.Itoa(keys) + "! (remaining keys: " + strconv.Itoa(SetedAvalibleKeys) + ")\n\n")
			}

			if counter == 4 {
				break
			} else {
				counter++
			}
		}

		fmt.Println("Iteration completed successfully!")

		if flag {
			values[2]++
		}

		var breakflag bool = true
		for i := 0; i < 4; i++ {
			if ServersRemaining[i] > 0 {
				breakflag = false
			}
		}
		if breakflag {
			fmt.Println("The demand is satisfied!")
			break
		}
	}
}
