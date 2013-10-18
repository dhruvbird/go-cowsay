package cowsay

import (
       // "fmt"
)



func format(text string) string {
     cowFooter := `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`

     return text + cowFooter
}
