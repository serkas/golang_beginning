## service we are going to deploy with docker


### Build with version
```bash
docker build -f multistaged_ver.Dockerfile --build-arg APP_VERSION=v0.0.1  -t server_multistaged_ver .
```

### Run 

```bash
docker run --rm  -p 8081:8081 server_multistaged_ver
```
