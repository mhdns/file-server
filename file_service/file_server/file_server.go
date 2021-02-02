package file

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
	"path"

	"github.com/mhdns/web_server/file_service/file_pb"
)

// Server struct to implement grpc methods
type Server struct {
	file_pb.UnimplementedFileServiceServer
	homeDirectory string
	maxFileSize int32
}

// NewFileServer will return a pointer to a Server
func NewFileServer(homeDirectory string, maxFileSize int32) *Server {
	return &Server{
		homeDirectory: homeDirectory,
		maxFileSize: maxFileSize,
	}
}

// FileUpload is a client streaming rpc to upload a file
func (server *Server) FileUpload(req file_pb.FileService_FileUploadServer) error {
	fileMetadata, err := req.Recv()
	if err != nil {
		return status.Errorf(codes.Internal, "unable to receive request: %v", err)
	}

	metadata := fileMetadata.GetMetadata()
	if metadata == nil {
		return status.Error(codes.InvalidArgument, "did not receive metadata")
	}

	filename := metadata.GetFilename()
	filesize := metadata.GetFilesize()

	// find folder with users email address
		// TODO
	userDirectory := path.Join(server.homeDirectory, "user")
	userDirectoryExists, err := exists(userDirectory)
	if err != nil {
		return status.Errorf(codes.Internal, "unable to find user directory %v", err)
	}
	// verify that the owner is the same as requester
		// TODO
	// create folder if folder not present
	if !userDirectoryExists {
		err = os.Mkdir(userDirectory, os.ModePerm)
		if err != nil {
			return status.Errorf(codes.Internal, "unable to create folder: %v", err)
		}
	}
	// check if filesize is withing limit from metadata
	if filesize > server.maxFileSize {
		return status.Error(codes.ResourceExhausted, "exceeds max file size")
	}

	// receive chunk of data and write to file
	fileBuffer := bytes.Buffer{}
	sizeUploaded := 0

	for {
		chunk, err := req.Recv()
		fileChunk := chunk.GetChunk()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "unable to upload file: %v", err)
		}
		if fileChunk == nil {
			return status.Error(codes.InvalidArgument, "did not receive file data")
		}

		n, err := fileBuffer.Write(fileChunk.GetChunk())
		if err != nil {
			return status.Errorf(codes.Internal, "unable to upload file: %v", err)
		}

		sizeUploaded += n
	}

	file, err := os.Create(path.Join(userDirectory, filename))
	if err != nil {
		return status.Errorf(codes.Internal, "unable to create file: %v", err)
	}

	_, err = fileBuffer.WriteTo(file)
	if err != nil {
		return status.Errorf(codes.Internal, "unable to create file: %v", err)
	}
	// return response and close
	res := &file_pb.FileUploadResponse{
		Filename: filename,
		Status: file_pb.Status_SUCCESS,
	}
	err = req.SendAndClose(res)
	if err != nil {
		return status.Errorf(codes.Internal, "unable to send response: %v", err)
	}
	return nil
}

// FileDelete is a unary rpc method to delete an existing file
func (server *Server) FileDelete(
	ctx context.Context,
	req *file_pb.FileDeleteRequest,) (*file_pb.FileDeleteResponse, error) {
	return nil, fmt.Errorf("unimplemented")
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}