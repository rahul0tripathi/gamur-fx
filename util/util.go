package util

import (
	"github.com/rahul0tripathi/gamur/util/logger"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	logger.Module,
)
