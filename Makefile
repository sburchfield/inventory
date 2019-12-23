
export APP_PORT=3000

export APP_PASSWORD_RESET_DOMAIN=http://localhost:3000

export TEMPLATES_FOLDER=./emailtemplates

export MG_API_KEY=3504cf99bb9ca7ecd35ba89235552849-09001d55-835dc64d
export MG_DOMAIN=sandboxdd9222a0f55e4f5bacc19fa520f5c955.mailgun.org
export MG_PUBLIC_API_KEY=pubkey-c672fc7acf8c9048edfa27af379a0e26
# export MG_URL="https://api.mailgun.net/v3/mg.pyaamailservices.com"

build:
	go build

local: export APP_MODE=local
local: export DB_HOST=localhost
local: export DB_NAME=inventory
local: export DB_SCHEMA=inventory
local: export DB_USER=sam
local: export DB_PASSWORD="1030Jaco"
local: export DB_SSLMODE=disable
local:
	./inventory
