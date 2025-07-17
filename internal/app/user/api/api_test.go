package api

// func TestApiGetAdvert(t *testing.T) { // go test -v -timeout 30s -run ^TestApiGetAdvert$ service-advert/internal/api
// 	req := httptest.NewRequest("GET", "/advert?id=1", nil)

// 	responseRecorder := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetAdvert)
// 	handler.ServeHTTP(responseRecorder, req)

// 	if status := responseRecorder.Code; status != http.StatusOK {
// 		t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
// 	}
// }

// func TestApiPostAdvert(t *testing.T) { // go test -v -timeout 30s -run ^TestApiPostAdvert$ service-advert/internal/api
// 	body := `{"name": "Продажа авто"}`
// 	req := httptest.NewRequest("POST", "/advert", bytes.NewReader([]byte(body)))

// 	t.Log("TestApiPostAdvert", "req", req)
// 	//slog.Info("TestApiPostAdvert", "req", req)

// 	responseRecorder := httptest.NewRecorder()
// 	handler := http.HandlerFunc(PostAdvert)
// 	handler.ServeHTTP(responseRecorder, req)

// 	//if status := responseRecorder.Code; status != http.StatusOK {
// 	status := responseRecorder.Code
// 	//slog.Info("TestApiPostAdvert", "status", status)
// 	t.Log("TestApiPostAdvert", "status", status)
// 	if status != http.StatusOK {
// 		t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
// 	}
// }
