package test
func main {
	fmt.Println("Starting Album Service Client")
	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	requestOpts = grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
	}
	client = pb.NewAlbumServiceClient(conn)
	//call the respo methods here
}
