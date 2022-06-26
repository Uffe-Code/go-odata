package modelGenerator

type generator struct {
	apiUrl string
}

func (g generator) metadataUrl() string {
	return g.apiUrl + "$metadata"
}
