package msg

import (
	"errors"
)

/*
   Creation Time: 2018 - Apr - 07
   Created by:  Ehsan N. Moosa (ehsan)
   Maintainers:
       1.  Ehsan N. Moosa (ehsan)
   Auditor: Ehsan N. Moosa
   Copyright Ronak Software Group 2018
*/

var (
	ErrNoHandler     = errors.New("no handler submitted")
	ErrTimeout       = errors.New("time out")
	ErrIncorrectSize = errors.New("incorrect size")
)

// Error Codes
const (
	ErrCodeInternal         = "E00"
	ErrCodeInvalid          = "E01"
	ErrCodeUnavailable      = "E02"
	ErrCodeTooMany          = "E03"
	ErrCodeTooFew           = "E04"
	ErrCodeIncomplete       = "E05"
	ErrCodeTimeout          = "E06"
	ErrCodeAccess           = "E07"
	ErrCodeAlreadyExists    = "E08"
	ErrCodeBusy             = "E09"
	ErrCodeOutOfRange       = "E10"
	ErrCodePartiallyApplied = "E11"
)

// Error Items
const (
	ErrItemPhone            = "PHONE"
	ErrItemPhoneCode        = "PHONE_CODE"
	ErrItemUserID           = "USER_ID"
	ErrItemPeer             = "PEER"
	ErrItemPeerType         = "PEER_TYPE"
	ErrItemInput            = "INPUT"
	ErrItemRequest          = "REQUEST"
	ErrItemMessage          = "MESSAGE"
	ErrItemMessageID        = "MESSAGE_ID"
	ErrItemServer           = "SERVER"
	ErrItemPq               = "PQ"
	ErrItemEncryption       = "ENCRYPTION"
	ErrItemRsaKey           = "RSA_KEY"
	ErrItemProto            = "PROTO"
	ErrItemDhKey            = "DH_KEY"
	ErrItemSignIn           = "SIGN_IN"
	ErrItemRandomID         = "RANDOM_ID"
	ErrItemAccessHash       = "ACCESS_HASH"
	ErrItemJobWorker        = "JOB_WORKER"
	ErrItemAuth             = "AUTH"
	ErrItemAuthID           = "AUTH_ID"
	ErrItemUsername         = "USERNAME"
	ErrItemChatText         = "CHAT_TEXT"
	ErrItemGroupTitle       = "GROUP_TITLE"
	ErrItemGroupID          = "GROUP_ID"
	ErrItemGroup            = "GROUP"
	ErrItemGroupMember      = "GROUP_MEMBER"
	ErrItemUsers            = "USERS"
	ErrItemRetractTime      = "RETRACT_TIME"
	ErrItemBio              = "BIO"
	ErrItemApi              = "API"
	ErrItemLastAdmin        = "LAST_ADMIN"
	ErrItemDeleteCreator    = "DELETE_CREATOR"
	ErrItemFilePartID       = "FILE_PART_ID"
	ErrItemFileParts        = "FILE_PARTS"
	ErrItemFilePartSize     = "FILE_PART_SIZE"
	ErrItemDeviceToken      = "DEVICE_TOKEN"
	ErrItemDeviceModel      = "DEVICE_MODEL"
	ErrItemDocument         = "DOCUMENT"
	ErrItemToken            = "TOKEN"
	ErrItemMedia            = "MEDIA"
	ErrItemPinnedDialogs    = "PINNED_DIALOGS"
	ErrItemSalt             = "SALT"
	ErrItemCounter          = "COUNTER"
	ErrItemAppUpdate        = "APP_UPDATE"
	ErrItemBindUser         = "BIND_USER"
	ErrItemActiveConnection = "ACTIVE_CONN"
	ErrItemInputFile        = "INPUT_FILE"
	ErrItemPhotoID          = "PHOTO_ID"
	ErrItemBot              = "BOT"
	ErrItemPhoto            = "PHOTO"
	ErrItemMember           = "MEMBER"
	ErrItemFeature          = "FEATURE"
)
