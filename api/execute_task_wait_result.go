package api

import (
	"github.com/mesg-foundation/core/execution"
	uuid "github.com/satori/go.uuid"
)

// ExecuteTaskAndWaitResult executes given task and listen for result.
func (a *API) ExecuteTaskAndWaitResult(serviceID, task string, inputs map[string]interface{}) (*execution.Execution, error) {
	tag := uuid.NewV4().String()
	result, err := a.ListenResult(serviceID, ListenResultTagFilters([]string{tag}))
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if _, err := a.ExecuteTask(serviceID, task, inputs, []string{tag}); err != nil {
		return nil, err
	}
	select {
	case execution := <-result.Executions:
		return execution, nil
	case err := <-result.Err:
		return nil, err
	}
}
