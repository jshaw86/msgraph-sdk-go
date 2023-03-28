package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\members\{directoryObject-id}
type ItemMembersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemMembersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    m := &ItemMembersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/members/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemMembersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphApplication casts the previous resource to application.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphApplication()(*ItemMembersItemGraphApplicationRequestBuilder) {
    return NewItemMembersItemGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphDevice()(*ItemMembersItemGraphDeviceRequestBuilder) {
    return NewItemMembersItemGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphGroup()(*ItemMembersItemGraphGroupRequestBuilder) {
    return NewItemMembersItemGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphOrgContact()(*ItemMembersItemGraphOrgContactRequestBuilder) {
    return NewItemMembersItemGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemMembersItemGraphServicePrincipalRequestBuilder) {
    return NewItemMembersItemGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemMembersItemGraphUserRequestBuilder) {
    return NewItemMembersItemGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of group entities.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) Ref()(*ItemMembersItemRefRequestBuilder) {
    return NewItemMembersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
