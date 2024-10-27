// handlers/order.go
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/superdev/ecommerce/gateway/grpc"
	"github.com/superdev/ecommerce/gateway/proto"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	client := grpc.NewOrderClient()
	var req proto.CreateOrderRequest
	json.NewDecoder(r.Body).Decode(&req)
	res, err := client.CreateOrder(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}
