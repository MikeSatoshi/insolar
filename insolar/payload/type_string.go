// Code generated by "stringer -type=Type"; DO NOT EDIT.

package payload

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TypeUnknown-0]
	_ = x[TypeMeta-1]
	_ = x[TypeError-2]
	_ = x[TypeID-3]
	_ = x[TypeState-4]
	_ = x[TypeGetObject-5]
	_ = x[TypePassState-6]
	_ = x[TypeObjIndex-7]
	_ = x[TypeObjState-8]
	_ = x[TypeIndex-9]
	_ = x[TypePass-10]
	_ = x[TypeGetCode-11]
	_ = x[TypeCode-12]
	_ = x[TypeSetCode-13]
	_ = x[TypeSetIncomingRequest-14]
	_ = x[TypeSetOutgoingRequest-15]
	_ = x[TypeGetFilament-16]
	_ = x[TypeFilamentSegment-17]
	_ = x[TypeSetResult-18]
	_ = x[TypeActivate-19]
	_ = x[TypeRequestInfo-20]
	_ = x[TypeGotHotConfirmation-21]
	_ = x[TypeDeactivate-22]
	_ = x[TypeUpdate-23]
	_ = x[_latestType-24]
}

const _Type_name = "TypeUnknownTypeMetaTypeErrorTypeIDTypeStateTypeGetObjectTypePassStateTypeObjIndexTypeObjStateTypeIndexTypePassTypeGetCodeTypeCodeTypeSetCodeTypeSetIncomingRequestTypeSetOutgoingRequestTypeGetFilamentTypeFilamentSegmentTypeSetResultTypeActivateTypeRequestInfoTypeGotHotConfirmationTypeDeactivateTypeUpdate_latestType"

var _Type_index = [...]uint16{0, 11, 19, 28, 34, 43, 56, 69, 81, 93, 102, 110, 121, 129, 140, 162, 184, 199, 218, 231, 243, 258, 280, 294, 304, 315}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
