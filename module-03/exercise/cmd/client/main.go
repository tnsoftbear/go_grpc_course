package main

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net/http"

	"github.com/cshep4/grpc-course/module3-exercise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// initialise a gRPC connection on server start
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewFileUploadServiceClient(conn)

	http.HandleFunc("/", downloadHandler(client))

	log.Printf("starting http server on address: %s", ":8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// downloadHandler is an example of a gRPC client making a request to a server streaming RPC.
// The gRPC call will stream a file in chunks back to the client.
// The file content will be buffered until the server stream is complete, then the content will be returned to the user.
func downloadHandler(client proto.FileUploadServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Query().Get("file")
		if file == "" {
			file = "gopher.png"
		}
		fmt.Println("File to download:", file)
		ctx := context.Background()
		req := &proto.DownloadFileRequest{
			Name: file,
		}
		downloadFileClient, err := client.DownloadFile(ctx, req)
		if err != nil {
			st := status.Convert(err)
			switch st.Code() {
			case codes.NotFound:
				http.Error(w, "File not found", 404)
				return
			case codes.InvalidArgument:
				http.Error(w, "Bad request", 400)
				return
			}
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println("DownloadFile server streaming RPC started:", file)

		contentBuf := bytes.NewBuffer([]byte{})
		for {
			res, err := downloadFileClient.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				http.Error(w, err.Error(), 500)
				return
			}
			content := res.GetContent()
			contentBuf.Write(content)
			fmt.Println("client Recv() content len:", len(content))
		}
		
		wrote, err := w.Write(contentBuf.Bytes())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println("Wrote content of len:", wrote)
	}
}
