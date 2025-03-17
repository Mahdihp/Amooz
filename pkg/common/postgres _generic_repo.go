package common

const (
	SearchParamsType_Integer  string = "integer"
	SearchParamsType_String   string = "string"
	SearchParamsType_DateTime string = "datetime"
)

type SearchParams struct {
	Type      string
	Key       string
	Value     string
	SortField string
}
type IGenericRepository[T any] interface {
	//Delete(filterParams SearchParams) error
	//Update(filterParams SearchParams, updateDoc T) error
	//Create(newItem T) error
	//FindAll(params []SearchParams, pageIndex, pageSize int) ([]T, error)
}
type GenericRepository[T any] struct {
	//postgres *postgres.PostgresDb
}
