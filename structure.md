CashCraft/
├── .gitignore
├── go.mod
├── go.sum
├── README.md
├── cmd/
│   └── main.go
├── configs/
│   └── app.yaml (or .env)
├── internal/
│   ├── controllers/
│   │   ├── auth.go
│   │   └── user.go
│   ├── database/
│   │   ├── gorm.go
│   │   └── migrations.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── logging.go
│   ├── models/
│   │   ├── user.go
│   │   └── response.go
│   ├── repositories/
│   │   └── user_repository.go
│   ├── routes/
│   │   └── api.go
│   └── utils/
│       ├── crypto.go
│       └── validators.go
├── pkg/
│   └── errors/
│       └── api_errors.go
├── views/
│   ├── auth/
│   │   ├── login.html
│   │   └── register.html
│   ├── dashboard/
│   │   ├── index.html
│   │   └── partials/
│   └── shared/
│       ├── layout.html
│       └── components/
├── public/
│   ├── css/
│   ├── js/
│   └── images/
├── scripts/
│   └── migrate.sh
└── server.go
