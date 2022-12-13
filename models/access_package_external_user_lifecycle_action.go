package models
import (
    "errors"
)
// Provides operations to manage the collection of applicationTemplate entities.
type AccessPackageExternalUserLifecycleAction int

const (
    NONE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION AccessPackageExternalUserLifecycleAction = iota
    BLOCKSIGNIN_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
    BLOCKSIGNINANDDELETE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
    UNKNOWNFUTUREVALUE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
)

func (i AccessPackageExternalUserLifecycleAction) String() string {
    return []string{"none", "blockSignIn", "blockSignInAndDelete", "unknownFutureValue"}[i]
}
func ParseAccessPackageExternalUserLifecycleAction(v string) (interface{}, error) {
    result := NONE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
    switch v {
        case "none":
            result = NONE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
        case "blockSignIn":
            result = BLOCKSIGNIN_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
        case "blockSignInAndDelete":
            result = BLOCKSIGNINANDDELETE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGEEXTERNALUSERLIFECYCLEACTION
        default:
            return 0, errors.New("Unknown AccessPackageExternalUserLifecycleAction value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageExternalUserLifecycleAction(values []AccessPackageExternalUserLifecycleAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
