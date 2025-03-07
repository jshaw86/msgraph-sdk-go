package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters user experience analytics work from anywhere hardware readiness metrics.
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics work from anywhere hardware readiness metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable), nil
}
// Patch update the navigation property userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation user experience analytics work from anywhere hardware readiness metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
