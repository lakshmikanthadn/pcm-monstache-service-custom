package main
import (
    "github.com/rwynn/monstache/v6/monstachemap"
    "fmt"
    "context"
    "reflect"
	"time"
    "go.mongodb.org/mongo-driver/bson"
   
)
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
    doc := input.Document
    con := input.MongoClient
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    booksCollection := con.Database("demo").Collection("books")
     titlesCollection := con.Database("demo").Collection("titles") 
  


    if doc["type"]== "book"{
        //  var k map[string]interface{}
        //  k:= doc["identifiers"]
        // fmt.Println(k)
        // for key, element := range k {
        //     fmt.Println("Key:", key, "=>", "Element:", element)
        // }
        // fmt.Println(doc["_id"])

        // mdi, err := page.Metadata()
        // md, ok := mdi.(doc["identifiers"])
        fmt.Println(reflect.TypeOf(doc["identifiers"]))
        // for key, element := range doc["identifiers"]{
        //     if key == "titleId" {
        //         fmt.Println("key", key, "Element:", element)
        //     }
        // }dock
        // for key, value := range doc {
        //     fmt.Println("[", key, "] has items:")
        //     for k, v := range value.([]map[string]interface{}){
        //         if k == "titleId" {
        //             fmt.Println("\t-->", k, ":", v)
        //         }else{
        //             fmt.Println("anything")
        //         }
        //     }
    
        // }
    
        titlesData, err := titlesCollection.Find(ctx, bson.D{{"titleId", bson.D{{"$eq", doc["identifiers"].(map[string]interface{})["titleId"]}}}}) 
        fmt.Println(titlesData)
        fmt.Println(doc["identifiers"].(map[string]interface{})["titleId"])

        if err != nil {
          panic(err)
        }
        var titles []bson.M
        if err = titlesData.All(ctx, &titles); err != nil {
            panic(err)
        }
            //    fmt.Println(titles)
            //    fmt.Println(reflect.TypeOf(doc))
            //    fmt.Println(doc)
            
       doc["variants"]= titles
         fmt.Println(doc)

    }

    if doc["titleId"]!= nil {
        fmt.Println("here ")
     booksData, err := booksCollection.Find(ctx, bson.D{{"identifiers.titleId", bson.D{{"$eq", doc["titleId"]}}}}) 
        if err != nil {
         fmt.Println("ssup")
         panic(err)
        }
     var books []bson.M
       if err = booksData.All(ctx, &books); err != nil {
          fmt.Println("hiii")
          panic(err)
        }
        fmt.Println(books)
      for k, v := range books {
        v["variants"]=doc
		books[k] = v
	  }
        doc = books[0]
        fmt.Println(doc)
    }

    

  
    
    // unwindStage := bson.D{{"$unwind", bson.D{{"path", "$empno"}, {"preserveNullAndEmptyArrays", false}}}}
    // showLoadedCursor, err := addressCollection.Find(ctx, matchQuery)
    // if err != nil {
    //     panic(err)
    // }
 
    //  convert bson to struct

    //  var s Struct1

    //  bsonBytes, _ := bson.Marshal(address)
    //  bson.Unmarshal(bsonBytes, &s)
    //  fmt.Println(bsonBytes)
    //  fmt.Println(s)

    
 
    // json.Unmarshal([]byte(address), &x)
    //  fmt.Println(titles[0])
    //  fmt.Println(books[0])
   
   
 



    // // var jsonDocuments map[string]interface{}

    // var temporaryBytes []byte

    // for cursor.Next(context.Background()) {
    // err = cursor.Decode(&episodes)
    // temporaryBytes, err = bson.MarshalExtJSON(episodes, true, true)
    //  err = json.Unmarshal(temporaryBytes, &doc)
    // jsonDocuments = append(jsonDocuments, doc)
    // }
    // for k, v := range doc {
    //     switch v.(type) {
    //     case string:
    //         doc[k] = strings.ToUpper(v.(string))
    //     }
    // }
    // fmt.Println(doc)
    // for k, v := range emp[0] {
	// 	doc[k] = v
	// }
    

    output = &monstachemap.MapperPluginOutput{Document: doc }
    return
}
