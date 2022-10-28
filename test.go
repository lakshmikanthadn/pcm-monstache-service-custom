package main
import (
    "github.com/rwynn/monstache/v6/monstachemap"
    "fmt"
	  "time"
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "reflect" 
    "github.com/olivere/elastic/v7"
)
func Process(input*monstachemap.ProcessPluginInput) (err error){
    fmt.Println("hellpnesstess")
    doc := input.Document
    con := input.MongoClient
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    client :=  input.ElasticClient



    booksCollection := con.Database("demo").Collection("books")
    // Index name for the bulk Elasticsearch documents goes here
    indexName := "demo.books-4"
  
  // Declare a string slice with the index name in it
  // indices := []string{indexName}
  bulk := client.Bulk()
    if doc["titleId"]!= nil {
     
                fmt.Println("here ")
                delete(doc,"_id")
                fmt.Println(reflect.TypeOf(doc))
              //  opts := options.Find().SetProjection(bson.D{{"_id", 0}})
             booksData, err := booksCollection.Find(ctx, bson.D{{"identifiers.titleId", bson.D{{"$eq", doc["titleId"]}}}})
                if err != nil {
                 fmt.Println("ssup")
                 panic(err)
                }
                fmt.Println(reflect.TypeOf(booksData))
                  // errs := booksCollection.removed(bson.M{"_id": &id})
                  //   if errs != nil {
                  //       fmt.Println("some")
                  //   }
               var books []bson.M
               if err = booksData.All(ctx, &books); err != nil {
                  fmt.Println("hiii")
                  panic(err)
                }
                // delete(books[0],"_id")
              //  fmt.Println(reflect.TypeOf(books))
              //  var fruitBasket map[string]interface{}
              //   errs := json.Unmarshal(books, &fruitBasket)
           
              //  if errs != nil {
              //      fmt.Println("JSON decode error!")
              //  }
                // fmt.Println(books)
                // books[variants]=doc
                fmt.Println(books[0]["_id"])
              for k, v := range books {
               
                v["variants"]=doc
        		    books[k] = v
        	     }
              //  temp := books.(primitive.D)
              //  metadata := temp.Map() // map to map[string]interface{}
              //  delete(metadata, "_id")

              //     if v, ok := metadata[prqKey]; ok { // check and use value
              //         commitID = v.(string)
              //     }
              //  delete(books,"_id")
                //  doc = books[0]
                //  delete(doc,"_id")

                 // Iterate over the slice of Elasticsearch documents
                //  var t int = len(books)
                //   fmt.Println(t)
                //   var arr = make([]string, t)
                //   for i := 0; i < len(books); i++ {
                //     var a = books[i]["_id"]
                //     fmt.Println(a)
                //     fmt.Println(reflect.TypeOf(a))
                //     j := a.(string)
                //     arr[i] = j
                   
                //     delete(books[i], "_id")
                //     // fmt.Println(result)
                //   }
	              //     fmt.Println(arr)
                  //  docID1:= 12
                  // fmt.Println(docID1)
                  // delete(books[0],"_id")
  for _, bookdoc := range books {
  
    // Incrementally change the _id number in each iteration
   
    // var a = bookdoc["_id"]
    docID := bookdoc["_id"]

    delete(bookdoc, "_id")
     i := docID.(string) 
    
    // Convert the _id integer into a string
    // idStr := strconv.Itoa(docID)
    
    // Create a new int64 float from time package for doc timestamp
    // doc.Timestamp = time.Now().Unix()
    // fmt.Println("ntime.Now().Unix():", doc.Timestamp)
    
    // Declate a new NewBulkIndexRequest() instance
    req := elastic.NewBulkIndexRequest()
    
    // Assign custom values to the NewBulkIndexRequest() based on the Elasticsearch
    // index and the request type
    req.OpType("index") // set type to "index" document
    req.Index(indexName)
    //req.Type("_doc") // Doc types are deprecated (default now _doc)
    req.Id(i)
    req.Doc(bookdoc)
    
    // Print information about the NewBulkIndexRequest object
    fmt.Println("req:", req)
    fmt.Println("req TYPE:", reflect.TypeOf(req))
    
    // Add the new NewBulkIndexRequest() to the client.Bulk() instance
    bulk = bulk.Add(req)
    fmt.Println("NewBulkIndexRequest().NumberOfActions():", bulk.NumberOfActions())
    }
  
    
    // Do() sends the bulk requests to Elasticsearch
    bulkResp, err := bulk.Do(ctx)
    
    // Check if the Do() method returned any errors
    if err != nil {
     fmt.Println(err)
    } else {
    
    // If there is no error then get the Elasticsearch API response
    indexed := bulkResp.Indexed()
    fmt.Println("nbulkResp.Indexed():", indexed)
    fmt.Println("bulkResp.Indexed() TYPE:", reflect.TypeOf(indexed))
    
    // Iterate over the bulkResp.Indexed() object returned from bulk.go
    t := reflect.TypeOf(indexed)
    fmt.Println("nt:", t)
    fmt.Println("NewBulkIndexRequest().NumberOfActions():", bulk.NumberOfActions())
    
    // Iterate over the document responses
    for i := 0; i < t.NumMethod(); i++ {
    method := t.Method(i)
    fmt.Println("nbulkResp.Indexed() METHOD NAME:", i, method.Name)
    fmt.Println("bulkResp.Indexed() method:", method)
    }
    
    // Return data on the documents indexed
    fmt.Println("nBulk response Index:", indexed)
    for _, info := range indexed {
    fmt.Println("nBulk response Index:", info)
    //fmt.Println("nBulk response Index:", info.Index)
    }
    }
        		// fmt.Println(books)
        		// fmt.Println("testing books and doc")
                
            //     fmt.Println(doc)
                // delete(doc,"_id")
              //   result, err := client.Update().
              //           Index("demo.books-1").
              //           Type("_doc").
              //           Id("id").
              //           Doc(doc). // change doc here if you want to include/exclude/modify fields
              //           DocAsUpsert(true).
              //           Do(ctx)

              // fmt.Println(result)
              // fmt.Println("test")
              // fmt.Println(err)
            }
	
    // titlesCollection := con.Database("demo").Collection("titles") 
    return
}
// func main() {

