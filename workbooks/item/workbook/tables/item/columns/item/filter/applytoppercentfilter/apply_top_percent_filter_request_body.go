package applytoppercentfilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApplyTopPercentFilterRequestBody struct {
    additionalData map[string]interface{};
    percent *int32;
}
func NewApplyTopPercentFilterRequestBody()(*ApplyTopPercentFilterRequestBody) {
    m := &ApplyTopPercentFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApplyTopPercentFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApplyTopPercentFilterRequestBody) GetPercent()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.percent
    }
}
func (m *ApplyTopPercentFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["percent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPercent(val)
        return nil
    }
    return res
}
func (m *ApplyTopPercentFilterRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ApplyTopPercentFilterRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("percent", m.GetPercent())
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
func (m *ApplyTopPercentFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApplyTopPercentFilterRequestBody) SetPercent(value *int32)() {
    m.percent = value
}
