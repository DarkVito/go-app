# API example (in the process of documenting)


<details>
  <summary>register</summary>
  
**URL** : `/api/register/`

**Method** : `POST`

```json
{
    "name":"a",
    "email": "a@aa123.com",
    "password": "1"
}
```

## Success Response

```json
{
    "id": 9,
    "name": "a",
    "email": "a@aa123.com"
}
```

## Error Response


```json
{
    "Number": 1062,
    "Message": "Duplicate entry 'a@aa123.com' for key 'email'"
}
```
</details>

<details>
  <summary>login</summary>
  
**URL** : `/api/login`

**Method** : `POST`

```json
{
    "email":"a@a.com",
    "password": "1"
}
```

## Success Response

```json
{
    "message": "success"
}
```

## Error Response


```json
{
    "message": "user not found"
}
```
</details>
