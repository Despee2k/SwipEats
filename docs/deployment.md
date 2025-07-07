# Deployment Documentation

This is the documentation for setting up SwipEats in DCISM Server.

## Prerequisite: Environment Set up (Server)

### Step 1: SSH into `web.dcism.org`

Instructions can be found in `admin.dcism.org`

### Step 2: Clone the repository (sparse checkout)

```
cd {{ PROJECT_DOMAIN }}

git clone --filter=blob:none --no-checkout https://github.com/Despee2k/SwipEats.git
cd SwipEats

git sparse-checkout init --cone
git sparse-checkout set server
git checkout main

mv * .[^.]* ..
cd ..
rmdir SwipEats
```

> \[!NOTE]
> This clones only the `SwipEats/server` directory to save space and simplify deployment.

### Step 3: Configure `.htaccess`

```
DirectoryIndex disabled

RewriteEngine on

RewriteCond %{SERVER_PORT} 80
RewriteRule ^.*$ https://%{HTTP_HOST}%{REQUEST_URI} [R=301,L]

RewriteRule (.*) http://127.0.0.1:{{ SERVER_PORT }}%{REQUEST_URI} [P,L]
```

> \[!NOTE]
> You must replace `{{ SERVER_PORT }}` with an allowed port within the range provided in `admin.dcism.org`