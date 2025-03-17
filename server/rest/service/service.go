package service

import (
	"context"
	"flash/calculator/pb"
	"log"

	"google.golang.org/grpc"
)

var client pb.OperationsClient

func prepareClient() {
	if client != nil {
		return
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	client = pb.NewOperationsClient(conn)
}

func Add(firstNum, secondNum int) (int, error) {
	prepareClient()
	ctx := context.Background()
	resp, err := client.Add(ctx, &pb.NumberRequest{
		FirstNum:  int32(firstNum),
		SecondNum: int32(secondNum),
	})
	if err != nil {
		log.Fatalf("Failed to add: %v", err)
		return 0, err
	}
	
	return int(resp.Result), nil
}

func Subtract(firstNum, secondNum int) (int, error) {
	prepareClient()
	ctx := context.Background()
	resp, err := client.Subtract(ctx, &pb.NumberRequest{
		FirstNum:  int32(firstNum),
		SecondNum: int32(secondNum),
	})
	if err != nil {
		log.Fatalf("Failed to subtract: %v", err)
		return 0, err
	}
	
	return int(resp.Result), nil
}

func Multiply(firstNum, secondNum int) (int, error) {
	prepareClient()
	ctx := context.Background()
	resp, err := client.Multiply(ctx, &pb.NumberRequest{
		FirstNum:  int32(firstNum),
		SecondNum: int32(secondNum),
	})
	if err != nil {
		log.Fatalf("Failed to multiply: %v", err)
		return 0, err
	}
	
	return int(resp.Result), nil
}

func Division(firstNum, secondNum int) (int, error) {
	prepareClient()
	ctx := context.Background()
	resp, err := client.Division(ctx, &pb.NumberRequest{
		FirstNum:  int32(firstNum),
		SecondNum: int32(secondNum),
	})
	if err != nil {
		log.Fatalf("Failed to divide: %v", err)
		return 0, err
	}
	
	return int(resp.Result), nil
}

func Power(firstNum, secondNum int) (int, error) {
	prepareClient()
	ctx := context.Background()
	resp, err := client.Power(ctx, &pb.NumberRequest{
		FirstNum:  int32(firstNum),
		SecondNum: int32(secondNum),
	})
	if err != nil {
		log.Fatalf("Failed to power: %v", err)
		return 0, err
	}
	
	return int(resp.Result), nil
}