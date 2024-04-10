package v1

// func JWTMiddleware(server *ApiV1Service, next echo.HandlerFunc, s string) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := c.Request().Context()
// 		path := c.Request().URL.Path
// 		method := c.Request().Method

// 		fmt.Println(path)

// 		//skip it for auth api's. i have to add /api prefix to those api's so that i can catch them here.
// 		//auth api's dosent need jwtauthentication because we are providing id's there
// 		// if server.defaultAuthSkipper(c) {
// 		// 	return next(c)
// 		// }

// 		return next(c)
// 	}
// }
