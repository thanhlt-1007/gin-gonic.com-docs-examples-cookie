package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func getCookieHandler(context *gin.Context) {
    cookie, err := context.Cookie("gin_cookie")
    if err != nil {
        cookie = "NotSet"
        context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
    }

    context.JSON(
        http.StatusOK,
        gin.H {
            "cookie": cookie,
        },
    )
}

func main() {
    engine := gin.Default()
    engine.GET("/cookie", getCookieHandler)
    engine.Run()
}