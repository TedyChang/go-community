package service_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logpanic := func(err error) {
		if err != nil {
			log.Panic(err)
		}
	}

	ctx := context.Background()

	logpanic(os.Setenv("PORT", "26777"))
	logpanic(os.Setenv("PROFILE", "test"))

	req := testcontainers.ContainerRequest{
		Image: "postgis/postgis:13-master",
		Env: map[string]string{
			"POSTGRES_USER":             "postgres",
			"POSTGRES_PASSWORD":         "postgres",
			"POSTGRES_DB":               "postgres",
			"POSTGRES_HOST_AUTH_METHOD": "trust",
			"POSTGRES_HOSTNAME":         "localhost",
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("listen ipv6 :: port 5432"),
	}

	pg, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return
	}

	defer func(pg testcontainers.Container, ctx context.Context) {
		logpanic(pg.Terminate(ctx))
	}(pg, ctx)

	port, _ := pg.MappedPort(ctx, "5432")

	db := fmt.Sprintf("postgres://postgres:postgres@localhost:%d/postgres?sslmode=disable", port.Int())

	logpanic(os.Setenv("DSN", db))

}

func setup(f func()) func(t *testing.T) {
	logpanic := func(err error) {
		if err != nil {
			log.Panic(err)
		}
	}

	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	logpanic(err)

	defer func(db *sql.DB) {
		logpanic(db.Close())
	}(db)

	pwd, _ := os.Getwd()

	basepath := fmt.Sprintf("%s/../", pwd)

	sqlPutter := func(sql string) string {
		dat1, _ := ioutil.ReadFile(fmt.Sprintf("%s/docs/s/"+sql, basepath))
		return string(dat1)
	}

	s := sqlPutter("1_board_domain.s") + ";"
	s += sqlPutter("1_user_domain.s") + ";"
	s += sqlPutter("2_board_data.s") + ";"
	s += sqlPutter("2_user_data.s")

	_, err = db.Exec(s)
	logpanic(err)

	return func(t *testing.T) {
		f()
	}

}