//   // Allow for custom formatting of log output
//   log.SetFlags(0)
  
//   // Use the Olivere package to get the Elasticsearch version number
//   fmt.Println("Version:", elastic.Version)
  
//   // Create a context object for the API calls
//   ctx := context.Background()
  
//   // Declare a client instance of the Olivere driver
//   client, err := elastic.NewClient(
//   elastic.SetSniff(true),
//   elastic.SetURL("http://localhost:9200"),
//   elastic.SetHealthcheckInterval(5*time.Second), // quit trying after 5 seconds
//   )
  
//   // Check and see if olivere's NewClient() method returned an error
//   if err != nil {
//   fmt.Println("elastic.NewClient() ERROR:", err)
//   log.Fatalf("quiting connection..")
//   } else {
//   // Print client information
//   fmt.Println("client:", client)
//   fmt.Println("client TYPE:", reflect.TypeOf(client))
//   }
  
//   // Index name for the bulk Elasticsearch documents goes here
//   indexName := "some_index"
  
//   // Declare a string slice with the index name in it
//   indices := []string{indexName}
  
//   // Instantiate a new *elastic.IndicesExistsService
//   existService := elastic.NewIndicesExistsService(client)
  
//   // Pass the slice with the index name to the IndicesExistsService.Index() method
//   existService.Index(indices)
  
//   // Have Do() return an API response by passing the Context object to the method call
//   exist, err := existService.Do(ctx)
  
//   // Check if the IndicesExistsService.Do() method returned any errors
//   if err != nil {
//   log.Fatalf("IndicesExistsService.Do() ERROR:", err)
  
//   } else if exist == false {
//   fmt.Println("nOh no! The index", indexName, "doesn't exist.")
//   fmt.Println("Create the index, and then run the Go script again")
  
//   /*
//   curl -XPUT 'http://localhost:9200/some_index' -H 'Content-Type: application/json' -d '
//   {
//   "settings" : {
//   "number_of_shards" : 3,
//   "number_of_replicas" : 2
//   }
//   }'
//   */
  
//   // Bulk index the Elasticsearch documents if index exists
//   } else if exist == true {
//   fmt.Println("Index name:", indexName, " exists!")
  
//   // Declare an empty slice for the Elasticsearch document struct objects
//   docs := []ElasticDocs{}
  
//   // Get the type of the 'docs' struct slice
//   fmt.Println("docs TYPE:", reflect.TypeOf(docs))
  
