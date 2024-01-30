API endpoints for auth

### 1. User sign up.
- Handles the registration or creation of a new user account.

**Authorization** 

| Key          | Value           |
| ------------ | --------------- |
|              |                 |
|              |                 |

**POST** sign up

```{{baseURL}}/api/v1/signup```

**Body** raw (json)
```json
{
    "first_name": "john",
    "middle_name": "",
    "last_name" : "doe",
    "email": "johndoe@gmail.com",
    "phone_number": "<phone number>",
    "password": "password"
}
```

**Response**
```json
{
    "status_code":201,
    "message": "user successfully signed up",
    "data":{},
    "error_text":""
}
```