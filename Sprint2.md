





Backend:
Alex (routing):
- Sorted backend into controller, model, responses, routes, and services files
- Added User, Image, and OHPost models - format to receive data from frontend
- Added Create User, Get User, Create Image, and Get Image routes & controllers
- Added Makefile to run frontend and backend easily

Routing documentation:

Create User:
Creates a user in the database with a userid, username, and more. 

URI - POST hostname/users/create
    Sample: POST localhost:8000/users/create

Body - 
{
    "_id": DatabaseID     
    "userid": "String",   
    "username": "String" 
}

_id (primitive.ObjectID)
    MongoDB created id.
    Don't change
    Ex. {"$oid":"63fd804978e7971b8b27e106"}
userid (String) - Required
    Will eventually be created by backend
username (String) - Required

Response - 
{
    "status": int,       
    "message": "success",
    "data": {
        "data": {
            "_id": "DatabaseID",
            "userid": "String",
            "username": "String"
        }
    }
}

Status (int):
    StatusCreated - 201
        Successfully created user.
    StatusBadRequest - 400 
        Could not parse request Body.
    
Message (string):
    success
        Successfully created user.
    error
        Could not parse request Body

Data (map[string]interface{})
    Contains data of user created
    Map corresponding to user data



Get User:

URI - GET hostname/users/get/:userid
    Sample: GET localhost:8000/users/get/123456

Response - 
{
    "status": 200,
    "message": "success",
    "data": {
        "data": {
            "_id": "63feddbf92b34f4119af0c82",
            "userid": "string",
            "username": "string"
        }
    }
}

Status (int):
    StatusOK - 200
        Successfully retrieved user

Message (string):
    success
        Successfully retrieved user
    error

Data

Create Image:

Body - 
{
    "_id": primitive.ObjectID
    "imageid": "string",
    "ohpostid": "string",
    "encoding": "string"
}

_id (primitive.ObjectID)
    MongoDB created id.
    Don't change
    Ex. {"$oid":"63fd804978e7971b8b27e106"}

ImageId (string) -

type Image struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	ImageID  string             `json:"imageid,omitempty"`
	OHPostID string             `json:"ohpostid,omitempty"`
	Encoding string             `json:"encoding" validate:"required"`
}

Database Documentation:

CreateAndStoreUserObject(username string) UserObject:
- Takes in username string and creates and stores User Object in MongoDB Atlas.
- Returns User Object created.
- Notes: Does NOT check if username exists.

CreateAndStoreOHPostObject(ohpostid string, userid string, description string) OHPostObject:
- Takes in parameters above and creates and stores OHPost Object in MongoDB Atlas.
- Returns OHPost Object created.
- Notes: Does NOT check if OHPost exists.

CreateAndStoreImageObject imageid string, base64encode string, userid string, ohpostid string) ImageObject:
- Takes in parameters above and creates and stores Image Object in MongoDB Atlas.
- Returns Image Object created.
- Notes: Does NOT check if image exists.