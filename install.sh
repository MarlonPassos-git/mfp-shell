#!/bin/bash

# Configurações
REPO_OWNER="MarlonPassos-git"
REPO_NAME="mp-shell"
VERSION="v1.0.2"  # Substitua pela versão desejada
BIN_NAME="mp-shell"
DOWNLOAD_URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$VERSION/$BIN_NAME"

# Diretório de instalação padrão e fallback
DEFAULT_INSTALL_DIR="/usr/local/bin"
FALLBACK_INSTALL_DIR="$HOME/.local/bin"

# Baixando o binário
echo "Baixando o binário do mp-shell..."
curl -L -o $BIN_NAME $DOWNLOAD_URL

# Tornando o binário executável
echo "Tornando o binário executável..."
chmod +x $BIN_NAME

# Escolhendo o diretório de instalação
if [ -w $DEFAULT_INSTALL_DIR ]; then
    INSTALL_DIR=$DEFAULT_INSTALL_DIR
else
    echo "Aviso: Você não tem permissão para escrever em $DEFAULT_INSTALL_DIR."
    echo "Instalando em $FALLBACK_INSTALL_DIR."
    mkdir -p $FALLBACK_INSTALL_DIR
    INSTALL_DIR=$FALLBACK_INSTALL_DIR
fi

# Movendo o binário para o diretório escolhido
echo "Movendo o binário para $INSTALL_DIR..."
mv $BIN_NAME $INSTALL_DIR

# Verificando se o diretório do fallback está no PATH
if [ "$INSTALL_DIR" == "$FALLBACK_INSTALL_DIR" ]; then
    if ! echo "$PATH" | grep -q "$FALLBACK_INSTALL_DIR"; then
        echo "Aviso: $FALLBACK_INSTALL_DIR não está no PATH."
        echo "Adicione a linha abaixo ao seu arquivo ~/.bashrc ou ~/.zshrc:"
        echo "export PATH=\$PATH:$FALLBACK_INSTALL_DIR"
    fi
fi

# Verificando a instalação
echo "Verificando a instalação..."
if command -v $BIN_NAME &> /dev/null; then
    echo "$BIN_NAME instalado com sucesso em $INSTALL_DIR!"
else
    echo "Erro: Não foi possível instalar $BIN_NAME."
fi
