package migration

type Migration interface {
	Up()
	Down()
}
