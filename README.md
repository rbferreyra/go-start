# Go

Aula básica sobre a linguaguem Go. Principais conceitos e fundamentos da linguagem.

## Tabela de conteúdo

* [Instalação](#instalação)
    * [Instalar a linguagem Go](#1-instalar-a-linguagem-go)
    * [Configurar o ambiente Go](#2-configurar-o-ambiente-go)
    * [Atualizar a sessão atual do shell](#3-atualizar-a-sessão-atual-do-shell)
    * [Verificar instalação](#4-verificar-instalação)



## Instalação

### 1. Instalar a linguagem Go

Ir para a página de [downloads](https://golang.org/dl) e  baixar o arquivo binário conforme o sistema operacional. Para instalar no sistema do Ubuntu, segue os comandos:

```bash
wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
```

Extrai o arquivo baixado e instala no diretório, geralmente localizado sobre o diretório `/usr/local`, sugerido por padrão.

```bash
sudo tar -xvf go1.14.linux-amd64.tar.gz
sudo mv go /usr/local
```

### 2. Configurar o ambiente Go

Vamos configurar as variáveis de ambiente da linguagem Go, como `GOROOT`, `GOPATH` e `PATH`.

`GOROOT` é aonde está localizado o pacotes instalados do Go no seu sistema.

`GOPATH` é aonde está localização do diretório dos projetos realizados.

Abre o arquivo `.profile` e inclui as variáveis globais no fim do arquivo.

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

### 3. Atualizar a sessão atual do shell

```bash
source ~/.profile
```

Isso irá permitir que você use os comandos do Go, sem reiniciar o terminal.

### 4. Verificar instalação

Você tem instalado com sucesso e configurado a linguagem Go no sistema.

Para verificar a versão instalada:

```bash
go version

// go version go1.14.1 linux/amd64
```
