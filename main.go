package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "github.com/joho/godotenv"
)

var (
    oauthStateString = "random"
    googleOauthConfig oauth2.Config
)

func init() {
    // Check if .env file exists
    if _, err := os.Stat(".env"); err == nil {
        // Load environment variables from .env file
        err := godotenv.Load()
        if err != nil {
            fmt.Println("Error loading .env file:", err)
        }
    }

    // Initialize Google OAuth config
    googleOauthConfig = oauth2.Config{
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        RedirectURL:  os.Getenv("REDIRECT_URI"),
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
            "openid",
            "https://www.googleapis.com/auth/userinfo.profile",
        },
        Endpoint: google.Endpoint,
    }
}

func main() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    // Load HTML templates
    templatesPath := "templates/*"
    if _, err := filepath.Glob(templatesPath); err != nil {
        fmt.Printf("Error loading templates: %v\n", err)
        return
    }
    router.LoadHTMLGlob(templatesPath)
    router.Static("/static", "./static")

    // API endpoint to get the API key
    router.GET("/api/key", func(c *gin.Context) {
        apiKey := "os.Getenv("OPENWEATHER_API_KEY")"
        c.JSON(http.StatusOK, gin.H{"apiKey": apiKey})
    })

    router.GET("/", index)
    router.GET("/login", login)
    router.GET("/callback", callback)
    router.GET("/logout", logout)

    // Start the server
    router.Run("0.0.0.0:5000")
}

func index(c *gin.Context) {
    if _, err := c.Cookie("google_id"); err != nil {
        c.HTML(http.StatusOK, "signin.html", nil)
    } else {
        name, err := c.Cookie("name")
        if err != nil {
            name = "Guest" // Fallback if cookie is not available
        }
        c.HTML(http.StatusOK, "index.html", gin.H{
            "name": name,
        })
    }
}

func login(c *gin.Context) {
    fmt.Println("Redirect URL:", googleOauthConfig.RedirectURL) // Debugging output
    url := googleOauthConfig.AuthCodeURL(oauthStateString)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

func callback(c *gin.Context) {
    if c.Query("state") != oauthStateString {
        c.Redirect(http.StatusTemporaryRedirect, "/login")
        return
    }

    code := c.Query("code")
    token, err := googleOauthConfig.Exchange(c, code) // Use c instead of oauth2.NoContext
    if err != nil {
        c.Redirect(http.StatusTemporaryRedirect, "/login")
        return
    }

    client := googleOauthConfig.Client(c, token) // Use c instead of oauth2.NoContext
    response, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
    if err != nil {
        c.Redirect(http.StatusTemporaryRedirect, "/login")
        return
    }
    defer response.Body.Close()

    var userInfo map[string]interface{}
    json.NewDecoder(response.Body).Decode(&userInfo)

    c.SetCookie("google_id", userInfo["id"].(string), 3600, "/", "", false, true)
    c.SetCookie("name", userInfo["name"].(string), 3600, "/", "", false, true)

    c.Redirect(http.StatusTemporaryRedirect, "/")
}

func logout(c *gin.Context) {
    c.SetCookie("google_id", "", -1, "/", "", false, true)
    c.SetCookie("name", "", -1, "/", "", false, true)
    c.Redirect(http.StatusTemporaryRedirect, "/")
}