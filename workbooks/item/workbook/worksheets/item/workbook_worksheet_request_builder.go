package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0ffc856a99ab076eacab92d71f01358a27ce7f82ca9cea1eeef71ccc0b690b04 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/protection"
    i100ec4503074fbce609d04c7af9f3038ecbe146a28dd1e9696d968fa800e43fb "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/usedrange"
    i63bfade034059e0bafdeccc55d1d84a9578056cd0b4452d244d1ea48d219bfd5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables"
    i9c65a5e0511223ed4b7aadfa985f3cecba8763d0234d695c0eba9470b08105fd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/pivottables"
    i9e565a8c9d35090b91ddaf8942fefd178ab4d8274ecd66a46cebd61471579ac8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/names"
    ia4b371526ab2f408171fc46a799dd775d49573d38da9d1f1459fa6b1e8cc5242 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/usedrangewithvaluesonly"
    iad940f8124c7334170f4bc84a78fc081f07f84f05e2bbcdad79b8c17146992f2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts"
    ib8d36af0717b2669dc8b8889287197f9036d4251ba8a717b88522dcf0bc194d8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/rangewithaddress"
    icbd8df2df6b23b4e1ba57be357e24d8f468fbc33c095797c5f66f4525b172267 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/range_escpaped"
    ie4b7c6c1c8b3e34e3e51127f9ebb673f0a04c788986d8229f1a710411d5f225d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/cellwithrowwithcolumn"
    i2b7eb38de9f32e0a84e0f4d15c3dfa5d1787f849e6127d28c3c84d7fb2c68f22 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/pivottables/item"
    i3ecee7d2013b82649254a2abcecff10c453e98848525efb8661b8198de377cbd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item"
    i8dc6bf9c3cd413c8da2a0e6e6c76573ab0a8670c078efc72a8efa32a1950ad10 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/names/item"
    iaf6e80887a038f65ff678934a224c20738b5518de1ef6bd042c68d64e3eb866a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item"
)

type WorkbookWorksheetRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorkbookWorksheetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *WorkbookWorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*ie4b7c6c1c8b3e34e3e51127f9ebb673f0a04c788986d8229f1a710411d5f225d.CellWithRowWithColumnRequestBuilder) {
    return ie4b7c6c1c8b3e34e3e51127f9ebb673f0a04c788986d8229f1a710411d5f225d.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookWorksheetRequestBuilder) Charts()(*iad940f8124c7334170f4bc84a78fc081f07f84f05e2bbcdad79b8c17146992f2.ChartsRequestBuilder) {
    return iad940f8124c7334170f4bc84a78fc081f07f84f05e2bbcdad79b8c17146992f2.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) ChartsById(id string)(*iaf6e80887a038f65ff678934a224c20738b5518de1ef6bd042c68d64e3eb866a.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return iaf6e80887a038f65ff678934a224c20738b5518de1ef6bd042c68d64e3eb866a.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorkbookWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetRequestBuilder) {
    m := &WorkbookWorksheetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}{?select,expand}";
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
func NewWorkbookWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookWorksheetRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
func (m *WorkbookWorksheetRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookWorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookWorksheetRequestBuilderGetQueryParameters)
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
func (m *WorkbookWorksheetRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *WorkbookWorksheetRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookWorksheetRequestBuilder) Get(q func (value *WorkbookWorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookWorksheet() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet), nil
}
func (m *WorkbookWorksheetRequestBuilder) Names()(*i9e565a8c9d35090b91ddaf8942fefd178ab4d8274ecd66a46cebd61471579ac8.NamesRequestBuilder) {
    return i9e565a8c9d35090b91ddaf8942fefd178ab4d8274ecd66a46cebd61471579ac8.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) NamesById(id string)(*i8dc6bf9c3cd413c8da2a0e6e6c76573ab0a8670c078efc72a8efa32a1950ad10.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return i8dc6bf9c3cd413c8da2a0e6e6c76573ab0a8670c078efc72a8efa32a1950ad10.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookWorksheetRequestBuilder) PivotTables()(*i9c65a5e0511223ed4b7aadfa985f3cecba8763d0234d695c0eba9470b08105fd.PivotTablesRequestBuilder) {
    return i9c65a5e0511223ed4b7aadfa985f3cecba8763d0234d695c0eba9470b08105fd.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) PivotTablesById(id string)(*i2b7eb38de9f32e0a84e0f4d15c3dfa5d1787f849e6127d28c3c84d7fb2c68f22.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return i2b7eb38de9f32e0a84e0f4d15c3dfa5d1787f849e6127d28c3c84d7fb2c68f22.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) Protection()(*i0ffc856a99ab076eacab92d71f01358a27ce7f82ca9cea1eeef71ccc0b690b04.ProtectionRequestBuilder) {
    return i0ffc856a99ab076eacab92d71f01358a27ce7f82ca9cea1eeef71ccc0b690b04.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) Range_escpaped()(*icbd8df2df6b23b4e1ba57be357e24d8f468fbc33c095797c5f66f4525b172267.RangeRequestBuilder) {
    return icbd8df2df6b23b4e1ba57be357e24d8f468fbc33c095797c5f66f4525b172267.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) RangeWithAddress(address *string)(*ib8d36af0717b2669dc8b8889287197f9036d4251ba8a717b88522dcf0bc194d8.RangeWithAddressRequestBuilder) {
    return ib8d36af0717b2669dc8b8889287197f9036d4251ba8a717b88522dcf0bc194d8.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorkbookWorksheetRequestBuilder) Tables()(*i63bfade034059e0bafdeccc55d1d84a9578056cd0b4452d244d1ea48d219bfd5.TablesRequestBuilder) {
    return i63bfade034059e0bafdeccc55d1d84a9578056cd0b4452d244d1ea48d219bfd5.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) TablesById(id string)(*i3ecee7d2013b82649254a2abcecff10c453e98848525efb8661b8198de377cbd.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return i3ecee7d2013b82649254a2abcecff10c453e98848525efb8661b8198de377cbd.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) UsedRange()(*i100ec4503074fbce609d04c7af9f3038ecbe146a28dd1e9696d968fa800e43fb.UsedRangeRequestBuilder) {
    return i100ec4503074fbce609d04c7af9f3038ecbe146a28dd1e9696d968fa800e43fb.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ia4b371526ab2f408171fc46a799dd775d49573d38da9d1f1459fa6b1e8cc5242.UsedRangeWithValuesOnlyRequestBuilder) {
    return ia4b371526ab2f408171fc46a799dd775d49573d38da9d1f1459fa6b1e8cc5242.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
