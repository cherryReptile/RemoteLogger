## Docker Compose Example
```yaml
version: '3'

services:
  auth:
    image: orendat/ws-auth:v1.1.1
    depends_on:
      - db
    environment:
      DB_HOST: db
      DB_PORT: 5432
      DB_USER: db
      DB_PASSWORD: db
      DB_NAME: db
      GIN_MODE: release
      GOOGLE_CLIENT_ID: ''
      GOOGLE_CLIENT_SECRET: ''
      GITHUB_CLIENT_ID: ''
      GITHUB_CLIENT_SECRET: ''
      TG_BOT_TOKEN: ''
      DOMAIN: ''
      JWT_KEY: ''
```
## Environment
`DB_HOST` - postgres host  
`DB_PORT` - postgres port  
`DB_USER` - postgres user  
`DB_PASSWORD` - postgres password  
`DB_NAME` - database name  
`GIN_MODE` - mode http server (debug/release)  
`GOOGLE_CLIENT_ID` -  
`GOOGLE_CLIENT_SECRET` -  
`GITHUB_CLIENT_ID` -  
`GITHUB_CLIENT_SECRET` -  
`TG_BOT_TOKEN` -  
`DOMAIN` -  
`JWT_KEY` -  