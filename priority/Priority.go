package priority

type SysLogLevel int

const (
	/* system is unusable */
	EMERG SysLogLevel = 0

	/* action must be taken immediately */
	ALERT SysLogLevel = 1

	/* critical conditions */
	CRIT SysLogLevel = 2

	/* error conditions */
	ERR SysLogLevel = 3

	/* warning conditions */
	WARNING SysLogLevel = 4

	/* normal but significant condition */
	NOTICE SysLogLevel = 5

	/* informational */
	INFO SysLogLevel = 6

	/* debug-level messages */
	DEBUG SysLogLevel = 7
)
