// Code generated by "stringer -type ActionType ."; DO NOT EDIT.

package send

import "strconv"

const _ActionType_name = "NopPostFavBoostReplyDel"

var _ActionType_index = [...]uint8{0, 3, 7, 10, 15, 20, 23}

func (i ActionType) String() string {
	if i < 0 || i >= ActionType(len(_ActionType_index)-1) {
		return "ActionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ActionType_name[_ActionType_index[i]:_ActionType_index[i+1]]
}
