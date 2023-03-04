# Up

## Dev

```bash
export PORT=2178
docker compose up -d
```

# Interactive

```bash
# algo container
docker compose exec algo /bin/sh
# client container
docker compose exec client /bin/sh
```

# Environment

> Images should be re-build when a package is added.