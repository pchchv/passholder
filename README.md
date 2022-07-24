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
        lehgth — Password length(optional)
        example: 
            "GET" :80/generete?characters=ul
            "GET" :80/generete?characters=ulns&length=33
            "GET" :80/generete?characters=n&length=5
```
### Params for ```.env``` file
```
HOST=127.0.0.1
PORT=8000
```