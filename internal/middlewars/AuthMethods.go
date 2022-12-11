package middlewars

//func CheckApp(c *gin.Context, db *sqlx.DB, t string) {
//user := new(models.User)
//token := new(models.AccessToken)
//claims, err := appauth.GetClaims(t)
//if err != nil {
//	if err.(*jwt.ValidationError).Errors == 16 {
//		token.GetByToken(db, t)
//		if token.ID == 0 {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//			return
//		}
//
//		token.Delete(db)
//	}
//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//	return
//}
//
//if err = user.FindByUniqueAndService(db, claims.Unique, claims.Service); err != nil {
//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//	return
//}
//
//token, err = user.GetTokenByStr(db, t)
//if err != nil {
//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//	return
//}
//if token.ID == 0 {
//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
//	return
//}
//c.Set("user", user.UniqueRaw)
//c.Set("service", user.AuthorizedBy)
//}
