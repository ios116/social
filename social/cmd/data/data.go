package data

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"math/rand"
	"social/internal/domain/entities"
	"time"
)

func DataForTest()  {

	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(0)


	user := entities.User{
		Login:       gofakeit.FirstName(),
		Password:     "123456",
		Email:       gofakeit.Email(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		City:        gofakeit.City(),
		Gender:      gofakeit.Gender(),
		Interests:   gofakeit.Digit(),
		Age:         20+rand.Int31n(80),
	}

	fmt.Println("",user.FirstName, user.LastName)
	fmt.Println("City",user.City)
	fmt.Println("Gender",user.Gender)
	fmt.Println("Interest",user.Interests)

}


