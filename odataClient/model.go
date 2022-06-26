package odataClient

type ODataModelDefinition[T any] interface {
	Name() string
	Url() string
}

type ODataModelCollection[T any] interface {
	ODataModelDefinition[T]
	DataSet(ODataClient) ODataDataSet[T, ODataModelDefinition[T]]
}
