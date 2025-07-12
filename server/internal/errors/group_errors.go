package errors

import "errors"

var (
	ErrGroupNotFound		  = errors.New("group not found")
	ErrUnableToGenerateGroupCode = errors.New("unable to generate a unique group code")
	ErrUserAlreadyInGroup = errors.New("user is already a member of the group")
	ErrUserNotInGroup	 = errors.New("user is not a member of the group")
	ErrGroupClosed		 = errors.New("group is closed")
	ErrUserNotAuthorized	 = errors.New("user is not authorized to perform this action")
)