# GMH

Simple projet back-end pour se familiariser avec le Golang.

## Structure du Projet

```sh
/gmh-api
├── /cmd
│   ├── /myapp
│   │   └── main.go
├── /pkg
│   ├── /mypackage
│   │   └── mypackage.go
├── /api
│   ├── /v1
│   │   ├── /handlers
│   │   ├── /middleware
│   │   └── routes.go
├── /scripts
├── /internal
│   ├── /configs
│   ├── /models
│   ├── /services
│   └── /utils
├── .env 
└── go.mod
```

## Prérequis

- [Go](https://golang.org/dl/) version 1.16 ou supérieure
- Un environnement de développement Go configuré

## Installation

Clonez le dépôt et installez les dépendances :

```sh
git clone https://github.com/c3k4ah/gmh-api.git
cd gmh-api
go mod tidy
```
