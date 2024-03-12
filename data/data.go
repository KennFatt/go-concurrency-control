package data

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"go-cc/ent"
	"time"

	"ariga.io/atlas/sql/sqltool"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(conf *Database) (*ent.Client, error) {
	// see https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.Username, conf.Password, conf.Hostname, conf.Port, conf.Name)

	db, err := sql.Open(conf.Driver, dsn)
	if err != nil {
		return nil, err
	}

	// see https://www.alexedwards.net/blog/configuring-sqldb
	if conf.MaxLifeTime != nil {
		db.SetConnMaxLifetime(*conf.MaxLifeTime)
	}

	if conf.MaxOpenConn != nil {
		// NOTE: adjust this value accordingly to test the concurrent updates
		// 	(amount of requests sent by the client at the same time).

		// Default 2 max open conns
		db.SetMaxOpenConns(*conf.MaxOpenConn)
	}

	if conf.MaxIdleConn != nil {
		// Default unlimited idle conns
		db.SetMaxIdleConns(*conf.MaxIdleConn)
	}

	var (
		driver          = entsql.OpenDB(conf.Driver, db)
		entClientDriver = ent.Driver(driver)

		client *ent.Client
	)

	ctx := context.Background()

	client = ent.NewClient(entClientDriver)

	opts := []schema.MigrateOption{
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	}

	// for debugging purpose, it will run automigration and generate new SQL on each schema updates.
	var buffer bytes.Buffer
	client.Schema.WriteTo(ctx, &buffer, opts...)
	client.Schema.Create(ctx, opts...)
	if len(buffer.Bytes()) > 0 {
		dir, _ := sqltool.NewGolangMigrateDir("migrations")
		now := time.Now().Unix()
		dir.WriteFile(fmt.Sprintf("%d.sql", now), buffer.Bytes())
	}

	return client, nil
}
