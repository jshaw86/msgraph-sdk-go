package windowsinformationprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ib765ad9fa24ed6bbb10d9fc8ca08fbb6f4a83060dd75b5daca83142349e3bfd0 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/windowsinformationprotection/assign"
)

type WindowsInformationProtectionRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WindowsInformationProtectionRequestBuilder) Assign()(*ib765ad9fa24ed6bbb10d9fc8ca08fbb6f4a83060dd75b5daca83142349e3bfd0.AssignRequestBuilder) {
    return ib765ad9fa24ed6bbb10d9fc8ca08fbb6f4a83060dd75b5daca83142349e3bfd0.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWindowsInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsInformationProtectionRequestBuilder) {
    m := &WindowsInformationProtectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.windowsInformationProtection";
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
func NewWindowsInformationProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
