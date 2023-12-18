package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	pb "github.com/MauricioCortesRo/Lab2_Proto"
	"google.golang.org/grpc"
)

var ServidorIP_puerto string = "dist006:9000"
var lineToRead = 1
var path = "names.txt"

func generateStatus() string {
	rand.Seed(time.Now().UnixNano())
	probability := rand.Float64()
	if probability <= 0.55 {
		return "Infected"
	} else {
		return "Dead"
	}
}

func ReadFile(filename string, lines chan string, serviceClient pb.Information_TradesClient) {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file" + err.Error())
	}
	defer file.Close()

	lineCount := 0

	scanner := bufio.NewScanner(file)
	i := 0
	for i < 5 && scanner.Scan() {
		line := scanner.Text()
		lineCount++
		if lineCount%4 == lineToRead {
			parts := strings.Split(line, " ")
			name := parts[0]
			lastname := parts[1]
			status := generateStatus()
			_, err := serviceClient.Notificate(context.Background(), &pb.Continent{
				Name:     name,
				LastName: lastname,
				Status:   status,
			})
			i++

			if err != nil {
				panic("Person not sended" + err.Error())
			}
			fmt.Printf("Status Sended: %s %s %s\n", name, lastname, status)
		}
	}
	for scanner.Scan() {
		time.Sleep(3 * time.Second)
		lineCount++
		if lineCount%4 == lineToRead {
			line := scanner.Text()
			parts := strings.Split(line, " ")
			name := parts[0]
			lastname := parts[1]
			status := generateStatus()
			_, err := serviceClient.Notificate(context.Background(), &pb.Continent{
				Name:     name,
				LastName: lastname,
				Status:   status,
			})

			if err != nil {
				panic("Person not sended\n" + err.Error())
			}
			fmt.Printf("Status Sended: %s %s %s\n", name, lastname, status)
		}
	}
	close(lines)
}

func main() {

	conn, err := grpc.Dial(ServidorIP_puerto, grpc.WithInsecure())

	if err != nil {
		panic("cannot connect to server" + err.Error())
	}

	serviceClient := pb.NewInformation_TradesClient(conn)

	go ReadFile(path, make(chan string), serviceClient)

	select {}
}
