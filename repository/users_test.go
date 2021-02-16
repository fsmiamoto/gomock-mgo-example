package repository_test

import (
	"testing"
	"time"

	"github.com/fsmiamoto/gomock-mgo-example/mocks/db"
	"github.com/fsmiamoto/gomock-mgo-example/models"
	"github.com/fsmiamoto/gomock-mgo-example/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestUserRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	session := db.NewMockSession(ctrl)
	database := db.NewMockDatabase(ctrl)
	collection := db.NewMockCollection(ctrl)
	query := db.NewMockQuery(ctrl)

	t.Run("Save user", func(t *testing.T) {
		user := newUser()

		session.EXPECT().DB().Return(database)
		database.EXPECT().C(repository.UsersCollection).Return(collection)
		collection.EXPECT().Insert(user)

		r := repository.NewUsersRepository(session)

		assert.NoError(t, r.Save(user))
	})

	t.Run("Get user by id", func(t *testing.T) {
		user := newUser()

		session.EXPECT().DB().Return(database)
		database.EXPECT().C(repository.UsersCollection).Return(collection)
		collection.EXPECT().FindId(user.Id).Return(query)
		query.EXPECT().One(gomock.Any()).SetArg(0, user).Return(nil)

		r := repository.NewUsersRepository(session)
		got, err := r.GetById(user.Id.Hex())
		assert.NoError(t, err)
		assert.Equal(t, user, got)
	})

	t.Run("Get user by email", func(t *testing.T) {
		user := newUser()

		session.EXPECT().DB().Return(database)
		database.EXPECT().C(repository.UsersCollection).Return(collection)
		collection.EXPECT().Find(bson.M{"email": user.Email}).Return(query)
		query.EXPECT().One(gomock.Any()).SetArg(0, user).Return(nil)

		r := repository.NewUsersRepository(session)
		got, err := r.GetByEmail(user.Email)
		assert.NoError(t, err)
		assert.Equal(t, user, got)
	})

	t.Run("Delete user", func(t *testing.T) {
		user := newUser()

		session.EXPECT().DB().Return(database)
		database.EXPECT().C(repository.UsersCollection).Return(collection)
		collection.EXPECT().RemoveId(user.Id).Return(nil)

		r := repository.NewUsersRepository(session)
		err := r.Delete(user.Id.Hex())
		assert.NoError(t, err)
	})
}

func newUser() *models.User {
	id := bson.NewObjectId()
	return &models.User{
		Id:      id,
		Name:    "TEST",
		Email:   id.Hex() + "@email.test",
		Created: time.Now(),
		Updated: time.Now(),
	}
}
