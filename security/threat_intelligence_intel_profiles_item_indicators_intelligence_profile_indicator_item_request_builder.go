package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder provides operations to manage the indicators property of the microsoft.graph.security.intelligenceProfile entity.
type ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderGetQueryParameters includes an assemblage of high-fidelity network indicators of compromise.
type ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderGetQueryParameters
}
// NewThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderInternal instantiates a new IntelligenceProfileIndicatorItemRequestBuilder and sets the default values.
func NewThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder) {
    m := &ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/intelProfiles/{intelligenceProfile%2Did}/indicators/{intelligenceProfileIndicator%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder instantiates a new IntelligenceProfileIndicatorItemRequestBuilder and sets the default values.
func NewThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get includes an assemblage of high-fidelity network indicators of compromise.
func (m *ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.IntelligenceProfileIndicatorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateIntelligenceProfileIndicatorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.IntelligenceProfileIndicatorable), nil
}
// ToGetRequestInformation includes an assemblage of high-fidelity network indicators of compromise.
func (m *ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceIntelProfilesItemIndicatorsIntelligenceProfileIndicatorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
