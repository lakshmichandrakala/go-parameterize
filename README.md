# go-parameterize
Go implementation of Ruby's string parameterize

Replaces contiguous special characters in a string with a '-' and converts all alphabets to lower case. When there are only special characters in the original string, the result of parameterize is an empty string.

Examples:
Parameterize("New York") // "new-york"
Parameterize("$#@") // ""
Parameterize("$#@1#$%") // "1"


