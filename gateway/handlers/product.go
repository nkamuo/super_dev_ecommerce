// handlers/product.go
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/superdev/ecommerce/gateway/grpc"
	"github.com/superdev/ecommerce/gateway/proto"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	client := grpc.NewProductClient()
	res, err := client.ListProducts(context.Background(), &proto.Empty{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res.Products)
}
