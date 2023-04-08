package domain

var PI float64 = 3.14
var HELP_INFO_TEXT string = `
USAGE:

  rpn                          Launch in interactive mode
  rpn [expression]             Evaluate a one-line expression
  rpn q                        To exit 

EXAMPLES

  rpn 1 2 + 3 + 4 + 5 +              => 15
  rpn pi cos                         => -1.0
  rpn                                => interactive mode
`

var GRACEFUL_SHUTDOWN_TEXT string = `
Thanks for using our RPN (Rever Polish Notation) calculator!

For more information, visit our project. Link below:
https://github.com/thiagoluiznunes/roon-labs-assessments

                       .-.
                      |_:_|
                     /(_Y_)\
.                   ( \/M\/ )    I'm not your father!
 '.               _.'-/'-'\-'._
   ':           _/.--'[[[[]'--.\_
     ':        /_'  : |::"| :  '.\
       ':     //   ./ |oUU| \.'  :\
         ':  _:'..' \_|___|_/ :   :|
           ':.  .'  |_[___]_|  :.':\
            [::\ |  :  | |  :   ; : \
             '-'   \/'.| |.' \  .;.' |
             |\_    \  '-'   :       |
             |  \    \ .:    :   |   |
             |   \    | '.   :    \  |
             /       \   :. .;       |
            /     |   |  :__/     :  \\
           |  |   |    \:   | \   |   ||
          /    \  : :  |:   /  |__|   /|
      snd |     : : :_/_|  /'._\  '--|_\
          /___.-/_|-'   \  \
                         '-'
`
