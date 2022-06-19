// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// // type ResUser struct {
// // 	ResUser controller.ResUser
// // }

// func main() {
// 	// db := models.DbConnect()

// 	// defer db.Close()
// 	// db.LogMode(true)

// 	// client := resty.New()
// 	// res, _ := client.R().
// 	// 	EnableTrace().
// 	// 	Get("http://api_user:3000/users/db8eae88-e098-4109-88c3-210bdb346562")
// 	// var user controller.ResUser
// 	// if err := json.Unmarshal(res.Body(), &user); err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Printf("%+v\n", user.ResUser.UUID)

// 	router := gin.Default()
// 	// router.GET("/test", controller.Post(user))
//     router.Run(":3001")
// }

package main

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	CompanyName string `json:"companyName"`
	Email  string `json:"email"`
}

func clients(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := Client{
		CompanyName: "John Inc.",
		Email: "john@example.com",
	}
	var clients []Client
	clients = append(clients, user)
	json.NewEncoder(w).Encode(clients)
}

func main() {
	http.HandleFunc("/clients", clients)
	http.ListenAndServe(":3001", nil)
}
