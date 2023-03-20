# gpgm



## Up

> Images should be re-build when a package is added.

### Dev

```bash
export PORT=2178
docker compose up -d
```

### Production

```
export PORT=2178
docker compose -f docker-compose.prod.yml build base
docker compose -f docker-compose.prod.yml up -d
```

> The app is then ready on ``localhost:2178``.