//   // New ElasticDocs struct instances
//   newDoc1 := ElasticDocs{SomeStr: "Hello, world!", SomeInt: 42, SomeBool: true, Timestamp: 0.0}
//   newDoc2 := ElasticDocs{SomeStr: "ä½ å¥½ï¼Œä¸–ç•Œï¼", SomeInt: 7654, SomeBool: false, Timestamp: 0.0}
//   newDoc3 := ElasticDocs{SomeStr: "Kumusta, mundo!", SomeInt: 1234, SomeBool: true, Timestamp: 0.0}
  
//   // Append the new Elasticsearch document struct objects to the slice
//   docs = append(docs, newDoc1)
//   docs = append(docs, newDoc2)
//   docs = append(docs, newDoc3)
  
//   // Declare a new Bulk() object using the client instance
//   bulk := client.Bulk()
  
//   // Elasticsearch _id counter starts at 0
//   docID := 0
  
//   // Iterate over the slice of Elasticsearch documents
//   for _, doc := range docs {
  
//   // Incrementally change the _id number in each iteration
//   docID++
  
//   // Convert the _id integer into a string
//   idStr := strconv.Itoa(docID)
  
//   // Create a new int64 float from time package for doc timestamp
//   doc.Timestamp = time.Now().Unix()
//   fmt.Println("ntime.Now().Unix():", doc.Timestamp)
  
//   // Declate a new NewBulkIndexRequest() instance
//   req := elastic.NewBulkIndexRequest()
  
//   // Assign custom values to the NewBulkIndexRequest() based on the Elasticsearch
//   // index and the request type
//   req.OpType("index") // set type to "index" document
//   req.Index(indexName)
//   //req.Type("_doc") // Doc types are deprecated (default now _doc)
//   req.Id(idStr)
//   req.Doc(doc)
  
//   // Print information about the NewBulkIndexRequest object
//   fmt.Println("req:", req)
//   fmt.Println("req TYPE:", reflect.TypeOf(req))
  
//   // Add the new NewBulkIndexRequest() to the client.Bulk() instance
//   bulk = bulk.Add(req)
//   fmt.Println("NewBulkIndexRequest().NumberOfActions():", bulk.NumberOfActions())
//   }
  
//   // Do() sends the bulk requests to Elasticsearch
//   bulkResp, err := bulk.Do(ctx)
  
//   // Check if the Do() method returned any errors
//   if err != nil {
//   log.Fatalf("bulk.Do(ctx) ERROR:", err)
//   } else {
  
//   // If there is no error then get the Elasticsearch API response
//   indexed := bulkResp.Indexed()
//   fmt.Println("nbulkResp.Indexed():", indexed)
//   fmt.Println("bulkResp.Indexed() TYPE:", reflect.TypeOf(indexed))
  
//   // Iterate over the bulkResp.Indexed() object returned from bulk.go
//   t := reflect.TypeOf(indexed)
//   fmt.Println("nt:", t)
//   fmt.Println("NewBulkIndexRequest().NumberOfActions():", bulk.NumberOfActions())
  
//   // Iterate over the document responses
//   for i := 0; i < t.NumMethod(); i++ {
//   method := t.Method(i)
//   fmt.Println("nbulkResp.Indexed() METHOD NAME:", i, method.Name)
//   fmt.Println("bulkResp.Indexed() method:", method)
//   }
  
//   // Return data on the documents indexed
//   fmt.Println("nBulk response Index:", indexed)
//   for _, info := range indexed {
//   fmt.Println("nBulk response Index:", info)
//   //fmt.Println("nBulk response Index:", info.Index)
//   }
//   }
//   }
//   }
// func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
//     doc := input.Document
//     con := input.MongoClient
//     ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	input= &monstachemap.MapperPluginInput{Operation: "u"}

//     booksCollection := con.Database("demo").Collection("books")
	

// //     // titlesCollection := con.Database("demo").Collection("titles") 
  


// //     // if doc["type"]== "book"{
// //     //     //  var k map[string]interface{}
// //     //     //  k:= doc["identifiers"]
// //     //     // fmt.Println(k)
// //     //     // for key, element := range k {
// //     //     //     fmt.Println("Key:", key, "=>", "Element:", element)
// //     //     // }
// //     //     // fmt.Println(doc["_id"])

