package dao

import (
	"gopkg.in/mgo.v2/bson"
	"user-service/common"
	"user-service/entities"
	"user-service/persistent"
	"user-service/utils"
)

// User manages User CRUD
type User struct {
	utils *utils.Jwt
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]entities.User, error) {
	sessionCopy := persistent.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(persistent.Database.Databasename).C(common.ColUsers)

	var users []entities.User
	err := collection.Find(bson.M{}).All(&users)
	return users, err
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (entities.User, error) {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return entities.User{}, err
	}

	sessionCopy := persistent.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(persistent.Database.Databasename).C(common.ColUsers)

	var user entities.User
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	sessionCopy := persistent.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(persistent.Database.Databasename).C(common.ColUsers)

	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Login User
func (u *User) Login(name string, password string) (entities.User, error) {
	sessionCopy := persistent.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(persistent.Database.Databasename).C(common.ColUsers)

	var user entities.User
	err := collection.Find(bson.M{"$and": []bson.M{bson.M{"name": name}, bson.M{"password": password}}}).One(&user)
	return user, err
}

// Insert adds a new User into database'
func (u *User) Insert(user entities.User) error {
	sessionCopy := persistent.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(persistent.Database.Databasename).C(common.ColUsers)

	err := collection.Insert(&user)
	return err
}

// Delete remove an existing User
func (u *User) Delete(user entities.User) error {
	sessionCopy := persistent.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(persistent.Database.Databasename).C(common.ColUsers)

	err := collection.Remove(&user)
	return err
}

// Update modifies an existing User
func (u *User) Update(user entities.User) error {
	sessionCopy := persistent.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(persistent.Database.Databasename).C(common.ColUsers)

	err := collection.UpdateId(user.ID, &user)
	return err
}
