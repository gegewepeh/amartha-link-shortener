# SETUP DB
```
CREATE TABLE users (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	"username" VARCHAR ( 50 ) UNIQUE NOT NULL,
	"password" VARCHAR ( 50 ),
	"createdAt" TIMESTAMPTZ NOT NULL,
	"updatedAt" TIMESTAMPTZ NOT NULL
);
```

```
CREATE TABLE links (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	"userId" UUID references users(id) NOT NULL,
	"fullLink" VARCHAR ( 100 ) NOT NULL,
	slug VARCHAR ( 6 ) UNIQUE NOT NULL,
	visit INT4 NOT NULL,
	"createdAt" TIMESTAMPTZ NOT NULL,
	"updatedAt" TIMESTAMPTZ NOT NULL
);
```

## Run Project

- run `go mod tidy`
- run `go mod download`
- copy .env.sample to .env (adjust for local user and password)
- run `export ENV=development`
- run `go run ./cmd/link_shortener.go`


## Available Endpoint

#### Create User
POST {{host}}/link-shortener/v1/slug

Body JSON example:
```
{
    "username": "testing",
    "password": "12345678"
}
```

#### Create Slug
POST {{host}}/link-shortener/v1/slugs

Body JSON example:
```
{
    "username": "testing",
    "password": "12345678",
    "fullLink": "www.google.com"
}
```

#### Show all Links created by user
GET {{host}}/link-shortener/v1/links?username=testing&password=12345678

#### Show full link with slug (will count as 1 visit)
GET {{host}}/link-shortener/v1/slug/:id

#### Update Slug by user
PUT {{host}}/link-shortener/v1/slug/:id

Body JSON example:
```
{
    "username": "testing",
    "password": "12345678"
}
```


## Number of possible unique slugs

### With permutation formula
k! / k! - n!

k: total possible characters
n: length of characters picked

available letters "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" (have 62 length)
k = 62
n = 6 (length of generated characters currently used in the generator)

62! / 62! - 6!
= 62! / 56!
= 62 x 61 x 60 x 59 x 58 x 57
= 44,261,653,680