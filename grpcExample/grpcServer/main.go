package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

const portNumber = "9000"

func main() {
	// net 표준라이브러리 패키지를 사용해 어떤 네트워크에 어떤 포트로 서버 실행시킬지 정의
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	// NewServer() 함수 호출해 gRPC 서버 생성
	grpcServer := grpc.NewServer()
	log.Printf("start gRPC server on %s port", portNumber)

	// net 사용해 만든 listener connection을 Server()함수의 인자로 넣음
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %s", err)
	}
}
