package models
import (
    "errors"
)
// 
type TeamsAppResourceSpecificPermissionType int

const (
    DELEGATED_TEAMSAPPRESOURCESPECIFICPERMISSIONTYPE TeamsAppResourceSpecificPermissionType = iota
    APPLICATION_TEAMSAPPRESOURCESPECIFICPERMISSIONTYPE
    UNKNOWNFUTUREVALUE_TEAMSAPPRESOURCESPECIFICPERMISSIONTYPE
)

func (i TeamsAppResourceSpecificPermissionType) String() string {
    return []string{"delegated", "application", "unknownFutureValue"}[i]
}
func ParseTeamsAppResourceSpecificPermissionType(v string) (any, error) {
    result := DELEGATED_TEAMSAPPRESOURCESPECIFICPERMISSIONTYPE
    switch v {
        case "delegated":
            result = DELEGATED_TEAMSAPPRESOURCESPECIFICPERMISSIONTYPE
        case "application":
            result = APPLICATION_TEAMSAPPRESOURCESPECIFICPERMISSIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSAPPRESOURCESPECIFICPERMISSIONTYPE
        default:
            return 0, errors.New("Unknown TeamsAppResourceSpecificPermissionType value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAppResourceSpecificPermissionType(values []TeamsAppResourceSpecificPermissionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
