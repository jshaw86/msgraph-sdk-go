package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder provides operations to call the activate method.
type LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal instantiates a new MicrosoftGraphIdentityGovernanceActivateRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/microsoft.graph.identityGovernance.activate", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder instantiates a new MicrosoftGraphIdentityGovernanceActivateRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run a workflow object on-demand. You can run any workflow on-demand, including scheduled workflows. Workflows created from the 'Real-time employee termination' template are run on-demand only. When you run a workflow on demand, the tasks are executed regardless of whether the user state matches the scope and trigger execution conditions.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/identitygovernance-workflow-activate?view=graph-rest-1.0
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder) Post(ctx context.Context, body LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateActivatePostRequestBodyable, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation run a workflow object on-demand. You can run any workflow on-demand, including scheduled workflows. Workflows created from the 'Real-time employee termination' template are run on-demand only. When you run a workflow on demand, the tasks are executed regardless of whether the user state matches the scope and trigger execution conditions.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, body LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateActivatePostRequestBodyable, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
