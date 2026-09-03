# libremail
weird server implementation of a mail protocol

intented to be run with a client that implements key hashing indepentend of the server side secret key

has lots of security issues don't use this seriously lmfao

<details>
<summary>endpoint documentation</summary>


### POST `{endpoint}/register`

example body
```
"username": "john doe"
```

example response
```
"Message": "Your secret key is voEEM7.6n3hB22eH note it down and do not lose it...."
```

### POST `{endpoint}/sendmail`

example body
```
"mail": {
        "title": "important discussions",
        "contents": "what's up"
},
"sender": "john doe",
"assignee": "jane doe",
"secretKey": "voEEM7.6n3hB22eH"
```

example response
`{body.mail} echoed back`

### QUERY `{endpoint}/mail`

example body
```
"username": "jane doe",
"secretKey": "pretend i could be bothered to gen another key here"
```

example response
```
[{
    "title": "important discussions",
    "contents": "what's up",
    "author": "john doe"
}]
```

</details>

### flaws

- messages are entirely anonymous (simple implementation change but I never did it)
- mail is stored unencrypted on the server side if a custom client does not independently hash the information before sending it
- data is stored in json, so terrible write speeds and scalability
- vulnerable to db race conditions
- no client implementation as of now