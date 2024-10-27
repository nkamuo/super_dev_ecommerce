
type ProductService interface {
	CheckProductAvailability(ctx context.Context, in *CheckProductRequest, opts ...grpc.CallOption) (*CheckProductResponse, error)
	GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error)
	ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductListResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error)
}