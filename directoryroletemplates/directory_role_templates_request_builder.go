package directoryroletemplates

import (
    i13f35543e7fd513b7896b53959450496b4b40a0d8ed709962f6871dc35b86ec4 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/getbyids"
    i14aa4ead11e74baa62edc9a283a7340417496098b3a9b5697fdc955d20086255 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/validateproperties"
    i1deef55c0cacc5fe1d33525ca2ca081a8626edb9cd782f2be58da2bdabec7837 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/getavailableextensionproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type DirectoryRoleTemplatesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DirectoryRoleTemplatesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewDirectoryRoleTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplatesRequestBuilder) {
    m := &DirectoryRoleTemplatesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/directoryRoleTemplates{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewDirectoryRoleTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DirectoryRoleTemplatesRequestBuilder) CreateGetRequestInformation(q func (value *DirectoryRoleTemplatesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DirectoryRoleTemplatesRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DirectoryRoleTemplatesRequestBuilder) CreatePostRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DirectoryRoleTemplatesRequestBuilder) Get(q func (value *DirectoryRoleTemplatesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*DirectoryRoleTemplatesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryRoleTemplatesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*DirectoryRoleTemplatesResponse), nil
}
func (m *DirectoryRoleTemplatesRequestBuilder) GetAvailableExtensionProperties()(*i1deef55c0cacc5fe1d33525ca2ca081a8626edb9cd782f2be58da2bdabec7837.GetAvailableExtensionPropertiesRequestBuilder) {
    return i1deef55c0cacc5fe1d33525ca2ca081a8626edb9cd782f2be58da2bdabec7837.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplatesRequestBuilder) GetByIds()(*i13f35543e7fd513b7896b53959450496b4b40a0d8ed709962f6871dc35b86ec4.GetByIdsRequestBuilder) {
    return i13f35543e7fd513b7896b53959450496b4b40a0d8ed709962f6871dc35b86ec4.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplatesRequestBuilder) Post(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleTemplate, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDirectoryRoleTemplate() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRoleTemplate), nil
}
func (m *DirectoryRoleTemplatesRequestBuilder) ValidateProperties()(*i14aa4ead11e74baa62edc9a283a7340417496098b3a9b5697fdc955d20086255.ValidatePropertiesRequestBuilder) {
    return i14aa4ead11e74baa62edc9a283a7340417496098b3a9b5697fdc955d20086255.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
