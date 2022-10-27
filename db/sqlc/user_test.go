package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/donMutua/tanzuapi/util"
	"github.com/stretchr/testify/require"
)


func CreateRandomUser(t *testing.T) User{
	arg := CreateUserParams{
		Username: faker.Username(),
		HashedPassword: "Secret",
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Role: util.RandomRole(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Role, user.Role)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}


func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Username, user.Username)
	require.Equal(t, user.HashedPassword, user.HashedPassword)
	require.Equal(t, user.FirstName, user2.FirstName)
	require.Equal(t, user.LastName, user2.LastName)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.Role, user2.Role)


	require.WithinDuration(t, user.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, user2.CreatedAt, time.Second)

}

// func TestUpdateUser(t *testing.T) {
// 	// user := CreateRandomUser(t)

// 	arg := UpdateUserParams{
// 		Username: faker.Username(),
// 		HashedPassword: "Secret",
// 		FirstName: faker.FirstName(),
// 		LastName: faker.LastName(),
// 		Email: faker.Email(),
// 		Role: util.RandomRole(),
// 	}

// 	user2, err := testQueries.UpdateUser(context.Background(),arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, arg.FirstName, user2.FirstName)
// 	require.Equal(t, arg.LastName, user2.LastName)
// 	require.Equal(t, arg.Email, user2.Email)
// 	require.Equal(t, arg.Role, user2.Role)

// 	require.NotZero(t, user2.Username)
// 	require.NotZero(t, user2.CreatedAt)
// }

func TestDeleteUser(t *testing.T) {
	user := CreateRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.Username)

	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user.Username)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestDeleteUsers(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	account2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListUsers(t *testing.T){
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}

	arg := ListUsersParams{
		Limit: 5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}



