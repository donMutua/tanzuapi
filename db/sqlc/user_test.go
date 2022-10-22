package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/donMutua/tanzuapi/util"
	"github.com/stretchr/testify/require"
)


func CreateUser(t *testing.T) User{
	arg := CreateUserParams{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Role: util.RandomRole(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Role, user.Role)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateUser(t)
}


func TestGetUser(t *testing.T) {
	user := CreateUser(t)

	user2, err := testQueries.GetUser(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.FirstName, user2.FirstName)
	require.Equal(t, user.LastName, user2.LastName)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.Role, user2.Role)


	require.NotZero(t, user2.ID)
	require.NotZero(t, user2.CreatedAt)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateUser(t)

	arg := UpdateUserParams{
		ID: user1.ID,
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Role: util.RandomRole(),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.FirstName, user2.FirstName)
	require.Equal(t, arg.LastName, user2.LastName)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, arg.Role, user2.Role)

	require.NotZero(t, user2.ID)
	require.NotZero(t, user2.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
	user := CreateUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)

	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestDeleteUsers(t *testing.T) {
	account1 := CreateUser(t)
	err := testQueries.DeleteUser(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListUsers(t *testing.T){
	for i := 0; i < 10; i++ {
		CreateUser(t)
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



