package database

var ConnectionString = "mongodb+srv://elite5000:Rp4PQ27PeiDoZ0MF@cluster0.3qzpcw3.mongodb.net/"

type DB struct {
	client *mongo.Client
}
