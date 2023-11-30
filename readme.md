# budget API

### Endpoints

1. Add Transactions

```http request
POST /add-transaction
Content-Type: application/json
Authorization: Basic dXNlcjpwYXNzd29yZA==

{
    "account_id": 1,
    "category_id": 1,
    "transaction_type": "Credit",
    "amount": 56.9,
    "charges": 33.0
}
```

_OK_: HTTP 201

```json
{
  "message": "Ok"
}
```

_NOT OK_: HTTP 401/ HTTP 400
```json
{
  "message": "An error occurred"
}
```

---

2. Read Accounts

```http request
GET /accounts
Authorization: Basic dXNlcjpwYXNzd29yZA==

```

_OK_: HTTP 200
```json
{
  "data": [
    {
      "account_id": 1,
      "name": "Cash"
    },
    {
      "account_id": 2,
      "name": "Mpesa"
    }
  ]
}
```



3. Read Categories

```http request
GET /categories
Authorization: Basic dXNlcjpwYXNzd29yZA==
```

_OK_: HTTP 200

```json
{
  "data": [
    {
      "category_id": 1,
      "name": "Health"
    },
    {
      "category_id": 2,
      "name": "Transport"
    }
  ]
}
```

---

### Data

1. Summary

`period` can be `{day|week|month|year}`
```http request
GET /summary/{period}
Authorization: Basic dXNlcjpwYXNzd29yZA==
```

_OK_:HTTP 200
```json
{
  "start_time": "2023-11-30 00:00:00",
  "end_time": "2023-11-30 23:59:59",
  "label": "Nov 30th",
  "credit": 459.09,
  "debit": 8754.09
}
```


2. Transactions

`period` can be `{day|week|month|year}`

```http request
GET /summary/transactions/{period}
Authorization: Basic dXNlcjpwYXNzd29yZA==
```

_OK_:HTTP 200
```json
{
  "start_time": "2023-11-30 00:00:00",
  "end_time": "2023-11-30 23:59:59",
  "label": "Nov 30th",
  "data": [
    {
      "transaction_id": 1,
      "account": {
        "account_id": 1,
        "name": "Cash"
      },
      "category": {
        "category_id": 1,
        "name": "Health"
      },
      "category_id": 2,
      "amount": 456.99,
      "charges": 45.00,
      "timestamp": "2023-11-30 12:34:34"
    }
  ]
}
```

3. Account Summary

```http request
GET /summary/account/{period}
Authorization: Basic dXNlcjpwYXNzd29yZA==
```

_OK_: HTTP 200
```json
{
  "start_time": "2023-11-30 00:00:00",
  "end_time": "2023-11-30 23:59:59",
  "label": "Nov 30th",
  "data": [
    {
      "account_id": 1,
      "name": "Cash",
      "balance": 569.09,
      "credit": 45958.0,
      "debit": 65674.0,
      "credit_percentage": 45.0,
      "debit_percentage": 43.0
    }
  ]
}
```

4. Category Summary
`period` options: `{day|week|month|year}`

```http request
GET /summary/category/{period}
Authorization: Basic dXNlcjpwYXNzd29yZA==
```

_OK_: HTTP 200

```json
{
  "start_time": "2023-11-30 00:00:00",
  "end_time": "2023-11-30 23:59:59",
  "label": "Nov 30th",
  "data": [
    {
      "category_id": 1,
      "name": "Health",
      "credit": 45958.0,
      "debit": 65674.0,
      "credit_percentage": 45.0,
      "debit_percentage": 43.0
    }
  ]
}

```