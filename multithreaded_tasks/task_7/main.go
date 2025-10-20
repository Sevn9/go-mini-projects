package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	UserId   int
	UserName string
	UserRole string
}

var jwtSecretKey = []byte("secret_key")

// 1. Создаем НЕЭКСПОРТИРУЕМЫЙ тип (contextKey)
// Пустая структура (struct{}) используется для экономии памяти,
// так как нам нужно только уникальное имя типа.
type contextKey struct{}

// 2. Создаем ЭКЗЕМПЛЯР этого типа, который и будет ключом.
// Он должен быть экспортирован, чтобы его могли использовать внешние пакеты.
var JwtTokenContextKey = contextKey{}

func main() {
	fmt.Println("task_7")

	var wg sync.WaitGroup

	ctx := context.Background()

	currentUser := User{UserId: 1, UserName: "Alice", UserRole: "Admin"}

	ctxWithToken, err := AddJWTToContext(ctx, currentUser.UserId)

	if err != nil {
		panic(err)
	}

	userId, err := ExtractUserIDFromContext(ctxWithToken)

	if err != nil {
		panic(err)
	}

	fmt.Println("userId:", userId)

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		userId, err := ExtractUserIDFromContext(ctxWithToken)
		if err != nil {
			panic(err)
		}

		fmt.Println("userId from goroutine:", userId)

	}(ctxWithToken)

	wg.Wait()

}

func AddJWTToContext(ctx context.Context, userId int) (context.Context, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	//создать токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//подписать токен
	tokenString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		return nil, errors.New("cannot signed in")
	}

	fmt.Printf("token created:\n%s\n", tokenString)

	//добавить токен в контекст
	ctxWithToken := context.WithValue(ctx, JwtTokenContextKey, tokenString)

	return ctxWithToken, nil
}

func ExtractUserIDFromContext(ctx context.Context) (int, error) {
	//достаем значение токена из контекста
	receivedToken, ok := ctx.Value(JwtTokenContextKey).(string)

	if !ok {
		errTokenNotFound := errors.New("jwt token not found in context")
		fmt.Println("LOG: jwt token not found in context")
		return 0, errTokenNotFound
	}

	// парсим и валидируем токен
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		// (Функция обратного вызова для предоставления секретного ключа)
		// Проверяем, что алгоритм токена точно соответствует нашему HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("ожидаемый алгоритм %s, получен %s",
				jwt.SigningMethodHS256.Alg(),
				token.Header["alg"])
		}

		// Возвращаем мой ключ библиотеке jwt
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token: %v", err)
	}

	//достать claim из токена
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("cannot parse claims")
	}

	//достать userId из claim
	userIdFloat, ok := claims["user_id"].(float64) // json числа приходят как float64

	if !ok {
		return 0, errors.New("user_id claim missing or invalid")
	}

	userId := int(userIdFloat)
	return userId, nil
}
