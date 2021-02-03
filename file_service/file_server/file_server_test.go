package file_test

import (
	"bufio"
	"context"
	"io"
	"net"
	"os"
	"testing"

	"github.com/mhdns/web_server/file_service/file_pb"
	file "github.com/mhdns/web_server/file_service/file_server"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestFileUpload(t *testing.T) {
	serverAddr := createFileServer(t, "/Users/anas/go/src/github.com/mhdns/web_server/file_service/tmp_destination")
	client := createFileClient(t, serverAddr)

	stream, err := client.FileUpload(context.Background())
	require.NoError(t, err)

	file, err := os.Open("/Users/anas/go/src/github.com/mhdns/web_server/tmp_source/test_file")
	require.NoError(t, err)
	defer file.Close()

	fileInfo, err := file.Stat()
	require.NoError(t, err)

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	req := &file_pb.FileUploadRequest{
		Data: &file_pb.FileUploadRequest_Metadata{
			Metadata: &file_pb.Metadata{
				Filename: fileInfo.Name(),
				Filetype: "",
				Filesize: int32(fileInfo.Size()),
			}},
	}

	err = stream.Send(req)
	require.NoError(t, err)

	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		require.NoError(t, err)

		chunk := &file_pb.FileUploadRequest{
			Data: &file_pb.FileUploadRequest_Chunk{
				Chunk: &file_pb.File{
					Chunk: buffer,
				},
			},
		}

		err = stream.Send(chunk)
		require.NoError(t, err)
	}

	_, err = stream.CloseAndRecv()
	require.NoError(t, err)
}

func TestFileDelete(t *testing.T) {
	serverAddr := createFileServer(t, "/Users/anas/go/src/github.com/mhdns/web_server/tmp_source")
	clientServer := createFileClient(t, serverAddr)
	filePath := ""
	req := &file_pb.FileDeleteRequest{
		Filename: filePath,
	}
	res, err := clientServer.FileDelete(context.Background(), req)
	require.NoError(t,err)
	require.NotEmpty(t, res)
}

func createFileClient(t *testing.T, addrString string) file_pb.FileServiceClient {
	conn, err := grpc.Dial(addrString, grpc.WithInsecure())
	require.NoError(t, err)

	return file_pb.NewFileServiceClient(conn)
}

func createFileServer(t *testing.T, homeDirectory string) string {
	server := file.NewFileServer(homeDirectory, 1000)

	grpcServer := grpc.NewServer()

	file_pb.RegisterFileServiceServer(grpcServer, server)

	li, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(li)
	return li.Addr().String()
}
