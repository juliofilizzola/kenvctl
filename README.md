# kenvctl

Ferramenta CLI para gerenciar secrets no Kubernetes e AWS Secrets Manager.

## Instalação

### Instalação rápida (Go 1.25+)

```sh
go install github.com/juliofilizzola/kenvctl/cmd/kenvctl@latest
```

### Build manual

```sh
git clone https://github.com/juliofilizzola/kenvctl.git
cd kenvctl
make build
```

O binário `kenvctl.exe` será gerado na pasta atual (Windows). Para Linux/Mac, use os targets `make build-linux` ou `make build-mac`.

### Releases

Baixe o binário mais recente em [Releases](https://github.com/juliofilizzola/kenvctl/releases).

## Uso

```sh
kenvctl --help
```

Exemplo para criar uma secret:
```sh
kenvctl add
```

## Requisitos
- Go 1.25+
- kubectl instalado e configurado
- AWS CLI instalado e configurado (para AWS Secrets Manager)

## Contribuição
Pull requests são bem-vindos!

## Licença
MIT
