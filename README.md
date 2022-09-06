
# Currency API Go

Currency API - show currencies & coins price powered by Golang and Redis.

## Supported currencies
```
Currencies:
US Dollar, Euro, British Pound, UAE Dirham, Norwegian Krone, Qatari Rial, Swiss Franc,
Thai Baht, Azerbaijani Manat, Hong Kong Dollar, Danish Krone, Turkish Lira, KSA Rial,
Indian Rupee, Afghan Afghani, Omani Rial, Canadian Dollar, Japanese Yen, Iraqi Dinar, 
Australian Dollar, Singapore Dollar, Chinese Yuan, Swedish Krone, Russian Ruble,
Armenian Dram, Ringgit, Kuwaiti Dinar, Bahraini Dinar

Coins:
Azadi Coin,  Azadi Half a coin, Azadi Quarter a coin, Gerami Coin, Emami Coin,

Gold:
Gram, Ounce, Mithqal,

Crypto currencies:
Bitcoin
```

## Dependencies

- Golang 1.18+
- Redis v7

## Getting started

```bash
$ git clone https://github.com/itsjoniur/CurrencyAPIGo.git
$ cd CurrencyAPIGo
```

- Rename .env.example to .env
- change the Redis info with yours

```bash
$ cd ./cmd/currency
$ go run main.go
```

OR

### Run as a docker container

```bash
$ docker-compose build --no-cache
$ docker-compose up -d

[+] Running 3/3
 ⠿ Container currencyapigo-redis-overcommit-1  Star...                                              0.2s
 ⠿ Container currencyapigo-cache-1             Started                                              0.5s
 ⠿ Container currencyapigo-web-1               Started                                              0.8s
```

## Routes

`GET /`

This route will show this `README` to show the features of the projects to everyone.

`GET /price`

This route will show the currencies' prices.

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

For a specific **currency**:

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

For a specific **coin**:

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

For a specific **gold**:

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

