package file

import (
	"context"
	"fmt"

	"github.com/mhdns/web_server/file_service/file_pb"
)

// Server struct to implement grpc methods
type Server struct {
	file_pb.UnimplementedFileServiceServer
	homeDirectory string
}

// NewFileServer will return a pointer to a Server
func NewFileServer(homeDirectory string) *Server {
	return &Server{
		homeDirectory: homeDirectory,
	}
}

// FileUpload is a client streaming rpc to upload a file
func (server *Server) FileUpload(req file_pb.FileService_FileUploadServer) error {
	// find folder with users email address

	// verify that the owner is the same as requester

	// create folder if folder not present

	// check if filesize is withing limit from metadata

	// recieve chunk of data and write to file

	// return response and close

	return fmt.Errorf("Unimplemented")
}

// FileDelete is a unary rpc method to delete an existing file
func (server *Server) FileDelete(ctx context.Context, req *file_pb.FileDeleteRequest) (*file_pb.FileDeleteResponse, error) {
	return nil, fmt.Errorf("Unimplemented")
}
