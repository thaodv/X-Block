package message

import (
	"XBlock/common"
	"XBlock/common/log"
	. "XBlock/net/protocol"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
)

type notFound struct {
	msgHdr
	hash common.Uint256
}

