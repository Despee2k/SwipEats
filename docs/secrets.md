# SwipEats GitHub Actions Environment Secrets

This document outlines all secrets used in the deployment workflows for the SwipEats Go server (backend) and frontend.  

- Secrets prefixed with `DCISM_` are used for **backend deployment via SSH**.
- Secrets prefixed with `SSH_` are used for **frontend deployment via SSH**.

## Secrets Table

| Secret Name           | Description |
|-----------------------|-------------|
| `DCISM_HOST`          | IP address or domain of the remote deployment server for backend |
| `DCISM_USERNAME`      | SSH username for connecting to the backend deployment server |
| `DCISM_PASSWORD`      | SSH password for backend authentication |
| `DCISM_PORT`          | SSH port used for the backend server connection |
| `DCISM_DOMAIN`        | Base directory path in the backend server where SwipEats is deployed |
| `SERVER_PORT`         | Port number that the Go backend server should listen on |
| `DB_PORT`             | Database server port (e.g., 5432 for PostgreSQL, 3306 for MySQL) |
| `DB_HOST`             | Host address of the database (e.g., `localhost` or internal IP) |
| `DB_USER`             | Username used to authenticate with the database |
| `DB_PASSWORD`         | Password for the database user |
| `DB_NAME`             | Name of the target database SwipEats uses |
| `SSH_HOST`            | IP address or domain of the remote server for frontend deployment |
| `SSH_USERNAME`        | SSH username for connecting to the frontend deployment server |
| `SSH_PASSWORD`        | SSH password for frontend authentication |
| `SSH_PORT`            | SSH port used for the frontend server connection |
| `JWT_SECRET`          | Secret key used to sign and verify JWTs.


