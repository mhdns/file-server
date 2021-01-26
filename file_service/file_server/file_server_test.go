package file_test

import (
	"testing"

	"github.com/mhdns/web_server/file_service/file_pb"
	file "github.com/mhdns/web_server/file_service/file_server"
)

func TestFileUpload(t *testing.T) {
	_ = file.NewFileServer("/Users/anas/go/src/github.com/mhdns/web_server/tmp")

	_ = &file_pb.FileUploadRequest{
		Data: &file_pb.FileUploadRequest_Metadata{
			Metadata: &file_pb.Metadata{
				Filename: "test_file.html",
				Filetype: "html",
				Filesize: 50,
			}},
	}
}

func TestFileDelete(t *testing.T) {

}
