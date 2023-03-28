# QuickSign

This repository contains backend code that helds QuickSign's Server-Side Application

## Requirements

- Golang >= 1.18.3
- PostgreSQL

# Environment Preparation

- Configure your Gmail to enable app password access :

```
1. Create a GMail account if you don't have any
2. Enable 2 Way Verification
3. Access App Password > Create Custom App > Acquire The Password > Fill to .env file
```

- Create a .env file containing these configs

```
CONFIG_SMTP_HOST = "smtp.gmail.com"
CONFIG_SMTP_PORT = 587
CONFIG_SENDER_NAME = "youremail@gmail.com"
CONFIG_AUTH_EMAIL = "youremail@gmail.com"
CONFIG_AUTH_PASSWORD = "your password"
DATABASE_URL=postgresql://${{ PGUSER }}:${{ PGPASSWORD }}@${{ PGHOST }}:${{ PGPORT }}/${{ PGDATABASE }}
PGHOST=
PGDATABASE=
PGPASSWORD=
PGPORT=
PGUSER=postgres
PORT=<APP_PORT> -> example 8080
JWT_EXPIRY_HOUR=<your_desired_hour_of_expiration>h -> example 24h
JWT_SECRET=<yourjwtsecret>
URL=http://localhost
```

- Create a PostgreSQL Database named `quicksign` or whatever you decided as long as it matches the .env file

## Usage and Installation

- Install the required dependencies using this command :

  > go mod tidy

- Run the project using this command :
  > go run main.go

## Endpoint Documentation

You can access the documentation on this <a href="">Link</a>
