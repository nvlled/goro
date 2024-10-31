package standard

import (
	"encoding/hex"

	"github.com/MagicalTux/goro/core"
	"github.com/MagicalTux/goro/core/phpv"
)

// > func string bin2hex ( string $str )
func fncBin2hex(ctx phpv.Context, args []*phpv.ZVal) (*phpv.ZVal, error) {
	var s string

	_, err := core.Expand(ctx, args, &s)
	if err != nil {
		return nil, err
	}

	return phpv.ZString(hex.EncodeToString([]byte(s))).ZVal(), nil
}
