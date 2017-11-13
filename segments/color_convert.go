package segments

import(
)

/*
var convertColors = map[color.Attribute]color.Attribute {
    color.FgBlack   : color.BgBlack,
    color.FgRed     : color.BgRed,
    color.FgGreen   : color.BgGreen,
    color.FgYellow  : color.BgYellow,
    color.FgBlue    : color.BgBlue,
    color.FgMagenta : color.BgMagenta,
    color.FgCyan    : color.BgCyan,
    color.FgWhite   : color.BgWhite,

    color.BgBlack   : color.FgBlack,
    color.BgRed     : color.FgRed,
    color.BgGreen   : color.FgGreen,
    color.BgYellow  : color.FgYellow,
    color.BgBlue    : color.FgBlue,
    color.BgMagenta : color.FgMagenta,
    color.BgCyan    : color.FgCyan,
    color.BgWhite   : color.FgWhite,
}
*/

/*
# Regular Colors
\[\033[0;30m\] # Black
\[\033[0;31m\] # Red
\[\033[0;32m\] # Green
\[\033[0;33m\] # Yellow
\[\033[0;34m\] # Blue
\[\033[0;35m\] # Purple
\[\033[0;36m\] # Cyan
\[\033[0;37m\] # White

# High Intensty
\[\033[0;90m\] # Black
\[\033[0;91m\] # Red
\[\033[0;92m\] # Green
\[\033[0;93m\] # Yellow
\[\033[0;94m\] # Blue
\[\033[0;95m\] # Purple
\[\033[0;96m\] # Cyan
\[\033[0;97m\] # White

# Background
\[\033[40m\] # Black
\[\033[41m\] # Red
\[\033[42m\] # Green
\[\033[43m\] # Yellow
\[\033[44m\] # Blue
\[\033[45m\] # Purple
\[\033[46m\] # Cyan
\[\033[47m\] # White

# High Intensty backgrounds
\[\033[0;100m\] # Black
\[\033[0;101m\] # Red
\[\033[0;102m\] # Green
\[\033[0;103m\] # Yellow
\[\033[0;104m\] # Blue
\[\033[10;95m\] # Purple
\[\033[0;106m\] # Cyan
\[\033[0;107m\] # White

#Replace any leading leading 0; with 1; for bold colors
#Replace any leading 0; with 4; to underline
*/
