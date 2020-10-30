package order

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"github.com/micro/go-micro/v2/client"
	order "github.com/3Rivers/order/proto/order"
)

func OrderCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("new")
	// call the backend service
	orderClient := order.NewOrderService("go.micro.service.order", client.DefaultClient)
	rsp, err := orderClient.Call(context.TODO(), &order.Request{
		Id: request["id"].(int64),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"username": rsp.User,
		"goods": rsp.Goods,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
