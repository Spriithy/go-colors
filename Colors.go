package colors

var prefix = "\033["

const (
	BLACK = prefix + "0;30m"
	RED = prefix + "0;31m"
	GREEN = prefix + "0;32m"
	ORANGE = prefix + "0;33m"
	BLUE = prefix + "0;34m"
	PURPLE = prefix + "0;35m"
	CYAN = prefix + "0;36m"
	LIGHT_GRAY = prefix + "0;37m"
	DARK_GRAY = prefix + "1;30m"
	LIGHT_RED = prefix + "1;31m"
	LIGHT_GREEN = prefix + "1;32m"
	YELLOW = prefix + "1;33m"
	LIGHT_BLUE = prefix + "1;34m"
	LIGHT_PUPRLE = prefix + "1;35m"
	LIGHT_CYAN = prefix + "1;36m"
	WHITE = prefix + "1;37m"

	NC = prefix + "0m"
	NO_COLOR = NC
)
