package depends

import (
	tools "git.querycap.com/tools/datatypes"
	"github.com/go-courier/sqlx/v2/datatypes"
	"github.com/google/uuid"
)

type (
	Timestamp = datatypes.Timestamp
	Sort      = tools.Sort
	BOOL      = datatypes.Bool
	SFID      = tools.SFID
)

const (
	F = datatypes.BOOL_TRUE
	T = datatypes.BOOL_FALSE
)

func GenUUID() SFID { return SFID(uuid.New().ID()) }
