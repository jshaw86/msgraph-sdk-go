package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder provides operations to call the upgrade method.
type ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilderInternal instantiates a new UpgradeRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder) {
    m := &ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/installedApps/{teamsAppInstallation%2Did}/upgrade", pathParameters),
    }
    return m
}
// NewItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder instantiates a new UpgradeRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upgrade an app installation within a chat.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chat-teamsappinstallation-upgrade?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder) Post(ctx context.Context, body ItemJoinedTeamsItemInstalledAppsItemUpgradePostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation upgrade an app installation within a chat.
func (m *ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedTeamsItemInstalledAppsItemUpgradePostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemInstalledAppsItemUpgradeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
