package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.customTaskExtension entity.
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetQueryParameters the unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderInternal instantiates a new CreatedByRequestBuilder and sets the default values.
func NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) {
    m := &LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/customTaskExtensions/{customTaskExtension%2Did}/createdBy{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder instantiates a new CreatedByRequestBuilder and sets the default values.
func NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
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
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) MailboxSettings()(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByMailboxSettingsRequestBuilder) {
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
