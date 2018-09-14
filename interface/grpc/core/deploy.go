package core

import (
	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/grpcclient"
	service "github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/service/importer"
)

// DeployService deploys a service from Git URL or service.tar.gz file. It'll send status
// events during the process and finish with sending service id or validation error.
func (s *Server) DeployService(stream grpcclient.Core_DeployServiceServer) error {
	statuses := make(chan api.DeployStatus, 0)
	go sendDeployStatus(statuses, stream)

	var (
		service         *service.Service
		validationError *importer.ValidationError
		err             error
	)

	sr := newDeployServiceStreamReader(stream)
	url, err := sr.GetURL()
	if err != nil {
		return err
	}
	if url != "" {
		service, validationError, err = s.api.DeployServiceFromURL(url, api.DeployServiceStatusOption(statuses))
	} else {
		service, validationError, err = s.api.DeployService(sr, api.DeployServiceStatusOption(statuses))
	}

	if err != nil {
		return err
	}
	if validationError != nil {
		return stream.Send(&grpcclient.DeployServiceReply{
			Value: &grpcclient.DeployServiceReply_ValidationError{ValidationError: validationError.Error()},
		})
	}

	return stream.Send(&grpcclient.DeployServiceReply{
		Value: &grpcclient.DeployServiceReply_ServiceID{ServiceID: service.ID},
	})
}

func sendDeployStatus(statuses chan api.DeployStatus, stream grpcclient.Core_DeployServiceServer) {
	for status := range statuses {
		var typ grpcclient.DeployServiceReply_Status_Type
		switch status.Type {
		case api.RUNNING:
			typ = grpcclient.DeployServiceReply_Status_RUNNING
		case api.DONE:
			typ = grpcclient.DeployServiceReply_Status_DONE
		}

		stream.Send(&grpcclient.DeployServiceReply{
			Value: &grpcclient.DeployServiceReply_Status_{
				Status: &grpcclient.DeployServiceReply_Status{
					Message: status.Message,
					Type:    typ,
				},
			},
		})
	}
}

type deployServiceStreamReader struct {
	stream grpcclient.Core_DeployServiceServer

	data []byte
	i    int64
}

func newDeployServiceStreamReader(stream grpcclient.Core_DeployServiceServer) *deployServiceStreamReader {
	return &deployServiceStreamReader{stream: stream}
}

func (r *deployServiceStreamReader) GetURL() (url string, err error) {
	message, err := r.stream.Recv()
	if err != nil {
		return "", err
	}
	r.data = message.GetChunk()
	return message.GetUrl(), err
}

func (r *deployServiceStreamReader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.data)) {
		message, err := r.stream.Recv()
		if err != nil {
			return 0, err
		}
		r.data = message.GetChunk()
		r.i = 0
		return r.Read(p)
	}
	n = copy(p, r.data[r.i:])
	r.i += int64(n)
	return n, nil
}
