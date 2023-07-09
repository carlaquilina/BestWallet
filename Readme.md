# BestWallet

BestWallet is a flagship eWallet product developed by BestFinance. It's designed to be highly scalable and offers the capability to serve multiple concurrent requests. It includes features like account creation, transactions between accounts, depositing and withdrawing funds, and much more.

## Features

- Request creation of personal accounts
- Access personal account data
- Deposit and withdraw funds
- Notification system after each operation
- Integrated third party verification and approval (KYC, KYT)

## Tech Stack

- Golang
- Docker
- Postgres
- Gin

## Setup and Installation

Make sure you have Docker and Docker Compose installed on your machine.

1. Clone the repository:
   ```
   git clone https://github.com/carlaquilina/BestWallet.git
   ```

2. Navigate into the cloned directory:
   ```
   cd your_repository_name
   ```

3. Build the Docker images and start the containers:
   ```
   docker-compose up --build
   ```

The application will be available at `http://localhost:8080`.

## Usage

### Account creation
Endpoint: `POST /api/v1/account`

Sample request body:
```json
{
  "FirstName": "John",
  "LastName": "Doe",
  "Email": "john.doe@email.com",
  "Address": "123 Street, City, Country",
  "BirthDate": "1995-12-12"
}
```

### Retrieve Account
Endpoint: `GET /api/v1/account/:id`

### Deposit Funds
Endpoint: `PUT /api/v1/account/:id/deposit`

Sample request body:
```json
{
  "amount": 500
}
```

### Withdraw Funds
Endpoint: `PUT /api/v1/account/:id/withdraw`

Sample request body:
```json
{
  "amount": 300
}
```

### Create Transaction
Endpoint: `POST /api/v1/transaction`

Sample request body:
```json
{
  "FromAccountID": 1,
  "ToAccountID": 2,
  "Amount": 150.75
}
```

## Testing

Run the tests using the following command:

```
go test ./...
```


## API Docs
https://documenter.getpostman.com/view/3624985/2s93zH1yYt
