package data

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"log"
	"math/rand"
	"social/cmd"
	"social/internal/domain/entities"
	"social/internal/domain/usecase"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func DataForTest() {

	rand.Seed(time.Now().UnixNano())

	container := cmd.BuildContainer()
	ctx := context.Background()
	err := container.Invoke(func(store usecase.UserService) {

		for i := 0; i < 1000000; i++ {

			lastName := gofakeit.LastName()
			firstName := gofakeit.FirstName()
			var b strings.Builder
			b.WriteString(firstName)
			b.WriteString("_")
			b.WriteString(String(4))
			gofakeit.Seed(0)

			user := &entities.User{
				Login:       b.String(),
				Password:    "123456",
				Email:       gofakeit.Email(),
				FirstName:   firstName,
				LastName:    lastName,
				City:        gofakeit.City(),
				Gender:      gofakeit.Gender(),
				Interests:   gofakeit.HipsterSentence(12),
				DateCreated: time.Now(),
				DateModify:  time.Now(),
				Age:         20 + rand.Int31n(80),
			}
			fmt.Println(user.ID, user.FirstName, user.LastName)
			_, err := store.AddUserUseCase(ctx, user)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("iteration= ", i)
		}
	});
	if err != nil {
		log.Println(err)
	}

}
