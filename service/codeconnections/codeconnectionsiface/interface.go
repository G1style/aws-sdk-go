// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codeconnectionsiface provides an interface to enable mocking the AWS CodeConnections service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codeconnectionsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codeconnections"
)

// CodeConnectionsAPI provides an interface to enable mocking the
// codeconnections.CodeConnections service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS CodeConnections.
//	func myFunc(svc codeconnectionsiface.CodeConnectionsAPI) bool {
//	    // Make svc.CreateConnection request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := codeconnections.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCodeConnectionsClient struct {
//	    codeconnectionsiface.CodeConnectionsAPI
//	}
//	func (m *mockCodeConnectionsClient) CreateConnection(input *codeconnections.CreateConnectionInput) (*codeconnections.CreateConnectionOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCodeConnectionsClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CodeConnectionsAPI interface {
	CreateConnection(*codeconnections.CreateConnectionInput) (*codeconnections.CreateConnectionOutput, error)
	CreateConnectionWithContext(aws.Context, *codeconnections.CreateConnectionInput, ...request.Option) (*codeconnections.CreateConnectionOutput, error)
	CreateConnectionRequest(*codeconnections.CreateConnectionInput) (*request.Request, *codeconnections.CreateConnectionOutput)

	CreateHost(*codeconnections.CreateHostInput) (*codeconnections.CreateHostOutput, error)
	CreateHostWithContext(aws.Context, *codeconnections.CreateHostInput, ...request.Option) (*codeconnections.CreateHostOutput, error)
	CreateHostRequest(*codeconnections.CreateHostInput) (*request.Request, *codeconnections.CreateHostOutput)

	CreateRepositoryLink(*codeconnections.CreateRepositoryLinkInput) (*codeconnections.CreateRepositoryLinkOutput, error)
	CreateRepositoryLinkWithContext(aws.Context, *codeconnections.CreateRepositoryLinkInput, ...request.Option) (*codeconnections.CreateRepositoryLinkOutput, error)
	CreateRepositoryLinkRequest(*codeconnections.CreateRepositoryLinkInput) (*request.Request, *codeconnections.CreateRepositoryLinkOutput)

	CreateSyncConfiguration(*codeconnections.CreateSyncConfigurationInput) (*codeconnections.CreateSyncConfigurationOutput, error)
	CreateSyncConfigurationWithContext(aws.Context, *codeconnections.CreateSyncConfigurationInput, ...request.Option) (*codeconnections.CreateSyncConfigurationOutput, error)
	CreateSyncConfigurationRequest(*codeconnections.CreateSyncConfigurationInput) (*request.Request, *codeconnections.CreateSyncConfigurationOutput)

	DeleteConnection(*codeconnections.DeleteConnectionInput) (*codeconnections.DeleteConnectionOutput, error)
	DeleteConnectionWithContext(aws.Context, *codeconnections.DeleteConnectionInput, ...request.Option) (*codeconnections.DeleteConnectionOutput, error)
	DeleteConnectionRequest(*codeconnections.DeleteConnectionInput) (*request.Request, *codeconnections.DeleteConnectionOutput)

	DeleteHost(*codeconnections.DeleteHostInput) (*codeconnections.DeleteHostOutput, error)
	DeleteHostWithContext(aws.Context, *codeconnections.DeleteHostInput, ...request.Option) (*codeconnections.DeleteHostOutput, error)
	DeleteHostRequest(*codeconnections.DeleteHostInput) (*request.Request, *codeconnections.DeleteHostOutput)

	DeleteRepositoryLink(*codeconnections.DeleteRepositoryLinkInput) (*codeconnections.DeleteRepositoryLinkOutput, error)
	DeleteRepositoryLinkWithContext(aws.Context, *codeconnections.DeleteRepositoryLinkInput, ...request.Option) (*codeconnections.DeleteRepositoryLinkOutput, error)
	DeleteRepositoryLinkRequest(*codeconnections.DeleteRepositoryLinkInput) (*request.Request, *codeconnections.DeleteRepositoryLinkOutput)

	DeleteSyncConfiguration(*codeconnections.DeleteSyncConfigurationInput) (*codeconnections.DeleteSyncConfigurationOutput, error)
	DeleteSyncConfigurationWithContext(aws.Context, *codeconnections.DeleteSyncConfigurationInput, ...request.Option) (*codeconnections.DeleteSyncConfigurationOutput, error)
	DeleteSyncConfigurationRequest(*codeconnections.DeleteSyncConfigurationInput) (*request.Request, *codeconnections.DeleteSyncConfigurationOutput)

	GetConnection(*codeconnections.GetConnectionInput) (*codeconnections.GetConnectionOutput, error)
	GetConnectionWithContext(aws.Context, *codeconnections.GetConnectionInput, ...request.Option) (*codeconnections.GetConnectionOutput, error)
	GetConnectionRequest(*codeconnections.GetConnectionInput) (*request.Request, *codeconnections.GetConnectionOutput)

	GetHost(*codeconnections.GetHostInput) (*codeconnections.GetHostOutput, error)
	GetHostWithContext(aws.Context, *codeconnections.GetHostInput, ...request.Option) (*codeconnections.GetHostOutput, error)
	GetHostRequest(*codeconnections.GetHostInput) (*request.Request, *codeconnections.GetHostOutput)

	GetRepositoryLink(*codeconnections.GetRepositoryLinkInput) (*codeconnections.GetRepositoryLinkOutput, error)
	GetRepositoryLinkWithContext(aws.Context, *codeconnections.GetRepositoryLinkInput, ...request.Option) (*codeconnections.GetRepositoryLinkOutput, error)
	GetRepositoryLinkRequest(*codeconnections.GetRepositoryLinkInput) (*request.Request, *codeconnections.GetRepositoryLinkOutput)

	GetRepositorySyncStatus(*codeconnections.GetRepositorySyncStatusInput) (*codeconnections.GetRepositorySyncStatusOutput, error)
	GetRepositorySyncStatusWithContext(aws.Context, *codeconnections.GetRepositorySyncStatusInput, ...request.Option) (*codeconnections.GetRepositorySyncStatusOutput, error)
	GetRepositorySyncStatusRequest(*codeconnections.GetRepositorySyncStatusInput) (*request.Request, *codeconnections.GetRepositorySyncStatusOutput)

	GetResourceSyncStatus(*codeconnections.GetResourceSyncStatusInput) (*codeconnections.GetResourceSyncStatusOutput, error)
	GetResourceSyncStatusWithContext(aws.Context, *codeconnections.GetResourceSyncStatusInput, ...request.Option) (*codeconnections.GetResourceSyncStatusOutput, error)
	GetResourceSyncStatusRequest(*codeconnections.GetResourceSyncStatusInput) (*request.Request, *codeconnections.GetResourceSyncStatusOutput)

	GetSyncBlockerSummary(*codeconnections.GetSyncBlockerSummaryInput) (*codeconnections.GetSyncBlockerSummaryOutput, error)
	GetSyncBlockerSummaryWithContext(aws.Context, *codeconnections.GetSyncBlockerSummaryInput, ...request.Option) (*codeconnections.GetSyncBlockerSummaryOutput, error)
	GetSyncBlockerSummaryRequest(*codeconnections.GetSyncBlockerSummaryInput) (*request.Request, *codeconnections.GetSyncBlockerSummaryOutput)

	GetSyncConfiguration(*codeconnections.GetSyncConfigurationInput) (*codeconnections.GetSyncConfigurationOutput, error)
	GetSyncConfigurationWithContext(aws.Context, *codeconnections.GetSyncConfigurationInput, ...request.Option) (*codeconnections.GetSyncConfigurationOutput, error)
	GetSyncConfigurationRequest(*codeconnections.GetSyncConfigurationInput) (*request.Request, *codeconnections.GetSyncConfigurationOutput)

	ListConnections(*codeconnections.ListConnectionsInput) (*codeconnections.ListConnectionsOutput, error)
	ListConnectionsWithContext(aws.Context, *codeconnections.ListConnectionsInput, ...request.Option) (*codeconnections.ListConnectionsOutput, error)
	ListConnectionsRequest(*codeconnections.ListConnectionsInput) (*request.Request, *codeconnections.ListConnectionsOutput)

	ListConnectionsPages(*codeconnections.ListConnectionsInput, func(*codeconnections.ListConnectionsOutput, bool) bool) error
	ListConnectionsPagesWithContext(aws.Context, *codeconnections.ListConnectionsInput, func(*codeconnections.ListConnectionsOutput, bool) bool, ...request.Option) error

	ListHosts(*codeconnections.ListHostsInput) (*codeconnections.ListHostsOutput, error)
	ListHostsWithContext(aws.Context, *codeconnections.ListHostsInput, ...request.Option) (*codeconnections.ListHostsOutput, error)
	ListHostsRequest(*codeconnections.ListHostsInput) (*request.Request, *codeconnections.ListHostsOutput)

	ListHostsPages(*codeconnections.ListHostsInput, func(*codeconnections.ListHostsOutput, bool) bool) error
	ListHostsPagesWithContext(aws.Context, *codeconnections.ListHostsInput, func(*codeconnections.ListHostsOutput, bool) bool, ...request.Option) error

	ListRepositoryLinks(*codeconnections.ListRepositoryLinksInput) (*codeconnections.ListRepositoryLinksOutput, error)
	ListRepositoryLinksWithContext(aws.Context, *codeconnections.ListRepositoryLinksInput, ...request.Option) (*codeconnections.ListRepositoryLinksOutput, error)
	ListRepositoryLinksRequest(*codeconnections.ListRepositoryLinksInput) (*request.Request, *codeconnections.ListRepositoryLinksOutput)

	ListRepositoryLinksPages(*codeconnections.ListRepositoryLinksInput, func(*codeconnections.ListRepositoryLinksOutput, bool) bool) error
	ListRepositoryLinksPagesWithContext(aws.Context, *codeconnections.ListRepositoryLinksInput, func(*codeconnections.ListRepositoryLinksOutput, bool) bool, ...request.Option) error

	ListRepositorySyncDefinitions(*codeconnections.ListRepositorySyncDefinitionsInput) (*codeconnections.ListRepositorySyncDefinitionsOutput, error)
	ListRepositorySyncDefinitionsWithContext(aws.Context, *codeconnections.ListRepositorySyncDefinitionsInput, ...request.Option) (*codeconnections.ListRepositorySyncDefinitionsOutput, error)
	ListRepositorySyncDefinitionsRequest(*codeconnections.ListRepositorySyncDefinitionsInput) (*request.Request, *codeconnections.ListRepositorySyncDefinitionsOutput)

	ListSyncConfigurations(*codeconnections.ListSyncConfigurationsInput) (*codeconnections.ListSyncConfigurationsOutput, error)
	ListSyncConfigurationsWithContext(aws.Context, *codeconnections.ListSyncConfigurationsInput, ...request.Option) (*codeconnections.ListSyncConfigurationsOutput, error)
	ListSyncConfigurationsRequest(*codeconnections.ListSyncConfigurationsInput) (*request.Request, *codeconnections.ListSyncConfigurationsOutput)

	ListSyncConfigurationsPages(*codeconnections.ListSyncConfigurationsInput, func(*codeconnections.ListSyncConfigurationsOutput, bool) bool) error
	ListSyncConfigurationsPagesWithContext(aws.Context, *codeconnections.ListSyncConfigurationsInput, func(*codeconnections.ListSyncConfigurationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*codeconnections.ListTagsForResourceInput) (*codeconnections.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *codeconnections.ListTagsForResourceInput, ...request.Option) (*codeconnections.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*codeconnections.ListTagsForResourceInput) (*request.Request, *codeconnections.ListTagsForResourceOutput)

	TagResource(*codeconnections.TagResourceInput) (*codeconnections.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *codeconnections.TagResourceInput, ...request.Option) (*codeconnections.TagResourceOutput, error)
	TagResourceRequest(*codeconnections.TagResourceInput) (*request.Request, *codeconnections.TagResourceOutput)

	UntagResource(*codeconnections.UntagResourceInput) (*codeconnections.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *codeconnections.UntagResourceInput, ...request.Option) (*codeconnections.UntagResourceOutput, error)
	UntagResourceRequest(*codeconnections.UntagResourceInput) (*request.Request, *codeconnections.UntagResourceOutput)

	UpdateHost(*codeconnections.UpdateHostInput) (*codeconnections.UpdateHostOutput, error)
	UpdateHostWithContext(aws.Context, *codeconnections.UpdateHostInput, ...request.Option) (*codeconnections.UpdateHostOutput, error)
	UpdateHostRequest(*codeconnections.UpdateHostInput) (*request.Request, *codeconnections.UpdateHostOutput)

	UpdateRepositoryLink(*codeconnections.UpdateRepositoryLinkInput) (*codeconnections.UpdateRepositoryLinkOutput, error)
	UpdateRepositoryLinkWithContext(aws.Context, *codeconnections.UpdateRepositoryLinkInput, ...request.Option) (*codeconnections.UpdateRepositoryLinkOutput, error)
	UpdateRepositoryLinkRequest(*codeconnections.UpdateRepositoryLinkInput) (*request.Request, *codeconnections.UpdateRepositoryLinkOutput)

	UpdateSyncBlocker(*codeconnections.UpdateSyncBlockerInput) (*codeconnections.UpdateSyncBlockerOutput, error)
	UpdateSyncBlockerWithContext(aws.Context, *codeconnections.UpdateSyncBlockerInput, ...request.Option) (*codeconnections.UpdateSyncBlockerOutput, error)
	UpdateSyncBlockerRequest(*codeconnections.UpdateSyncBlockerInput) (*request.Request, *codeconnections.UpdateSyncBlockerOutput)

	UpdateSyncConfiguration(*codeconnections.UpdateSyncConfigurationInput) (*codeconnections.UpdateSyncConfigurationOutput, error)
	UpdateSyncConfigurationWithContext(aws.Context, *codeconnections.UpdateSyncConfigurationInput, ...request.Option) (*codeconnections.UpdateSyncConfigurationOutput, error)
	UpdateSyncConfigurationRequest(*codeconnections.UpdateSyncConfigurationInput) (*request.Request, *codeconnections.UpdateSyncConfigurationOutput)
}

var _ CodeConnectionsAPI = (*codeconnections.CodeConnections)(nil)