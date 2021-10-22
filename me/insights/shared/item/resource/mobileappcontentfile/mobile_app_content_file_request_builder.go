package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i16fa1d9e42a0ca776693e29a59e9e4b9e2a2d52e05b446312fbeaba51b8f5efa "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/mobileappcontentfile/commit"
    iffe53ce30a6bcb0c29968ee1424419feda780d84bf038a72cbceaf30915477ee "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/mobileappcontentfile/renewupload"
)

type MobileAppContentFileRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*i16fa1d9e42a0ca776693e29a59e9e4b9e2a2d52e05b446312fbeaba51b8f5efa.CommitRequestBuilder) {
    return i16fa1d9e42a0ca776693e29a59e9e4b9e2a2d52e05b446312fbeaba51b8f5efa.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.mobileAppContentFile";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewMobileAppContentFileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppContentFileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*iffe53ce30a6bcb0c29968ee1424419feda780d84bf038a72cbceaf30915477ee.RenewUploadRequestBuilder) {
    return iffe53ce30a6bcb0c29968ee1424419feda780d84bf038a72cbceaf30915477ee.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
