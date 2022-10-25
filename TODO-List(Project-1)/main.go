package main


var ren *renderer.Render
var db *mgo.Database

const(
	hostName string = "localhost:27017"
	dbName string = "todo-list"
	collectionName string = "todo"
	port string = ":8080"
)

type (
	todoModel struct {
		ID bson.ObjectId `bson:"_id,omitempty"`
		Title string `bson:"title"`
		Completed bool `bson:"completed"`
		CreatedAt time.Time `bson:createAt`
	}

	todo struct {
		ID string = `json:"id"`
		Title string = `json:"title"`
		Completed bool = `json:"completed"`
		CreatedAt time.Time = `json:"created_at"`
	}
)

func init(){
	ren = renderer.New()
	sess, err = mgo.Dial(hostName)
	checkErr
}