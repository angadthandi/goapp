package auth

// func TestAuthenticate(t *testing.T) {
// 	q := AuthRecieve{
// 		Username: "test",
// 		Password: "test",
// 	}

// 	// Test Request
// 	req, err := json.Marshal(q)
// 	if err != nil {
// 		log.Errorf("test authenticate JSON Marshal error: %v",
// 			err)
// 		return
// 	}

// 	b, err := Authenticate(req)
// 	if err != nil {
// 		log.Errorf("test authenticate error: %v", err)
// 		return
// 	}

// 	var response AuthResponse
// 	err = json.Unmarshal(b, &response)
// 	if err != nil {
// 		log.Errorf("test authenticate JSON unmarshal error: %v",
// 			err)
// 		return
// 	}

// 	fmt.Printf("Authenticate response: %v", response)
// }
