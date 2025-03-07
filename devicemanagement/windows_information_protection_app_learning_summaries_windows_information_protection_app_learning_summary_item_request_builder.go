package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
type WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetQueryParameters read properties and relationships of the windowsInformationProtectionAppLearningSummary object.
type WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetQueryParameters
}
// WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal instantiates a new WindowsInformationProtectionAppLearningSummaryItemRequestBuilder and sets the default values.
func NewWindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) {
    m := &WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsInformationProtectionAppLearningSummaries/{windowsInformationProtectionAppLearningSummary%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewWindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder instantiates a new WindowsInformationProtectionAppLearningSummaryItemRequestBuilder and sets the default values.
func NewWindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a windowsInformationProtectionAppLearningSummary.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-wip-windowsinformationprotectionapplearningsummary-delete?view=graph-rest-1.0
func (m *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get read properties and relationships of the windowsInformationProtectionAppLearningSummary object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-wip-windowsinformationprotectionapplearningsummary-get?view=graph-rest-1.0
func (m *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable), nil
}
// Patch update the properties of a windowsInformationProtectionAppLearningSummary object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-wip-windowsinformationprotectionapplearningsummary-update?view=graph-rest-1.0
func (m *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable, requestConfiguration *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable), nil
}
// ToDeleteRequestInformation deletes a windowsInformationProtectionAppLearningSummary.
func (m *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the windowsInformationProtectionAppLearningSummary object.
func (m *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a windowsInformationProtectionAppLearningSummary object.
func (m *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLearningSummaryable, requestConfiguration *WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
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
