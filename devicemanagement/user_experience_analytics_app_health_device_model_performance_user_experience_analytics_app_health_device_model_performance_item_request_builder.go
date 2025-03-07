package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters user experience analytics appHealth Model Performance
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    m := &UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/{userExperienceAnalyticsAppHealthDeviceModelPerformance%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance for deviceManagement
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Model Performance
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance in deviceManagement
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance for deviceManagement
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation user experience analytics appHealth Model Performance
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance in deviceManagement
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
