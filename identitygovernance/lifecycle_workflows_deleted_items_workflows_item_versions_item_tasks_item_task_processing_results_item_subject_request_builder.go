package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder provides operations to manage the subject property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderGetQueryParameters the unique identifier of the Azure AD user targeted for the task execution.Supports $filter(eq, ne) and $expand.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderInternal instantiates a new SubjectRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}/tasks/{task%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/subject{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder instantiates a new SubjectRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the unique identifier of the Azure AD user targeted for the task execution.Supports $filter(eq, ne) and $expand.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// MailboxSettings the mailboxSettings property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder) MailboxSettings()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectMailboxSettingsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the unique identifier of the Azure AD user targeted for the task execution.Supports $filter(eq, ne) and $expand.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
