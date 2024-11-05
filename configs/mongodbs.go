package configs

// var mongoClient *mongo.Client

// func ConnectMongo(ctx context.Context, DBCollection ...string) *mongo.Database {
// 	connection := fmt.Sprintf("%s:%s", MONGOHost, MONGOPort)
// 	fmt.Println("Connection Mongo:", connection)
// 	clientOptions := options.Client()
// 	clientOptions.ApplyURI(connection)
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {

// 		return nil
// 	}

// 	err = client.Connect(ctx)
// 	if err != nil {

// 		return nil
// 	}

// 	var col string
// 	if AppEnv == "DEV" {
// 		col = MONGODBDEV
// 	}

// 	if len(DBCollection) > EMPTY_VALUE_INT {
// 		col = DBCollection[EMPTY_VALUE_INT]
// 	}

// 	mongoClient = client

// 	return client.Database(col)
// }

// func CloseMongo(ctx context.Context) {
// 	mongoClient.Disconnect(ctx)
// }
