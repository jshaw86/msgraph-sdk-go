package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsHelloForBusinessAuthenticationMethod struct {
    AuthenticationMethod
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    device *Device;
    displayName *string;
    keyStrength *AuthenticationMethodKeyStrength;
}
func NewWindowsHelloForBusinessAuthenticationMethod()(*WindowsHelloForBusinessAuthenticationMethod) {
    m := &WindowsHelloForBusinessAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
func (m *WindowsHelloForBusinessAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *WindowsHelloForBusinessAuthenticationMethod) GetDevice()(*Device) {
    if m == nil {
        return nil
    } else {
        return m.device
    }
}
func (m *WindowsHelloForBusinessAuthenticationMethod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *WindowsHelloForBusinessAuthenticationMethod) GetKeyStrength()(*AuthenticationMethodKeyStrength) {
    if m == nil {
        return nil
    } else {
        return m.keyStrength
    }
}
func (m *WindowsHelloForBusinessAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["device"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDevice() })
        if err != nil {
            return err
        }
        m.SetDevice(val.(*Device))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["keyStrength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodKeyStrength)
        if err != nil {
            return err
        }
        cast := val.(AuthenticationMethodKeyStrength)
        m.SetKeyStrength(&cast)
        return nil
    }
    return res
}
func (m *WindowsHelloForBusinessAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
func (m *WindowsHelloForBusinessAuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetKeyStrength() != nil {
        cast := m.GetKeyStrength().String()
        err = writer.WriteStringValue("keyStrength", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsHelloForBusinessAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *WindowsHelloForBusinessAuthenticationMethod) SetDevice(value *Device)() {
    m.device = value
}
func (m *WindowsHelloForBusinessAuthenticationMethod) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *WindowsHelloForBusinessAuthenticationMethod) SetKeyStrength(value *AuthenticationMethodKeyStrength)() {
    m.keyStrength = value
}
