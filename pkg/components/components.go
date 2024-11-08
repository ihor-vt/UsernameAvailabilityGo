package components

import (
	"fmt"

	"github.com/ihor-vt/UsernameAvailabilityGo/pkg/style"
)

var Header = `
 _   _
| | | |___  ___ _ __ _ __   __ _ _ __ ___   ___
| | | / __|/ _ \ '__| '_ \ / _` + "`" + ` | '_ ` + "`" + ` _ \ / _ \
| |_| \__ \  __/ |  | | | | (_| | | | | | |  __/
 \___/|___/\___|_|_ |_| |_|\__,_|_| |_| |_|\___|     ____
   / \__   ____ _(_) | __ _| |__ (_) (_) |_ _   _   / ___| ___
  / _ \ \ / / _` + "`" + ` | | |/ _` + "`" + ` | '_ \| | | | __| | | | | |  _ / _ \
 / ___ \ V / (_| | | | (_| | |_) | | | | |_| |_| | | |_| | (_) |
/_/   \_\_/ \__,_|_|_|\__,_|_.__/|_|_|_|\__|\__, |  \____|\___/
                                            |___/
`

var Prompt = fmt.Sprintf("  Enter a username to check availability %s", style.Dim.Colorize("[q to quit]"))

var TextInput = "  @ "
