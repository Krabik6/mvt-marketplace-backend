### Requirements
- Docker

### Run
```
docker-compose up --build
```

### Migrate up

```
migrate -database postgres://postgres:qwerty@localhost:5435/postgres?sslmode=disable -path migrations up
```

### Migrate down

```
migrate -database postgres://postgres:qwerty@localhost:5435/postgres?sslmode=disable -path migrations down
```

## Endpoints

### Fill sale (POST)

```
localhost:8000/sale/
```

Body:
```
{
    "NftId":          <NFT_ID(int)>,
    "NftCollection": "<COLLECTION_ADDRESS>",
    "Seller":        "<SELLER_ADDRESS>",
    "Cost":           <COST(int)>,
    "Token":         "<ERC20_TOKEN_ADDRESS>"
}
```

### Get sale info (GET)
```
localhost:8000/sale/?id=<NFT_ID>&collection=<COLLECTION_ADDRESS>
```

