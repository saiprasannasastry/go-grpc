package client

import (
	"context"
	"fmt"
	albumpb "github.com/crud-grpc/pkg/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
	"io"
	//"encoding/json"
)

const (
	server = "10.196.105.125:50051"
)

type ClientConn interface {
	Close() error
}
type Album struct {
	Title  string `json:"title"`
	UserId string `json:"userId"`
	Id     string `json:"id"`
}

func Connect(serverAddr string) (*grpc.ClientConn, error) {
	log.Infof("trying to connect to %s", serverAddr)
	dialOpt := grpc.WithInsecure()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, serverAddr,
		dialOpt,
		grpc.WithBlock(),
	)

	if err != nil {
		log.Errorf("unable to create connection %v", err)
		return nil, err
	}
	log.Infof("Established connection with %s", serverAddr)
	return conn, nil
}

func NewServer() (albumpb.AlbumServiceClient, ClientConn, error) {
	conn, err := Connect(server)
	if err != nil {
		return nil, conn, err
	}
	return albumpb.NewAlbumServiceClient(conn), conn, nil
}
func GetAlbumRequest() {
	client, conn, err := NewServer()
	if err != nil {
		log.Errorf("NewServer() - error building the grpc client. Reason: %v", err)
		return
	}
	defer conn.Close()
	ctx := metadata.AppendToOutgoingContext(context.Background())
	stream, err := client.ListAlbum(ctx, &albumpb.ListAlbumRequest{})
	if err != nil {
		log.Errorf("GetAlbumRequest() -ListAlbum failed. Reason: %v", err)
		return
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("blah")
			return 
		}
		fmt.Println(res.String())
	}
	resp, err := client.GetAlbum(ctx, &albumpb.Albumreq{
		Id: "3",
	})
	if err != nil {
		log.Errorf("GetAlbumRequest() - failed. Reason: %v", err)
		return
	}
	fmt.Println(resp)
}
