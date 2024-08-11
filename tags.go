package lgo

type Tag string

const(
	DEBUG_TAG Tag	= "\033[1;32m[DEBUG]\033[0m: " // Green
	ERROR_TAG Tag	= "\033[1;31m[ERROR]\033[0m: " // Red
	INFO_TAG  Tag	= "\033[1;36m [INFO]\033[0m: " // Teal
)
