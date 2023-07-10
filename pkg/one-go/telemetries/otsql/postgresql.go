package otsql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/luna-duclos/instrumentedsql"
	"github.com/wisnuragaprawida/project/pkg/one-go/telemetries/otsql/opentracing"
	"github.com/wisnuragaprawida/project/pkg/one-go/utils"
)

var otPgDriver = "opentracing-psotgresql"

func SqlMustConnect(dsn string) *sqlx.DB {

	// checking first in available drivers
	if !utils.Contains(sql.Drivers(), otPgDriver) {
		sql.Register(otPgDriver, instrumentedsql.WrapDriver(
			&pq.Driver{},
			instrumentedsql.WithTracer(opentracing.NewTracer(true)),
			instrumentedsql.WithOmitArgs(),
			instrumentedsql.WithOpsExcluded(opsExluded...),
		))
	}

	db, err := sqlx.Connect(otPgDriver, dsn)
	if err != nil {
		panic(err)
	}
	return db
}