// //     //     // mdi, err := page.Metadata()
// //     //     // // md, ok := mdi.(doc["identifiers"])
// //     //     // fmt.Println(reflect.TypeOf(doc["identifiers"]))
// //     //     // for key, element := range doc["identifiers"]{
// //     //     //     if key == "titleId" {
// //     //     //         fmt.Println("key", key, "Element:", element)
// //     //     //     }
// //     //     // }dock
// //     //     // for key, value := range doc {
// //     //     //     fmt.Println("[", key, "] has items:")
// //     //     //     for k, v := range value.([]map[string]interface{}){
// //     //     //         if k == "titleId" {
// //     //     //             fmt.Println("\t-->", k, ":", v)
// //     //     //         }else{
// //     //     //             fmt.Println("anything")
// //     //     //         }
// //     //     //     }
    
// //     //     // }

// //     //     titlesData, err := titlesCollection.Find(ctx, bson.D{{"titleId", bson.D{{"$eq", doc["identifiers"].(map[string]interface{})["titleId"]}}}}) 
// //     //     // fmt.Println(titlesData)
// //     //     fmt.Println(doc["identifiers"].(map[string]interface{})["titleId"])

// //     //     if err != nil {
// //     //       panic(err)
// //     //     } 
// //     //     var titles []bson.M
// //     //     if err = titlesData.All(ctx, &titles); err != nil {
// //     //         panic(err)
// //     //     }
// //     //         //    fmt.Println(titles)
// //     //         //    fmt.Println(reflect.TypeOf(doc))
// //     //         //    fmt.Println(doc)
            
// //     //    doc["variants"]= titles
// //     //      fmt.Println(doc)

// //     // }

//     if doc["titleId"]!= nil {
//         fmt.Println("here ")
//         delete(doc,"_id")
//      booksData, err := booksCollection.Find(ctx, bson.D{{"identifiers.titleId", bson.D{{"$eq", doc["titleId"]}}}}) 
//         if err != nil {
//          fmt.Println("ssup")
//          panic(err)
//         }
//        var books []bson.M
//        if err = booksData.All(ctx, &books); err != nil {
//           fmt.Println("hiii")
//           panic(err)
//         }
//         // fmt.Println(books)
//       for k, v := range books {
//         v["variants"]=doc
// 		    books[k] = v
// 	     }
//         doc = books[0]
// 		fmt.Println(books)
// 		fmt.Println("testing books and doc")
        
//         fmt.Println(doc)
//     }
// //       //  var final map[string]interface{}
// //       //  result, err := booksCollection.InsertOne(context.TODO(), doc)
// //       //  final = result
      
    

  
    
// //     // unwindStage := bson.D{{"$unwind", bson.D{{"path", "$empno"}, {"preserveNullAndEmptyArrays", false}}}}
// //     // showLoadedCursor, err := addressCollection.Find(ctx, matchQuery)
// //     // if err != nil {
// //     //     panic(err)
// //     // }
 
// //     //  convert bson to struct

// //     //  var s Struct1

// //     //  bsonBytes, _ := bson.Marshal(address)
// //     //  bson.Unmarshal(bsonBytes, &s)
// //     //  fmt.Println(bsonBytes)
// //     //  fmt.Println(s)

    
 
// //     // json.Unmarshal([]byte(address), &x)
// //     //  fmt.Println(titles[0])
// //     //  fmt.Println(books[0])
   
   
 



// //     // // var jsonDocuments map[string]interface{}

// //     // var temporaryBytes []byte

// //     // for cursor.Next(context.Background()) {
// //     // err = cursor.Decode(&episodes)
// //     // temporaryBytes, err = bson.MarshalExtJSON(episodes, true, true)
// //     //  err = json.Unmarshal(temporaryBytes, &doc)
// //     // jsonDocuments = append(jsonDocuments, doc)
// //     // }
// //     // for k, v := range doc {
// //     //     switch v.(type) {
// //     //     case string:
// //     //         doc[k] = strings.ToUpper(v.(string))
// //     //     }
// //     // }
// //     // fmt.Println(doc)
// //     // for k, v := range emp[0] {
// // 	// 	doc[k] = v
// // 	// }
// // //    result, err := booksCollection.Find(bson.D{{"identifiers.titleId", bson.D{{"$eq", doc["titleId"]}}}}).Select(bson.D{"identifiers.titleId": 1}).All(&result)
// // //     fmt.Println(result)
// //     str:= "demo.books"
//   output = &monstachemap.MapperPluginOutput{Document: doc , Index: "demo.books" , ID: "something" }
//     return
// }


