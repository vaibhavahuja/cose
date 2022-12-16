// (c) 2022-2022, LDC Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package aesgcm

import (
	"github.com/ldclabs/cose/go/key"
)

func init() {
	key.RegisterEncryptor(key.KtySymmetric, key.AlgA128GCM, NewAESGCM)
	key.RegisterEncryptor(key.KtySymmetric, key.AlgA192GCM, NewAESGCM)
	key.RegisterEncryptor(key.KtySymmetric, key.AlgA256GCM, NewAESGCM)
}
