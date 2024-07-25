package demo_command

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CommandHandler() http.Handler {
	fmt.Println("aaaaaaa")
	cs := CommandService{executor: CommandExecutor{}}
	cs.helloCommand()
	fn := func(w http.ResponseWriter, r *http.Request) {
		// 檢查 HTTP 方法是否為 GET
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 構建一個簡單的響應數據
		response := map[string]string{"message": "Hello, World!"}

		// 設置內容類型為 JSON
		w.Header().Set("Content-Type", "application/json")

		// 將響應數據編碼為 JSON 並寫入響應
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}

	// 將函數轉換為 http.Handler
	return http.HandlerFunc(fn)
}

func GoCommand() {
	http.Handle("/hello", CommandHandler())
	http.ListenAndServe(":8080", nil)
}
