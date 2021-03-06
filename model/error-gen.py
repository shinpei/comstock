#!/bin/python
from string import Template
import os

definedErrors = [
'SessionExpiresError',
'SessionNotFoundError',
'SessionInvalidError',
'UserAlreadyExistError',
'TooWeakPasswordError',
'InvalidMailError',
'UserNotFoundError',
'IncorrectPasswordError',
'AuthenticationFailedError',
'CommandNotFoundError',
'IllegalArgumentError',
'AlreadyLoginError',
'ServerSystemError'
]

codeHeader='''// Data structure used in comstock server, client
// INFO: This file is generated by error-gen.py
package model
'''
snippet = '''type $error struct {
    msg string
}

func (e *$error) Error() string{
    return e.msg
}

func (e *$error) SetError(msg string) *$error{
    e.msg = msg
    return e
}
'''

FILENAME="error-gen.go"
def generate() :
       fo = open(FILENAME, "w")

       fo.write(codeHeader)
       s = Template(snippet)
       for item in definedErrors:
              fo.write(s.substitute(error=item))
                     
       fo.close()

if __name__ == '__main__':
    generate();
    os.system("gofmt " + FILENAME + " > " + "model/error.go")
    os.system("rm " + FILENAME)
