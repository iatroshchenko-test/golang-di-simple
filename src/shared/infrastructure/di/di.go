package di

import (
	"go.uber.org/dig"
	"goapi/src/mymodule/application/use_case/list_users"
	"goapi/src/shared/infrastructure/db"
	"gorm.io/gorm"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Регистрируем базу данных
	container.Provide(func() *gorm.DB {
		return db.Connect()
	})

	// Регистрируем хендлер через его конструктор
	container.Provide(list_users.NewListUsers)

	return container
}
