package LoggerDomain

import (
	"github.com/yeencloud/lib-shared/log"
)

var LogFieldTraceScope = log.Path{Identifier: "trace"}
var LogFieldTraceDump = log.Path{Parent: &LogFieldTraceScope, Identifier: "dump"}
