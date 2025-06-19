# URLs-AC 🚀🔗

Welcome to **URLs-AC**!  
A simple and efficient URL management project.

## Features ✨
- Fast and lightweight
- Built with Go 🐹
- Postgres for storage
- Redis to cache hot URLs
- Prometheus/Grafana for monitoring
- Kubernetes deployment ready
- TLS
- DB migrations with `go-migrate`



## Running in Development 🛠️
1. Clone the repository:
   ```bash
   git clone git@github.com:cweiser22/urls-ac.git
   ```

2. Run make build-dev:
   ```bash
   make build-dev
   ```
   
3. Start the application:
   ```bash
    make dev
    ```

## Simulating Production 🚀
1. Create directory `./ssl`:
```bash
mkdir ssl 
```

2. Get real SSL certs and place them in `./ssl'`, or generate self-signed certificates with OpenSSL:
   ```bash
   openssl ecparam -genkey -name prime256v1 -out url_ac_ecdsa.key
   openssl x509 -in url_ac_ecdsa.crt -text -noout

   ```
   
3. Build the application:
   ```bash
   make build-prod
   ```
   
4. Start the application:
   ```bash
    make prod
    ```
   
   



## Contributing 🤝

Pull requests are welcome! For major changes, please open an issue first.

## License 📄

[MIT](LICENSE)

---
Made by Cooper Weiser
