# URLs-AC 🚀🔗

Welcome to **URLs-AC**!  
A simple and efficient URL management project.

## Features ✨
- Fast and lightweight
- Built with Go 🐹
- Postgres for storage
- Redis to cache hot URLs
- Prometheus/Grafana for monitoring


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
1. Create direct `./ssl`

2. Generate fake certs in `./ssl'`:
   ```bash
   openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
   -keyout privkey.pem \
   -out fullchain.pem \
   -subj "/CN=localhost" \
   -addext "subjectAltName=DNS:localhost"
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
