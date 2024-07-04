package main

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "go.etcd.io/bbolt"
    "log"
    "net/http"
    "regexp"
    "strconv"
    "time"
)

var db *bbolt.DB

type User struct {
    Username    string    `json:"username"`
    DateOfBirth time.Time `json:"dateOfBirth"`
}

func main() {
    var err error
    db, err = bbolt.Open("users.db", 0600, nil)
    if err != nil {
        log.Fatalf("failed to open database: %v", err)
    }
    defer db.Close()

    db.Update(func(tx *bbolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("Users"))
        return err
    })

    r := gin.Default()

    r.PUT("/hello/:username", putUser)
    r.GET("/hello/:username", getUser)

    r.Run(":8080")
}

func putUser(c *gin.Context) {
    username := c.Param("username")
    if !isValidUsername(username) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
        return
    }

    var input struct {
        DateOfBirth string `json:"dateOfBirth"`
    }

    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    dateOfBirth, err := time.Parse("2006-01-02", input.DateOfBirth)
    if err != nil || dateOfBirth.After(time.Now()) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date of birth"})
        return
    }

    user := User{Username: username, DateOfBirth: dateOfBirth}

    db.Update(func(tx *bbolt.Tx) error {
        bucket := tx.Bucket([]byte("Users"))
        userBytes, _ := json.Marshal(user)
        return bucket.Put([]byte(username), userBytes)
    })

    c.Status(http.StatusNoContent)
}

func getUser(c *gin.Context) {
    username := c.Param("username")

    var user User
    db.View(func(tx *bbolt.Tx) error {
        bucket := tx.Bucket([]byte("Users"))
        userBytes := bucket.Get([]byte(username))
        if userBytes == nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return nil
        }
        json.Unmarshal(userBytes, &user)
        return nil
    })

    if user.Username == "" {
        return
    }

    // Today's date without the time
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)

    // The user's birthday without the year
    birthday := user.DateOfBirth
    birthdayThisYear := time.Date(today.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.Local)

    var message string
	// If the birthday is today
	if today.Equal(birthdayThisYear) {
		message = "Hello, " + user.Username + "! Happy birthday!"
	}  else if today.Before(birthdayThisYear) { 
		// If the birthday has not occurred yet this year
		// Calculate the number of days until the birthday
		daysUntilBirthday := birthdayThisYear.Sub(today).Hours() / 24
		
		// If the birthday is today
		if daysUntilBirthday == 0 {
			message = "Hello, " + user.Username + "! Happy birthday!"
		} else {
			message = "Hello, " + user.Username + "! Your birthday is in " + strconv.Itoa(int(daysUntilBirthday)) + " day(s)"
		}
	} else {
		// If the birthday has already occurred this year
		nextBirthday := birthdayThisYear.AddDate(1, 0, 0)
		daysUntilBirthday := nextBirthday.Sub(today).Hours() / 24
		message = "Hello, " + user.Username + "! Your birthday is in " + strconv.Itoa(int(daysUntilBirthday)) + " day(s)"
	}

    c.JSON(http.StatusOK, gin.H{"message": message})
}

func isValidUsername(username string) bool {
    re := regexp.MustCompile("^[a-zA-Z]+$")
    return re.MatchString(username)
}
