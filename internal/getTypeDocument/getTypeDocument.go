package GetTypeDocument

import "context"

type Repository interface {
	GetTypeDocumentRepository(ctx context.Context) ([]GetTypeDocumentResponse, error)
}

type Service interface {
	GetTypeDocumentService(ctx context.Context) ([]GetTypeDocumentResponse, error)
}

type GetTypeDocumentResponse struct {
	Id               int    `json:"id"`
	NameTypeDocument string `json:"nameTypeDocument"`
}
