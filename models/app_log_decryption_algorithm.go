package models
import (
    "errors"
)
// 
type AppLogDecryptionAlgorithm int

const (
    // decrypting using Aes256.
    AES256_APPLOGDECRYPTIONALGORITHM AppLogDecryptionAlgorithm = iota
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_APPLOGDECRYPTIONALGORITHM
)

func (i AppLogDecryptionAlgorithm) String() string {
    return []string{"aes256", "unknownFutureValue"}[i]
}
func ParseAppLogDecryptionAlgorithm(v string) (any, error) {
    result := AES256_APPLOGDECRYPTIONALGORITHM
    switch v {
        case "aes256":
            result = AES256_APPLOGDECRYPTIONALGORITHM
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPLOGDECRYPTIONALGORITHM
        default:
            return 0, errors.New("Unknown AppLogDecryptionAlgorithm value: " + v)
    }
    return &result, nil
}
func SerializeAppLogDecryptionAlgorithm(values []AppLogDecryptionAlgorithm) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
