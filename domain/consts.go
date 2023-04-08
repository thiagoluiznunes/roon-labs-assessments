package domain

var HELP_INFO_TEXT string = `
USAGE:

  rpn                          Launch in interactive mode
  rpn [expression]             Evaluate a one-line expression

RC FILE

  rpn will execute the contents of ~/.rpnrc at startup if it exists.

EXAMPLES

  rpn 1 2 + 3 + 4 + 5 +              => 15
  rpn pi cos                         => -1.0
  rpn                                => interactive mode
`
