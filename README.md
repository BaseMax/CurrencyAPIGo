# Currency API Go
Currency API - show currencies & coins price powered by Golang

## Dependencies
- Golang 1.18+
- Redis v7
## Getting started
```bash
git clone https://github.com/itsjoniur/CurrencyAPIGo.git
cd CurrencyAPIGo
```
- Rename .env.example to .env
- change the Redis info with yours
```bash
cd ./cmd/currency
go run main.go
```
## Routes
`GET /`
This route will show this README to show the features of the projects to everyone.

`GET /price`
This route will show currencies/currency price
- STRING `q`  (optional)

Example response:
```json
{
   "status": true,
   "currencies": {
	   "usd": {
		   "currency": "US Dollar",
		   "sell": 29800,
		   "buy": 29500
	   },
	   "eur": {
		   "currency": "Euro",
		   "sell": 30000,
		   "buy": 29900
	   }
   },
   "coins" : {
	   "Azadi1/2": {
		   "name": "Azadi 1/2",
		   "sell": 12500
		   "buy": 12300
	   }
   },
   "last_modify": "YYYY-mm-dd HH:MM:SS"
}
```

For specific currency:
```json
{
   "status": true,
   "currency": {
	   "code": "USD",
	   "currency": "US Dollar",
	   "sell": 29800,
	   "buy": 29500
   }
}  
```

For specific coin:
```json
{
   "status": true,
   "coin": {
	   "name": "Azadi 1/2",
	   "sell": 12500,
	   "buy": 12300
   }
}  
```
For specific gold:
```json
{
   "status": true,
   "gold": {
	   "name": "Gram",
	   "price" 10000
   }
}  
```
OR
```json
{
   "status": false,
   "message": "Error message"
}
```

## Authors
- Its Joniur
- Max Base
