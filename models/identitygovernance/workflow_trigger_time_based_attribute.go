package identitygovernance
import (
    "errors"
)
// 
type WorkflowTriggerTimeBasedAttribute int

const (
    EMPLOYEEHIREDATE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE WorkflowTriggerTimeBasedAttribute = iota
    EMPLOYEELEAVEDATETIME_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
    UNKNOWNFUTUREVALUE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
    CREATEDDATETIME_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
)

func (i WorkflowTriggerTimeBasedAttribute) String() string {
    return []string{"employeeHireDate", "employeeLeaveDateTime", "unknownFutureValue", "createdDateTime"}[i]
}
func ParseWorkflowTriggerTimeBasedAttribute(v string) (any, error) {
    result := EMPLOYEEHIREDATE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
    switch v {
        case "employeeHireDate":
            result = EMPLOYEEHIREDATE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
        case "employeeLeaveDateTime":
            result = EMPLOYEELEAVEDATETIME_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
        case "createdDateTime":
            result = CREATEDDATETIME_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
        default:
            return 0, errors.New("Unknown WorkflowTriggerTimeBasedAttribute value: " + v)
    }
    return &result, nil
}
func SerializeWorkflowTriggerTimeBasedAttribute(values []WorkflowTriggerTimeBasedAttribute) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
