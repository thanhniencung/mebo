#### How to build
```
    1. install Docker > https://docs.docker.com/docker-for-windows/install/
    2. clone this repo and cd to mebo folder
    2. run command: make build
    3. Done :)
```

#### How to use 

##### 1. Config
```
    go to > mebo/cmd/main.go and fill your config
    const (
    	MAIL_FROM = "" // your email
    	MAIL_PASS = "" // your email password
    )
```
##### 2. Send an email
    
```
    Call api
    POST: http://localhost:3000/mail-service/send
    
    Params:
    {
        "to": "code4func@gmail.com",
        "subject": "Email marketing",
        "name": "Ryan"
    }
```

##### 3. Mailbox history

```
    Call api
    GET: http://localhost:3000/mail-service/list
    
    Response:
    {
        "status": 200,
        "message": "OK",
        "data": [
            {
                "id": "5cde9f197be36efae664b325",
                "name": "Ryan",
                "email": "ryan.vinova@gmail.com",
                "date": "2019-05-17T11:46:33.894Z"
            }
        ]
    }
```

##### 4. Open web ui
```
    Open web browser with url http://localhost:3000
```

##### 5. check email template
```
    you can check email template at mebo/template/mail_template.html
