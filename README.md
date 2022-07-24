# API implementation of the password manager
## HTTP Methods
```
"GET" /ping — Checking the server connection

    example: 
        "GET" :80/ping
```
```
"GET" / — Checking the server connection

    example: 
        "GET" :80/
```
```
"GET" /generete — Generates a new password
    options:
        characters — Set of letters: 
            u — Uppercase 
            l — Lowercase
            n — Numbers
            s — Symbols
        example: 
            "GET" :80/generete?characters=ul
            "GET" :80/generete?characters=ulns
            "GET" :80/generete?characters=n
```
### Params for ```.env``` file
```
HOST=127.0.0.1
PORT=80
```