package downloadUtils

import "errors"

// pixiv error vars
var (
	ErrPixivBodyNil     = errors.New("pixiv body is nil")
	ErrPixivPageInvalid = errors.New("invalid page specified, page number should start from 0")
)
