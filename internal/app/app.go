package app

import (
	"context"
	"log"
	"mono-tracker/internal/domain/fundraising"
	"mono-tracker/pkg"
)

// App struct
type App struct {
	ctx                context.Context
	db                 *pkg.SQLiteClient
	fundraisingService fundraising.IFundraisingService
}

// NewApp creates a new App application struct
func NewApp(db *pkg.SQLiteClient) *App {
	return &App{
		fundraisingService: fundraising.NewFundraisingService(db),
		db:                 db,
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	err := a.prepareDB(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	a.ctx = ctx
}

func (a *App) prepareDB(ctx context.Context) error {
	_, err := a.db.Db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS fundraising (id INTEGER PRIMARY KEY, name TEXT, description TEXT, goal FLOAT, url TEXT)")
	if err != nil {
		return err
	}

	_, err = a.db.Db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS fundraising_history (id INTEGER PRIMARY KEY, fundraising_id INTEGER, raised FLOAT, sync_time TEXT)")

	return nil
}
