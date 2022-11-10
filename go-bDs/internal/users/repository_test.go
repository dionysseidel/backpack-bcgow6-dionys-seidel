package users

import (
	"context"
	"os"
	"testing"

	"github.com/bootcamp-go/go-bDs/internal/domain"
	"github.com/bootcamp-go/go-bDs/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestUpdate_txdb(t *testing.T) {
	userToStore := &domain.User{
		Firstname: "Pablo",
	}
	userToUpdate := domain.User{
		Firstname: "Dionys",
	}

	os.Setenv("DYNAMOID", "local")
	os.Setenv("DYNAMOSECRET", "local")
	os.Setenv("DYNAMOTOKEN", "")
	db, err := db.ConnectDynamoDB()
	assert.NoError(t, err)

	repository := NewRepository(db)

	ctx := context.TODO()

	newID, err := repository.Store(ctx, userToStore)
	assert.NoError(t, err)

	userToStore, err = repository.GetOne(ctx, newID)
	assert.NoError(t, err)

	userToUpdate.Id = userToStore.Id

	err = repository.Update(ctx, userToUpdate)
	assert.NoError(t, err)

	updatedUser, err := repository.GetOne(ctx, userToUpdate.Id)
	assert.NoError(t, err)

	assert.Equal(t, userToUpdate.Id, updatedUser.Id)
}

func TestStore(t *testing.T) {
	userStub := domain.User{
		Firstname: "Dionys",
	}

	os.Setenv("DYNAMOID", "local")
	os.Setenv("DYNAMOSECRET", "local")
	os.Setenv("DYNAMOTOKEN", "")
	db, err := db.ConnectDynamoDB()
	assert.NoError(t, err)

	repository := NewRepository(db)

	ctx := context.TODO()

	newID, err := repository.Store(ctx, &userStub)
	assert.NoError(t, err)

	userSaved, err := repository.GetOne(ctx, newID)
	assert.NoError(t, err)

	assert.Equal(t, newID, userSaved.Id)
}
