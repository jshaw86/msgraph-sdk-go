package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IdentityUserFlowAttributeAssignment struct {
    Entity
    displayName *string;
    isOptional *bool;
    requiresVerification *bool;
    userAttribute *IdentityUserFlowAttribute;
    userAttributeValues []UserAttributeValuesItem;
    userInputType *IdentityUserFlowAttributeInputType;
}
func NewIdentityUserFlowAttributeAssignment()(*IdentityUserFlowAttributeAssignment) {
    m := &IdentityUserFlowAttributeAssignment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *IdentityUserFlowAttributeAssignment) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *IdentityUserFlowAttributeAssignment) GetIsOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOptional
    }
}
func (m *IdentityUserFlowAttributeAssignment) GetRequiresVerification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requiresVerification
    }
}
func (m *IdentityUserFlowAttributeAssignment) GetUserAttribute()(*IdentityUserFlowAttribute) {
    if m == nil {
        return nil
    } else {
        return m.userAttribute
    }
}
func (m *IdentityUserFlowAttributeAssignment) GetUserAttributeValues()([]UserAttributeValuesItem) {
    if m == nil {
        return nil
    } else {
        return m.userAttributeValues
    }
}
func (m *IdentityUserFlowAttributeAssignment) GetUserInputType()(*IdentityUserFlowAttributeInputType) {
    if m == nil {
        return nil
    } else {
        return m.userInputType
    }
}
func (m *IdentityUserFlowAttributeAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOptional(val)
        return nil
    }
    res["requiresVerification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequiresVerification(val)
        return nil
    }
    res["userAttribute"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityUserFlowAttribute() })
        if err != nil {
            return err
        }
        m.SetUserAttribute(val.(*IdentityUserFlowAttribute))
        return nil
    }
    res["userAttributeValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserAttributeValuesItem() })
        if err != nil {
            return err
        }
        res := make([]UserAttributeValuesItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserAttributeValuesItem))
        }
        m.SetUserAttributeValues(res)
        return nil
    }
    res["userInputType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityUserFlowAttributeInputType)
        if err != nil {
            return err
        }
        cast := val.(IdentityUserFlowAttributeInputType)
        m.SetUserInputType(&cast)
        return nil
    }
    return res
}
func (m *IdentityUserFlowAttributeAssignment) IsNil()(bool) {
    return m == nil
}
func (m *IdentityUserFlowAttributeAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOptional", m.GetIsOptional())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requiresVerification", m.GetRequiresVerification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userAttribute", m.GetUserAttribute())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserAttributeValues()))
        for i, v := range m.GetUserAttributeValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userAttributeValues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserInputType() != nil {
        cast := m.GetUserInputType().String()
        err = writer.WriteStringValue("userInputType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *IdentityUserFlowAttributeAssignment) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *IdentityUserFlowAttributeAssignment) SetIsOptional(value *bool)() {
    m.isOptional = value
}
func (m *IdentityUserFlowAttributeAssignment) SetRequiresVerification(value *bool)() {
    m.requiresVerification = value
}
func (m *IdentityUserFlowAttributeAssignment) SetUserAttribute(value *IdentityUserFlowAttribute)() {
    m.userAttribute = value
}
func (m *IdentityUserFlowAttributeAssignment) SetUserAttributeValues(value []UserAttributeValuesItem)() {
    m.userAttributeValues = value
}
func (m *IdentityUserFlowAttributeAssignment) SetUserInputType(value *IdentityUserFlowAttributeInputType)() {
    m.userInputType = value
}
