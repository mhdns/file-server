package file_test

import (
	"context"
	"net"
	"testing"

	"github.com/mhdns/web_server/file_service/file_pb"
	file "github.com/mhdns/web_server/file_service/file_server"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestFileUpload(t *testing.T) {
	serverAddr := createFileServer(t, "/Users/anas/go/src/github.com/mhdns/web_server/tmp")
	client := createFileClient(t, serverAddr)

	stream, err := client.FileUpload(context.Background())
	require.NoError(t, err)

	req := &file_pb.FileUploadRequest{
		Data: &file_pb.FileUploadRequest_Metadata{
			Metadata: &file_pb.Metadata{
				Filename: "test_file.html",
				Filetype: "html",
				Filesize: 50,
			}},
	}

	err = stream.Send(req)
	require.NoError(t, err)

	_, err = stream.CloseAndRecv()
	require.NoError(t, err)
}

func TestFileDelete(t *testing.T) {

}

func createFileClient(t *testing.T, addrString string) file_pb.FileServiceClient {
	conn, err := grpc.Dial(addrString, grpc.WithInsecure())
	require.NoError(t, err)

	return file_pb.NewFileServiceClient(conn)
}

func createFileServer(t *testing.T, homeDirectory string) string {
	server := file.NewFileServer(homeDirectory)

	grpcServer := grpc.NewServer()

	file_pb.RegisterFileServiceServer(grpcServer, server)

	li, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(li)
	return li.Addr().String()
}
