package api

import (
	"context"
	"fmt"

	"github.com/mesg-foundation/core/database/services"
	"github.com/mesg-foundation/core/execution"
	"github.com/mesg-foundation/core/grpcclient"
	"github.com/mesg-foundation/core/service"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

const (
	RESOLVER_SERVICE_ID = "aoehiaegjbkejfbq"
	RESOLVER_TASK_KEY   = "resolve"
)

// taskExecutor provides functionalities to execute a MESG task.
type taskExecutor struct {
	api *API
}

// newTaskExecutor creates a new taskExecutor with given api.
func newTaskExecutor(api *API) *taskExecutor {
	return &taskExecutor{
		api: api,
	}
}

// ExecuteTask executes a task tasKey with inputData and tags for service serviceID.
func (e *taskExecutor) Execute(serviceID, taskKey string, inputData map[string]interface{},
	tags []string) (executionID string, err error) {
	s, err := services.Get(serviceID)
	if _, ok := err.(services.NotFound); ok {
		address, err := e.resolve(serviceID)
		if err != nil {
			return "", err
		}
		res, err := e.delegateExecution(address, serviceID, taskKey, inputData)
		return res.ExecutionID, err
		return "", err
	}
	if err != nil {
		return "", err
	}
	s, err = service.FromService(s, service.ContainerOption(e.api.container))
	if err != nil {
		return "", err
	}
	if err := e.checkServiceStatus(s); err != nil {
		return "", err
	}
	return e.execute(s, taskKey, inputData, tags)
}

// checkServiceStatus checks service status. A task should be executed only if
// task's service is running.
func (e *taskExecutor) checkServiceStatus(s *service.Service) error {
	status, err := s.Status()
	if err != nil {
		return err
	}
	if status != service.RUNNING {
		return &NotRunningServiceError{ServiceID: s.ID}
	}
	return nil
}

// execute executes task.
func (e *taskExecutor) execute(s *service.Service, taskKey string, taskInputs map[string]interface{},
	tags []string) (executionID string, err error) {
	exc, err := execution.Create(s, taskKey, taskInputs, tags)
	if err != nil {
		return "", err
	}
	return exc.ID, exc.Execute()
}

func (e *taskExecutor) resolve(serviceID string) (string, error) {
	tags := []string{uuid.NewV4().String()}
	s, err := services.Get(RESOLVER_SERVICE_ID)
	if err != nil {
		return "", err
	}
	ln, err := e.api.ListenResult(RESOLVER_SERVICE_ID, ListenResultTagFilters(tags))
	if err != nil {
		return "", err
	}
	defer ln.Close()

	_, err = e.execute(s, RESOLVER_TASK_KEY, map[string]interface{}{
		"serviceID": serviceID,
	}, tags)
	if err != nil {
		return "", err
	}
	for {
		select {
		case err := <-ln.Err:
			return "", err
		case execution := <-ln.Executions:
			return execution.OutputData["address"].(string), nil
		}
	}
}

func (e *taskExecutor) delegateExecution(address, serviceID, taskKey string, inputs map[string]interface{}) (*grpcclient.ExecuteTaskReply, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := grpcclient.NewCoreClient(conn)
	return client.ExecuteTask(context.Background(), &grpcclient.ExecuteTaskRequest{})
}

// NotRunningServiceError is an error returned when the service is not running that
// a task needed to be executed on.
type NotRunningServiceError struct {
	ServiceID string
}

func (e *NotRunningServiceError) Error() string {
	return fmt.Sprintf("Service %q is not running", e.ServiceID)
}
