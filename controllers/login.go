package controllers

import (
    "fmt"
    "net/http"
    "sirjtaa/utils"
    "sirjtaa/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
    if utils.GetCurrentUser(r) != nil {
        http.Redirect(w, r, "/", http.StatusFound)
        return
    }

    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")

        var error string
        err, _ := models.AuthenticateUser(username, password)
        if err != nil {
            error = "Invalid username or password"
        }

        if error != "" {
            utils.RenderTemplate(w, r, "login.html", map[string]interface{} {
                "error": error,
            })
            return
        }
        
        session, _ := utils.Store.Get(r, "sirsid")
        session.Values["username"] = username
        session.Values["password"] = password
        fmt.Println("[notice] Auth success!")

        err = session.Save(r, w)
        if err != nil {
            fmt.Printf("[error] Could not save session (%s)\n", err.Error())
        }

        http.Redirect(w, r, "/", http.StatusFound)
        return
    }

    utils.RenderTemplate(w, r, "login.html", nil)
}