package changescreensharingrole

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type ChangeScreenSharingRoleRequestBody struct {
    additionalData map[string]interface{};
    role *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ScreenSharingRole;
}
func NewChangeScreenSharingRoleRequestBody()(*ChangeScreenSharingRoleRequestBody) {
    m := &ChangeScreenSharingRoleRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChangeScreenSharingRoleRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChangeScreenSharingRoleRequestBody) GetRole()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ScreenSharingRole) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
func (m *ChangeScreenSharingRoleRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseScreenSharingRole)
        if err != nil {
            return err
        }
        cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ScreenSharingRole)
        m.SetRole(&cast)
        return nil
    }
    return res
}
func (m *ChangeScreenSharingRoleRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ChangeScreenSharingRoleRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetRole() != nil {
        cast := m.GetRole().String()
        err := writer.WriteStringValue("role", &cast)
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
func (m *ChangeScreenSharingRoleRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChangeScreenSharingRoleRequestBody) SetRole(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ScreenSharingRole)() {
    m.role = value
}
