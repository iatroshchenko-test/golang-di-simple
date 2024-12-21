package dicontainer

import (
	"github.com/goioc/di"
	"goapi/src/mymodule/application/use_case/list_posts"
	"goapi/src/mymodule/application/use_case/list_users"
	"goapi/src/shared/infrastructure/db"
	"reflect"
)

func BuildContainer() {
	_, _ = di.RegisterBeanInstance("db", db.Connect())
	_, _ = di.RegisterBean("listUsers", reflect.TypeOf((*list_users.ListUsers)(nil)))
	_, _ = di.RegisterBean("listPosts", reflect.TypeOf((*list_posts.ListPosts)(nil)))
	_ = di.InitializeContainer()
}
