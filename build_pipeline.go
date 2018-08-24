package pipeline

//go:generate echo "installing go dependancy"
//go:generate go get -v github.com/rubenv/sql-migrate/...

//go:generate echo "running migration"
//go:generate sql-migrate up -env=development
