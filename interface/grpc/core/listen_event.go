package core

import (
	"encoding/json"

	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/grpcclient"
)

// ListenEvent listens events matches with eventFilter on serviceID.
func (s *Server) ListenEvent(request *grpcclient.ListenEventRequest, stream grpcclient.Core_ListenEventServer) error {
	ln, err := s.api.ListenEvent(request.ServiceID, api.ListenEventKeyFilter(request.EventFilter))
	if err != nil {
		return err
	}
	defer ln.Close()

	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case err := <-ln.Err:
			return err

		case ev := <-ln.Events:
			evData, err := json.Marshal(ev.Data)
			if err != nil {
				return err
			}

			if err := stream.Send(&grpcclient.EventData{
				EventKey:  ev.Key,
				EventData: string(evData),
			}); err != nil {
				return err
			}
		}
	}
}
