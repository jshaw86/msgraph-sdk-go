package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarSharingMessageAction 
type CalendarSharingMessageAction struct {
    // The action property
    action *CalendarSharingAction
    // The actionType property
    actionType *CalendarSharingActionType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The importance property
    importance *CalendarSharingActionImportance
}
// NewCalendarSharingMessageAction instantiates a new calendarSharingMessageAction and sets the default values.
func NewCalendarSharingMessageAction()(*CalendarSharingMessageAction) {
    m := &CalendarSharingMessageAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCalendarSharingMessageActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalendarSharingMessageActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalendarSharingMessageAction(), nil
}
// GetAction gets the action property value. The action property
func (m *CalendarSharingMessageAction) GetAction()(*CalendarSharingAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetActionType gets the actionType property value. The actionType property
func (m *CalendarSharingMessageAction) GetActionType()(*CalendarSharingActionType) {
    if m == nil {
        return nil
    } else {
        return m.actionType
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CalendarSharingMessageAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalendarSharingMessageAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarSharingAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*CalendarSharingAction))
        }
        return nil
    }
    res["actionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarSharingActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val.(*CalendarSharingActionType))
        }
        return nil
    }
    res["importance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarSharingActionImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*CalendarSharingActionImportance))
        }
        return nil
    }
    return res
}
// GetImportance gets the importance property value. The importance property
func (m *CalendarSharingMessageAction) GetImportance()(*CalendarSharingActionImportance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// Serialize serializes information the current object
func (m *CalendarSharingMessageAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActionType() != nil {
        cast := (*m.GetActionType()).String()
        err := writer.WriteStringValue("actionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
        err := writer.WriteStringValue("importance", &cast)
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
// SetAction sets the action property value. The action property
func (m *CalendarSharingMessageAction) SetAction(value *CalendarSharingAction)() {
    if m != nil {
        m.action = value
    }
}
// SetActionType sets the actionType property value. The actionType property
func (m *CalendarSharingMessageAction) SetActionType(value *CalendarSharingActionType)() {
    if m != nil {
        m.actionType = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CalendarSharingMessageAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetImportance sets the importance property value. The importance property
func (m *CalendarSharingMessageAction) SetImportance(value *CalendarSharingActionImportance)() {
    if m != nil {
        m.importance = value
    }
}
