package persistent

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net"
	"time"
	"user-service/common"
	"user-service/entities"
)

// Database shares global database instance
var (
	Database MongoDB
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbSession  *mgo.Session
	Databasename string
}

// Init initializes mongo database
func (db *MongoDB) Init() error {
	log.Info("Init Database Resource")
	db.Databasename = common.K8sConfig.Out.MgDbName
	//DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:   []string{common.K8sConfig.Out.MgAddrs}, // Get HOST + PORT
		Timeout: 60 * time.Second,
		//Database: common.K8sConfig.Out.MgDbName,                   // Database name
		Username: common.K8sConfig.Out.MgDbUsername, // Username
		Password: common.K8sConfig.Out.MgDbPassword, // Password
	}

	//databaseName := common.K8sConfig.Out.MgDbName
	//dialInfo, err := mgo.ParseURL(mongoURI)
	//Below part is similar to above.
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		fmt.Println("DialServer error", err)
		return conn, err
	}

	// Create a session which maintains a pool of socket connections
	// to the DB MongoDB database.
	var err error
	db.MgDbSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Error("Can't connect to mongo, go error", err)
		return err
	}
	log.Error("success DataSource init")
	return db.initData()
}

// InitData initializes default data
func (db *MongoDB) initData() error {
	var err error
	var count int
	// Check if user collection has at least one document
	sessionCopy := db.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Databasename).C(common.ColUsers)
	count, err = collection.Find(bson.M{}).Count()

	if count < 1 {
		// Create admin/admin account
		var user entities.User = entities.User{bson.NewObjectId(), "admin", "admin"}
		err = collection.Insert(&user)
	}

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Close()
	}
}
