#!/bin/bash

# Configurações
REPO_OWNER="MarlonPassos-git"
REPO_NAME="mp-shell"
VERSION="v1.0.2"  # Substitua pela versão desejada
BIN_NAME="mp-shell"
DOWNLOAD_URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$VERSION/$BIN_NAME"

# Baixando o binário
echo "Baixando o binário do mp-shell..."
curl -L -o $BIN_NAME $DOWNLOAD_URL

# Tornando o binário executável
echo "Tornando o binário executável..."
chmod +x $BIN_NAME

# Movendo para /usr/local/bin
echo "Movendo o binário para /usr/local/bin..."
sudo mv $BIN_NAME /usr/local/bin

# Verificando a instalação
echo "Verificando a instalação..."
if command -v $BIN_NAME &> /dev/null; then
    echo "$BIN_NAME instalado com sucesso!"
else
    echo "Erro: Não foi possível instalar $BIN_NAME."
fi