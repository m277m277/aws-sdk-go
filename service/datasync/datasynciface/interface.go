// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package datasynciface provides an interface to enable mocking the AWS DataSync service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package datasynciface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/datasync"
)

// DataSyncAPI provides an interface to enable mocking the
// datasync.DataSync service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS DataSync.
//    func myFunc(svc datasynciface.DataSyncAPI) bool {
//        // Make svc.CancelTaskExecution request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := datasync.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDataSyncClient struct {
//        datasynciface.DataSyncAPI
//    }
//    func (m *mockDataSyncClient) CancelTaskExecution(input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDataSyncClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DataSyncAPI interface {
	CancelTaskExecution(*datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error)
	CancelTaskExecutionWithContext(aws.Context, *datasync.CancelTaskExecutionInput, ...request.Option) (*datasync.CancelTaskExecutionOutput, error)
	CancelTaskExecutionRequest(*datasync.CancelTaskExecutionInput) (*request.Request, *datasync.CancelTaskExecutionOutput)

	CreateAgent(*datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error)
	CreateAgentWithContext(aws.Context, *datasync.CreateAgentInput, ...request.Option) (*datasync.CreateAgentOutput, error)
	CreateAgentRequest(*datasync.CreateAgentInput) (*request.Request, *datasync.CreateAgentOutput)

	CreateLocationEfs(*datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error)
	CreateLocationEfsWithContext(aws.Context, *datasync.CreateLocationEfsInput, ...request.Option) (*datasync.CreateLocationEfsOutput, error)
	CreateLocationEfsRequest(*datasync.CreateLocationEfsInput) (*request.Request, *datasync.CreateLocationEfsOutput)

	CreateLocationFsxLustre(*datasync.CreateLocationFsxLustreInput) (*datasync.CreateLocationFsxLustreOutput, error)
	CreateLocationFsxLustreWithContext(aws.Context, *datasync.CreateLocationFsxLustreInput, ...request.Option) (*datasync.CreateLocationFsxLustreOutput, error)
	CreateLocationFsxLustreRequest(*datasync.CreateLocationFsxLustreInput) (*request.Request, *datasync.CreateLocationFsxLustreOutput)

	CreateLocationFsxOntap(*datasync.CreateLocationFsxOntapInput) (*datasync.CreateLocationFsxOntapOutput, error)
	CreateLocationFsxOntapWithContext(aws.Context, *datasync.CreateLocationFsxOntapInput, ...request.Option) (*datasync.CreateLocationFsxOntapOutput, error)
	CreateLocationFsxOntapRequest(*datasync.CreateLocationFsxOntapInput) (*request.Request, *datasync.CreateLocationFsxOntapOutput)

	CreateLocationFsxOpenZfs(*datasync.CreateLocationFsxOpenZfsInput) (*datasync.CreateLocationFsxOpenZfsOutput, error)
	CreateLocationFsxOpenZfsWithContext(aws.Context, *datasync.CreateLocationFsxOpenZfsInput, ...request.Option) (*datasync.CreateLocationFsxOpenZfsOutput, error)
	CreateLocationFsxOpenZfsRequest(*datasync.CreateLocationFsxOpenZfsInput) (*request.Request, *datasync.CreateLocationFsxOpenZfsOutput)

	CreateLocationFsxWindows(*datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error)
	CreateLocationFsxWindowsWithContext(aws.Context, *datasync.CreateLocationFsxWindowsInput, ...request.Option) (*datasync.CreateLocationFsxWindowsOutput, error)
	CreateLocationFsxWindowsRequest(*datasync.CreateLocationFsxWindowsInput) (*request.Request, *datasync.CreateLocationFsxWindowsOutput)

	CreateLocationHdfs(*datasync.CreateLocationHdfsInput) (*datasync.CreateLocationHdfsOutput, error)
	CreateLocationHdfsWithContext(aws.Context, *datasync.CreateLocationHdfsInput, ...request.Option) (*datasync.CreateLocationHdfsOutput, error)
	CreateLocationHdfsRequest(*datasync.CreateLocationHdfsInput) (*request.Request, *datasync.CreateLocationHdfsOutput)

	CreateLocationNfs(*datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error)
	CreateLocationNfsWithContext(aws.Context, *datasync.CreateLocationNfsInput, ...request.Option) (*datasync.CreateLocationNfsOutput, error)
	CreateLocationNfsRequest(*datasync.CreateLocationNfsInput) (*request.Request, *datasync.CreateLocationNfsOutput)

	CreateLocationObjectStorage(*datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error)
	CreateLocationObjectStorageWithContext(aws.Context, *datasync.CreateLocationObjectStorageInput, ...request.Option) (*datasync.CreateLocationObjectStorageOutput, error)
	CreateLocationObjectStorageRequest(*datasync.CreateLocationObjectStorageInput) (*request.Request, *datasync.CreateLocationObjectStorageOutput)

	CreateLocationS3(*datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error)
	CreateLocationS3WithContext(aws.Context, *datasync.CreateLocationS3Input, ...request.Option) (*datasync.CreateLocationS3Output, error)
	CreateLocationS3Request(*datasync.CreateLocationS3Input) (*request.Request, *datasync.CreateLocationS3Output)

	CreateLocationSmb(*datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error)
	CreateLocationSmbWithContext(aws.Context, *datasync.CreateLocationSmbInput, ...request.Option) (*datasync.CreateLocationSmbOutput, error)
	CreateLocationSmbRequest(*datasync.CreateLocationSmbInput) (*request.Request, *datasync.CreateLocationSmbOutput)

	CreateTask(*datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error)
	CreateTaskWithContext(aws.Context, *datasync.CreateTaskInput, ...request.Option) (*datasync.CreateTaskOutput, error)
	CreateTaskRequest(*datasync.CreateTaskInput) (*request.Request, *datasync.CreateTaskOutput)

	DeleteAgent(*datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error)
	DeleteAgentWithContext(aws.Context, *datasync.DeleteAgentInput, ...request.Option) (*datasync.DeleteAgentOutput, error)
	DeleteAgentRequest(*datasync.DeleteAgentInput) (*request.Request, *datasync.DeleteAgentOutput)

	DeleteLocation(*datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error)
	DeleteLocationWithContext(aws.Context, *datasync.DeleteLocationInput, ...request.Option) (*datasync.DeleteLocationOutput, error)
	DeleteLocationRequest(*datasync.DeleteLocationInput) (*request.Request, *datasync.DeleteLocationOutput)

	DeleteTask(*datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error)
	DeleteTaskWithContext(aws.Context, *datasync.DeleteTaskInput, ...request.Option) (*datasync.DeleteTaskOutput, error)
	DeleteTaskRequest(*datasync.DeleteTaskInput) (*request.Request, *datasync.DeleteTaskOutput)

	DescribeAgent(*datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error)
	DescribeAgentWithContext(aws.Context, *datasync.DescribeAgentInput, ...request.Option) (*datasync.DescribeAgentOutput, error)
	DescribeAgentRequest(*datasync.DescribeAgentInput) (*request.Request, *datasync.DescribeAgentOutput)

	DescribeLocationEfs(*datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error)
	DescribeLocationEfsWithContext(aws.Context, *datasync.DescribeLocationEfsInput, ...request.Option) (*datasync.DescribeLocationEfsOutput, error)
	DescribeLocationEfsRequest(*datasync.DescribeLocationEfsInput) (*request.Request, *datasync.DescribeLocationEfsOutput)

	DescribeLocationFsxLustre(*datasync.DescribeLocationFsxLustreInput) (*datasync.DescribeLocationFsxLustreOutput, error)
	DescribeLocationFsxLustreWithContext(aws.Context, *datasync.DescribeLocationFsxLustreInput, ...request.Option) (*datasync.DescribeLocationFsxLustreOutput, error)
	DescribeLocationFsxLustreRequest(*datasync.DescribeLocationFsxLustreInput) (*request.Request, *datasync.DescribeLocationFsxLustreOutput)

	DescribeLocationFsxOntap(*datasync.DescribeLocationFsxOntapInput) (*datasync.DescribeLocationFsxOntapOutput, error)
	DescribeLocationFsxOntapWithContext(aws.Context, *datasync.DescribeLocationFsxOntapInput, ...request.Option) (*datasync.DescribeLocationFsxOntapOutput, error)
	DescribeLocationFsxOntapRequest(*datasync.DescribeLocationFsxOntapInput) (*request.Request, *datasync.DescribeLocationFsxOntapOutput)

	DescribeLocationFsxOpenZfs(*datasync.DescribeLocationFsxOpenZfsInput) (*datasync.DescribeLocationFsxOpenZfsOutput, error)
	DescribeLocationFsxOpenZfsWithContext(aws.Context, *datasync.DescribeLocationFsxOpenZfsInput, ...request.Option) (*datasync.DescribeLocationFsxOpenZfsOutput, error)
	DescribeLocationFsxOpenZfsRequest(*datasync.DescribeLocationFsxOpenZfsInput) (*request.Request, *datasync.DescribeLocationFsxOpenZfsOutput)

	DescribeLocationFsxWindows(*datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error)
	DescribeLocationFsxWindowsWithContext(aws.Context, *datasync.DescribeLocationFsxWindowsInput, ...request.Option) (*datasync.DescribeLocationFsxWindowsOutput, error)
	DescribeLocationFsxWindowsRequest(*datasync.DescribeLocationFsxWindowsInput) (*request.Request, *datasync.DescribeLocationFsxWindowsOutput)

	DescribeLocationHdfs(*datasync.DescribeLocationHdfsInput) (*datasync.DescribeLocationHdfsOutput, error)
	DescribeLocationHdfsWithContext(aws.Context, *datasync.DescribeLocationHdfsInput, ...request.Option) (*datasync.DescribeLocationHdfsOutput, error)
	DescribeLocationHdfsRequest(*datasync.DescribeLocationHdfsInput) (*request.Request, *datasync.DescribeLocationHdfsOutput)

	DescribeLocationNfs(*datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error)
	DescribeLocationNfsWithContext(aws.Context, *datasync.DescribeLocationNfsInput, ...request.Option) (*datasync.DescribeLocationNfsOutput, error)
	DescribeLocationNfsRequest(*datasync.DescribeLocationNfsInput) (*request.Request, *datasync.DescribeLocationNfsOutput)

	DescribeLocationObjectStorage(*datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error)
	DescribeLocationObjectStorageWithContext(aws.Context, *datasync.DescribeLocationObjectStorageInput, ...request.Option) (*datasync.DescribeLocationObjectStorageOutput, error)
	DescribeLocationObjectStorageRequest(*datasync.DescribeLocationObjectStorageInput) (*request.Request, *datasync.DescribeLocationObjectStorageOutput)

	DescribeLocationS3(*datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error)
	DescribeLocationS3WithContext(aws.Context, *datasync.DescribeLocationS3Input, ...request.Option) (*datasync.DescribeLocationS3Output, error)
	DescribeLocationS3Request(*datasync.DescribeLocationS3Input) (*request.Request, *datasync.DescribeLocationS3Output)

	DescribeLocationSmb(*datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error)
	DescribeLocationSmbWithContext(aws.Context, *datasync.DescribeLocationSmbInput, ...request.Option) (*datasync.DescribeLocationSmbOutput, error)
	DescribeLocationSmbRequest(*datasync.DescribeLocationSmbInput) (*request.Request, *datasync.DescribeLocationSmbOutput)

	DescribeTask(*datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error)
	DescribeTaskWithContext(aws.Context, *datasync.DescribeTaskInput, ...request.Option) (*datasync.DescribeTaskOutput, error)
	DescribeTaskRequest(*datasync.DescribeTaskInput) (*request.Request, *datasync.DescribeTaskOutput)

	DescribeTaskExecution(*datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error)
	DescribeTaskExecutionWithContext(aws.Context, *datasync.DescribeTaskExecutionInput, ...request.Option) (*datasync.DescribeTaskExecutionOutput, error)
	DescribeTaskExecutionRequest(*datasync.DescribeTaskExecutionInput) (*request.Request, *datasync.DescribeTaskExecutionOutput)

	ListAgents(*datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error)
	ListAgentsWithContext(aws.Context, *datasync.ListAgentsInput, ...request.Option) (*datasync.ListAgentsOutput, error)
	ListAgentsRequest(*datasync.ListAgentsInput) (*request.Request, *datasync.ListAgentsOutput)

	ListAgentsPages(*datasync.ListAgentsInput, func(*datasync.ListAgentsOutput, bool) bool) error
	ListAgentsPagesWithContext(aws.Context, *datasync.ListAgentsInput, func(*datasync.ListAgentsOutput, bool) bool, ...request.Option) error

	ListLocations(*datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error)
	ListLocationsWithContext(aws.Context, *datasync.ListLocationsInput, ...request.Option) (*datasync.ListLocationsOutput, error)
	ListLocationsRequest(*datasync.ListLocationsInput) (*request.Request, *datasync.ListLocationsOutput)

	ListLocationsPages(*datasync.ListLocationsInput, func(*datasync.ListLocationsOutput, bool) bool) error
	ListLocationsPagesWithContext(aws.Context, *datasync.ListLocationsInput, func(*datasync.ListLocationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *datasync.ListTagsForResourceInput, ...request.Option) (*datasync.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*datasync.ListTagsForResourceInput) (*request.Request, *datasync.ListTagsForResourceOutput)

	ListTagsForResourcePages(*datasync.ListTagsForResourceInput, func(*datasync.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *datasync.ListTagsForResourceInput, func(*datasync.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ListTaskExecutions(*datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error)
	ListTaskExecutionsWithContext(aws.Context, *datasync.ListTaskExecutionsInput, ...request.Option) (*datasync.ListTaskExecutionsOutput, error)
	ListTaskExecutionsRequest(*datasync.ListTaskExecutionsInput) (*request.Request, *datasync.ListTaskExecutionsOutput)

	ListTaskExecutionsPages(*datasync.ListTaskExecutionsInput, func(*datasync.ListTaskExecutionsOutput, bool) bool) error
	ListTaskExecutionsPagesWithContext(aws.Context, *datasync.ListTaskExecutionsInput, func(*datasync.ListTaskExecutionsOutput, bool) bool, ...request.Option) error

	ListTasks(*datasync.ListTasksInput) (*datasync.ListTasksOutput, error)
	ListTasksWithContext(aws.Context, *datasync.ListTasksInput, ...request.Option) (*datasync.ListTasksOutput, error)
	ListTasksRequest(*datasync.ListTasksInput) (*request.Request, *datasync.ListTasksOutput)

	ListTasksPages(*datasync.ListTasksInput, func(*datasync.ListTasksOutput, bool) bool) error
	ListTasksPagesWithContext(aws.Context, *datasync.ListTasksInput, func(*datasync.ListTasksOutput, bool) bool, ...request.Option) error

	StartTaskExecution(*datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error)
	StartTaskExecutionWithContext(aws.Context, *datasync.StartTaskExecutionInput, ...request.Option) (*datasync.StartTaskExecutionOutput, error)
	StartTaskExecutionRequest(*datasync.StartTaskExecutionInput) (*request.Request, *datasync.StartTaskExecutionOutput)

	TagResource(*datasync.TagResourceInput) (*datasync.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *datasync.TagResourceInput, ...request.Option) (*datasync.TagResourceOutput, error)
	TagResourceRequest(*datasync.TagResourceInput) (*request.Request, *datasync.TagResourceOutput)

	UntagResource(*datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *datasync.UntagResourceInput, ...request.Option) (*datasync.UntagResourceOutput, error)
	UntagResourceRequest(*datasync.UntagResourceInput) (*request.Request, *datasync.UntagResourceOutput)

	UpdateAgent(*datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error)
	UpdateAgentWithContext(aws.Context, *datasync.UpdateAgentInput, ...request.Option) (*datasync.UpdateAgentOutput, error)
	UpdateAgentRequest(*datasync.UpdateAgentInput) (*request.Request, *datasync.UpdateAgentOutput)

	UpdateLocationHdfs(*datasync.UpdateLocationHdfsInput) (*datasync.UpdateLocationHdfsOutput, error)
	UpdateLocationHdfsWithContext(aws.Context, *datasync.UpdateLocationHdfsInput, ...request.Option) (*datasync.UpdateLocationHdfsOutput, error)
	UpdateLocationHdfsRequest(*datasync.UpdateLocationHdfsInput) (*request.Request, *datasync.UpdateLocationHdfsOutput)

	UpdateLocationNfs(*datasync.UpdateLocationNfsInput) (*datasync.UpdateLocationNfsOutput, error)
	UpdateLocationNfsWithContext(aws.Context, *datasync.UpdateLocationNfsInput, ...request.Option) (*datasync.UpdateLocationNfsOutput, error)
	UpdateLocationNfsRequest(*datasync.UpdateLocationNfsInput) (*request.Request, *datasync.UpdateLocationNfsOutput)

	UpdateLocationObjectStorage(*datasync.UpdateLocationObjectStorageInput) (*datasync.UpdateLocationObjectStorageOutput, error)
	UpdateLocationObjectStorageWithContext(aws.Context, *datasync.UpdateLocationObjectStorageInput, ...request.Option) (*datasync.UpdateLocationObjectStorageOutput, error)
	UpdateLocationObjectStorageRequest(*datasync.UpdateLocationObjectStorageInput) (*request.Request, *datasync.UpdateLocationObjectStorageOutput)

	UpdateLocationSmb(*datasync.UpdateLocationSmbInput) (*datasync.UpdateLocationSmbOutput, error)
	UpdateLocationSmbWithContext(aws.Context, *datasync.UpdateLocationSmbInput, ...request.Option) (*datasync.UpdateLocationSmbOutput, error)
	UpdateLocationSmbRequest(*datasync.UpdateLocationSmbInput) (*request.Request, *datasync.UpdateLocationSmbOutput)

	UpdateTask(*datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error)
	UpdateTaskWithContext(aws.Context, *datasync.UpdateTaskInput, ...request.Option) (*datasync.UpdateTaskOutput, error)
	UpdateTaskRequest(*datasync.UpdateTaskInput) (*request.Request, *datasync.UpdateTaskOutput)

	UpdateTaskExecution(*datasync.UpdateTaskExecutionInput) (*datasync.UpdateTaskExecutionOutput, error)
	UpdateTaskExecutionWithContext(aws.Context, *datasync.UpdateTaskExecutionInput, ...request.Option) (*datasync.UpdateTaskExecutionOutput, error)
	UpdateTaskExecutionRequest(*datasync.UpdateTaskExecutionInput) (*request.Request, *datasync.UpdateTaskExecutionOutput)
}

var _ DataSyncAPI = (*datasync.DataSync)(nil)
