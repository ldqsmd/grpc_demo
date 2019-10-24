package main

//func main() {
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//	mux := runtime.NewServeMux()
//	client := gtls.NewClientTLS()
//	creds ,err :=  client.GetTLSCredentials()
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = pb.RegisterSearchServiceHandlerFromEndpoint(ctx, mux, "localhost:9001", []grpc.DialOption{grpc.WithTransportCredentials(creds)}, )
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("rpc  rest-api start")
//	err = http.ListenAndServe(":8899", mux)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
