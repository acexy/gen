package tests_test

import (
	"context"
	"sync"

	"github.com/acexy/gen/tests/.expect/dal_test/query"
	relquery "github.com/acexy/gen/tests/.expect/dal_test_relation/query"
)

var useOnce sync.Once
var ctx = context.Background()

func CRUDInit() {
	query.Use(DB)
	query.SetDefault(DB)
	relquery.Use(DB)
	relquery.SetDefault(DB)
}
