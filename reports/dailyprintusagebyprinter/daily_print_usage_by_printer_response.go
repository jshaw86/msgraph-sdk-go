package dailyprintusagebyprinter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type DailyPrintUsageByPrinterResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintUsageByPrinter;
}
func NewDailyPrintUsageByPrinterResponse()(*DailyPrintUsageByPrinterResponse) {
    m := &DailyPrintUsageByPrinterResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DailyPrintUsageByPrinterResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DailyPrintUsageByPrinterResponse) GetNextLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nextLink
    }
}
func (m *DailyPrintUsageByPrinterResponse) GetValue()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *DailyPrintUsageByPrinterResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["@odata.nextLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNextLink(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintUsageByPrinter, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintUsageByPrinter))
        }
        m.SetValue(res)
        return nil
    }
    return res
}
func (m *DailyPrintUsageByPrinterResponse) IsNil()(bool) {
    return m == nil
}
func (m *DailyPrintUsageByPrinterResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetNextLink())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DailyPrintUsageByPrinterResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DailyPrintUsageByPrinterResponse) SetNextLink(value *string)() {
    m.nextLink = value
}
func (m *DailyPrintUsageByPrinterResponse) SetValue(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintUsageByPrinter)() {
    m.value = value
}
