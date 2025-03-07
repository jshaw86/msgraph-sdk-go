package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters read properties and relationships of the appLogCollectionRequest object.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal instantiates a new AppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) {
    m := &MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}/appLogCollectionRequests/{appLogCollectionRequest%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder instantiates a new AppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDownloadUrl provides operations to call the createDownloadUrl method.
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) CreateDownloadUrl()(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) {
    return NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete deletes a appLogCollectionRequest.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-devices-applogcollectionrequest-delete?view=graph-rest-1.0
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the appLogCollectionRequest object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-devices-applogcollectionrequest-get?view=graph-rest-1.0
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable), nil
}
// Patch update the properties of a appLogCollectionRequest object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-devices-applogcollectionrequest-update?view=graph-rest-1.0
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable), nil
}
// ToDeleteRequestInformation deletes a appLogCollectionRequest.
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation read properties and relationships of the appLogCollectionRequest object.
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a appLogCollectionRequest object.
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppLogCollectionRequestable, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
