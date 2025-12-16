// Komunikacija po protokolu gRPC
// odjemalec

package main

import (
	"api/grpc/protobufStorage"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Client(url string) {
	// vzpostavimo povezavo s strežnikom
	fmt.Printf("gRPC client connecting to %v\n", url)
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// vzpostavimo izvajalno okolje
	contextCRUD, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// vzpostavimo vmesnik gRPC
	grpcClient := protobufStorage.NewCRUDClient(conn)

	// pripravimo strukture, ki jih uporabljamo kot argumente pri klicu oddaljenih metod
	lecturesCreate := protobufStorage.Todo{Task: "predavanja", Completed: false}
	lecturesUpdate := protobufStorage.Todo{Task: "predavanja", Completed: true}
	practicals := protobufStorage.Todo{Task: "vaje", Completed: false}

	done := make(chan bool)

	go func() {
		if stream, err := grpcClient.Subscribe(contextCRUD, &emptypb.Empty{}); err != nil {
			panic(err)
		} else {
			for {
				todoEvent, err := stream.Recv()
				if err != nil {
					fmt.Println("No more action.")
					return
				}
				fmt.Printf("-> Event: %v %v\n", todoEvent.T.GetTask(), todoEvent.Action)
			}
		}
	}()

	go func() {
		// ustvarimo zapis
		fmt.Print("1. Create: ")
		if _, err := grpcClient.Create(contextCRUD, &lecturesCreate); err != nil {
			panic(err)
		}
		fmt.Println("done")

		time.Sleep(time.Second)

		// ustvarimo zapis
		fmt.Print("2. Create: ")
		if _, err := grpcClient.Create(contextCRUD, &practicals); err != nil {
			panic(err)
		}
		fmt.Println("done")

		time.Sleep(time.Second)

		// posodobimo zapis
		fmt.Print("3. Update: ")
		if _, err := grpcClient.Update(contextCRUD, &lecturesUpdate); err != nil {
			panic(err)
		}
		fmt.Println("done")

		time.Sleep(time.Second)

		// izbrišemo zapis
		fmt.Print("4. Delete: ")
		if _, err := grpcClient.Delete(contextCRUD, &practicals); err != nil {
			panic(err)
		}
		fmt.Println("done")

		time.Sleep(time.Second)

		// izbrišemo zapis
		fmt.Print("5. Delete: ")
		if _, err := grpcClient.Delete(contextCRUD, &lecturesUpdate); err != nil {
			panic(err)
		}
		fmt.Println("done")

		time.Sleep(time.Second)

		done <- true
	}()

	// Wait for operations to complete
	<-done
}
