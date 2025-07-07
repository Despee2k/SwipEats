# SwipEats GitHub Actions Environment Secrets

This table documents all secrets used in the deployment workflows for the SwipEats Go server and environment variable injection.

| Secret Name           | Description |
|-----------------------|-------------|
| `DCISM_HOST`          | IP address or domain of the remote deployment server |
| `DCISM_USERNAME`      | SSH username for connecting to the deployment server |
| `DCISM_PASSWORD`      | SSH password for authentication (consider switching to key-based auth) |
| `DCISM_PORT`          | SSH port used for the connection |
| `DCISM_DOMAIN`        | Base directory path in the server where SwipEats is deployed |
| `SERVER_PORT`         | Port number that the Go server should listen on |
| `DB_PORT`             | Database server port (e.g., 5432 for PostgreSQL, 3306 for MySQL) |
| `DB_HOST`             | Host address of the database (e.g., `localhost` or internal IP) |
| `DB_USER`             | Username used to authenticate with the database |
| `DB_PASSWORD`         | Password for the database user |
| `DB_NAME`             | Name of the target database SwipEats uses |
