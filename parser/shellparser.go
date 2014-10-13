// shell script parser
package parser

import ()

type Parser struct {
}

// Parse almost all commands, but not complicated one yet
func Parse(line string) (cmds []string, err error) {
	// detect pipes
	constructingCommand := ""
	for idx, s := range line {
		if s == '|' {
			if line[idx-1] == '\\' {
				// it's escaped

			} else {
				// should divide
				cmds = append(cmds, constructingCommand)
				constructingCommand = ""
				continue
			}
		}
		constructingCommand += string(s) // TODO:  fix string new
	}
	cmds = append(cmds, constructingCommand)
	return
}
