package stream

import (
	"fmt"
	"github.com/cshep4/grpc-course/module3-exercise/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
)

const bufSize = 5 * 1024

type StreamingService struct {
	*proto.UnimplementedFileUploadServiceServer
}

func New() *StreamingService {
	return &StreamingService{}
}

func (s StreamingService) DownloadFile(req *proto.DownloadFileRequest, stream proto.FileUploadService_DownloadFileServer) error {
	filename := req.GetName()
	if filename == "" {
		return status.Error(codes.InvalidArgument, "file name cannot be empty")
	}

	fd, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return status.Error(codes.NotFound, fmt.Sprintf("cannot open the file %s", filename))
		}
		return err
	}
	defer fd.Close()

	// read file in chunks and send them to stream
	buf := make([]byte, bufSize)
	var n int
	for {
		n, err = fd.Read(buf)
		fmt.Println("service read n:", n)
		if err != nil {
			if err == io.EOF {
				break
			}
			// return err
			return status.Error(codes.Internal, fmt.Sprintf("error reading file: %v", err))
		}
		if n == 0 {
			break
		}
		res := &proto.DownloadFileResponse{
			Content: buf[:n],
		}
		err = stream.Send(res)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s StreamingService) UploadFile(stream proto.FileUploadService_UploadFileServer) error {
	// your implementation goes here ...
	panic("implement me")
}
