# Deployment Documentation

This is the documentation for setting up SwipEats in DCISM Server.

## Prerequisite: Environment Set up (Server)

### Step 1: SSH into `web.dcism.org`
Instructions can be found in `admin.dcism.org`

### Step 2:  Clone the repository
```
cd {{ PROJECT_DOMAIN }}
git clone https://github.com/Despee2k/SwipEats.git
cd SwipEats
mv * .[^.]* ../
cd ..
rmdir SwipEats
```

> [!NOTE]
> Replace {{ PROJECT_DOMAIN }} with actual project domain

### Step 3: Configure `.htaccess`
```
DirectoryIndex disabled

RewriteEngine on

RewriteCond %{SERVER_PORT} 80
RewriteRule ^.*$ https://%{HTTP_HOST}%{REQUEST_URI} [R=301,L]

RewriteRule (.*) http://127.0.0.1:{{ SERVER_PORT }}%{REQUEST_URI} [P,L]
```

> [!NOTE]
> You must replace {{SERVER_PORT}} with an allowed port within the range provided in `admin.dcism.org`