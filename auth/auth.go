package auth

type AuthRecieve struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserToken
}

type UserToken struct {
	Token     string `json:"token"`
	ExpiresAt int    `json:"expiresAt"`
}

// func Authenticate(
// 	jwtAuthSecret string,
// 	jsonMsg json.RawMessage,
// ) (json.RawMessage, error) {
// 	var (
// 		ret     json.RawMessage
// 		resp    AuthResponse
// 		recieve AuthRecieve
// 		err     error
// 		userID  int
// 	)

// 	err = json.Unmarshal(jsonMsg, &recieve)
// 	if err != nil {
// 		log.Errorf("authenticate JSON unmarshal error: %v", err)
// 		return nil, err
// 	}

// 	userID, err = ValidateDBUser(recieve.Username, recieve.Password)
// 	if err != nil {
// 		log.Errorf("invalid user: %v", err)
// 		return nil, err
// 	}

// 	resp.UserToken, err = CreateToken(jwtAuthSecret, userID)
// 	if err != nil {
// 		log.Errorf("unable to create token: %v", err)
// 		return nil, err
// 	}

// 	// Response
// 	ret, err = json.Marshal(resp)
// 	if err != nil {
// 		log.Errorf("authenticate JSON Marshal error: %v", err)
// 		return nil, err
// 	}

// 	return ret, err
// }

// func ValidateDBUser(
// 	username string,
// 	password string,
// ) (int, error) {
// 	var (
// 		userID int
// 		err    error
// 	)

// 	// TODO check DB
// 	// recieve.Username
// 	// recieve.Password

// 	// STUB
// 	userID = 1

// 	return userID, err
// }
