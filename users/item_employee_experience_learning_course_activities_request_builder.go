package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperienceUser entity.
type ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEmployeeExperienceLearningCourseActivitiesRequestBuilderGetQueryParameters get a list of the learningCourseActivity objects (assigned or self-initiated) for a user.
type ItemEmployeeExperienceLearningCourseActivitiesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemEmployeeExperienceLearningCourseActivitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeExperienceLearningCourseActivitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEmployeeExperienceLearningCourseActivitiesRequestBuilderGetQueryParameters
}
// ByLearningCourseActivityId provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperienceUser entity.
func (m *ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder) ByLearningCourseActivityId(learningCourseActivityId string)(*ItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if learningCourseActivityId != "" {
        urlTplParams["learningCourseActivity%2Did"] = learningCourseActivityId
    }
    return NewItemEmployeeExperienceLearningCourseActivitiesLearningCourseActivityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEmployeeExperienceLearningCourseActivitiesRequestBuilderInternal instantiates a new LearningCourseActivitiesRequestBuilder and sets the default values.
func NewItemEmployeeExperienceLearningCourseActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder) {
    m := &ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/employeeExperience/learningCourseActivities{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemEmployeeExperienceLearningCourseActivitiesRequestBuilder instantiates a new LearningCourseActivitiesRequestBuilder and sets the default values.
func NewItemEmployeeExperienceLearningCourseActivitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEmployeeExperienceLearningCourseActivitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder) Count()(*ItemEmployeeExperienceLearningCourseActivitiesCountRequestBuilder) {
    return NewItemEmployeeExperienceLearningCourseActivitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the learningCourseActivity objects (assigned or self-initiated) for a user.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/learningcourseactivity-list?view=graph-rest-1.0
func (m *ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEmployeeExperienceLearningCourseActivitiesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningCourseActivityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityCollectionResponseable), nil
}
// ToGetRequestInformation get a list of the learningCourseActivity objects (assigned or self-initiated) for a user.
func (m *ItemEmployeeExperienceLearningCourseActivitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEmployeeExperienceLearningCourseActivitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
