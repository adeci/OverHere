Frontend:

Cypress Test:
it('Reroutes on button press', () => {})
-basic e2e cypress test that find the 'Sign Up' button on the home page/login 
component and clicks it, and checks the url to ensure that it correctly 
reroutes to localhost:4200/new-user (the new user/signup screen).

**Unit tests (jasmine spec files):**  
**NewuserPagesComponent**  
- should link to login on back click
- should display submit and back buttons
- should link to afterlogin on confirm click
- should create newuser
- should disp new title  
**PhotoLibPagesComponent**  
- should link to afterlogin on back click
- should disp photolib title
- should display nextimg and back buttons
- should create photolib  
**LoginpagesComponent**  
- should link to login on signupbutton click
- should display buttons
- should link to login on loginbutton click
- should disp title
- should create  
**AfterloginPagesComponent**  
- should link to map on mapbutton click
- should create afterlogin
- should disp title
- should link to photoup on photoup click
- should link to home on signout click
- should link to lib on library click  
**ReturninguserPagesComponent**  
- should link to login on back click
- should display submit and back buttons
- should create returninguser
- should link to afterlogin on confirm click
- should disp returning title  
**PhotoupPagesComponent**  
- should link to homepage on back click
- should create photoup  

Each of these unit tests are created with the intent of checking the  
correctness of the displayed items for each page, as well as testing the   
routing functions that occur on button presses to ensure that they send the   
user to the correct page.  

**Work completed for sprint 2:**  
For this sprint, we created several pages and routes to help portray   
the intended user experience on the app. There is a home page, labeled login   
component, a sign up/new user screen, a login/returning user screen, a map screen, and   
photo upload and view pages which we plan to further expand and add functionality to   
in the future. We also integrated a user variable, which is passed through the different   
components after sign up and log in to keep track of who is using the app.   
Additionally, we created several unit tests to ensure that the app displays what it is   
supposed to on each page. We also extensively tested the routing functions that occur on  
button presses.

Backend:

Nico (database):
- Migrated to MongoDB Atlas (cloud server database) from locally ran MongoDB client.
- Created post functions for creating User, OHPost, and Image objects and storing them in DB.
- Created get function for User objects
- Tested database structure workflow designed by Alex G.
- Wrote tests for database and imageprocessing files.

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

GetUserObject(userid string) UserObject:
- Take in username string and returns User Object containing that username stored in MongoDB Atlas.

Tests (database):

TestCreateAndStoreUserObject(t *testing.T):
- Tests the creation and storage of a User Object in MongoDB Atlas.

TestCreateAndStoreOHPostObject(t *testing.T):
- Tests the creation and storage of an OHPost Object in MongoDB Atlas.

TestCreateAndStoreImageObject(t *testing.T):
- Tests the creation and storage of an Image Object in MongoDB Atlas.

TestGetUserObject(t *testing.T):
- Tests get function for a User Object.

Tests (imageprocessing):

TestToBase64(t *testing.T):
- Tests conversion of bytes to base64 string.

TestDecodePNG(t *testing.T):
- Tests conversion of local image to base64 string.

TestEncodePNG(t *testing.T):
- Tests conversion of base64 string to locally written image.



Alex (routing):
- Designed data formatting between User, OHPost, and Image
- Sorted backend into controller, model, responses, routes, and services files
- Added User, Image, and OHPost models - format to receive data from frontend
- Added Create User, Get User, Create Image, and Get Image routes & controllers
- Added Postman mock frontend requests CreateImage_Dummy, GetImage_Dummy, CreateUser_Dummy, GetUser_Dummy
- Added Routing Tests based on Postman requests
- Added Makefile to run frontend and backend easily
- Added documentation below

**Routing Tests**
TestCreateUserRoute
TestGetUserRoute
TestCreateImageRoute
TestGetImageRoutes

**Routing documentation**:

The backend applications acts as a server. 
Currently runs on localhost:8000 (port 8000)
Run backend using "make runback" command. It'll then be listening for HTTP messages.
Outlined below are messages you can send the backend using.

Format:
Name - Colloquial name of request, for humans to understand
URI - HTTP knows what request to handle from URI. /:value is parsed as *actual data* from this URI.
Body - HTTP messages have bodys. Backend parses the body. Must match this format in a string like "{"key":value}"
Response - Backend will send an HTTP message back with a Body. Use this data to continue frontend activities.


**Create User:**
Creates a user in the database with a userid, username, and more. 

**URI** - POST hostname/users/create
    Ex: http.message(POST, localhost:8000/users/create

**Body** - 
{   
    "username": "String" 
}

Username (String) - Required
    Name of the new user. Not unique.
    
**Response** - 
Sample response
Sample
{
    "status": 201,       
    "message": "success",
    "data": {
        "data": {
            "userid": "string",
            "username": "String"
        }
    }
}

Status (int):
    StatusCreated - 201
        Successfully created user.
    StatusBadRequest - 400 
        Could not parse request Body. Missing required fields
    
Message (string):
    success
        Successfully created user.
    error
        Did not retrieve user

Data (map[string]interface{})
    Contains data of user created
    Map corresponding to user data

    UserID (string) -
        Name of created users ID, assigned by backend. Is unique.
    Username (string) - 
        Name of created user, based on name from Body


**Get User:**
Retrieves user data from database

**URI** - GET hostname/users/get/:userid
    Ex: http.message(GET, localhost:8000/users/get/123456)

**Body** - No body for GET request.
    Ex. nil or null

**Response** - 
Sample response
{
    "status": 200,
    "message": "success",
    "data": {
        "data": {
            "userid": "string",
            "username": "string"
        }
    }
}

Status (int):
    StatusOK - 200
        Successfully retrieved user
    StatusBadRequest - 400
        Could not parse request URI. Missing required fields (likely missing :userid)
        Did not retrieve user

Message (string):
    success
        Successfully retrieved user
    error
        Did not retrieve user

Data (map[string]interface{})
    Contains data of user created
    Map corresponding to user data

    UserID (string) -
        Name of retrieved users ID, assigned by backend
    Username (string) - 
        Name of retrieved user


**Create Image:**

**URI** - POST hostname/images/create
    Ex. http.message(POST, localhost:8000/images/create)

**Body** - 
{
    "imageid": "string",
    "userid": "string",
    "ohpostid": "string",
    "encoding": "string"
}

ImageID (string) - 
    ID of created image
    *will be created by database in the future*

UserID (string) -
    User the created image will belong to

OHPostID (string) -
    OHPost the created image will belong to

Encoding (string) - 
    Image encoding that represents the image
    Backend stores in base64
    Will convert frontend's encoding.

**Response** - 
Sample response
{
    "status": 201,
    "message": "success",
    "data": {
        "data": {
            "imageid": "string",
            "userid": "string",
            "ohpostid": "string",
            "encoding": "string"
        }
    }
}

Status (int):
    StatusOK - 200
        Successfully retrieved user
    StatusBadRequest - 400
        Could not parse request Body. Missing required fields
        Did not retrieve user

Message (string):
    success
        Successfully created image.
    error
        Did not create image

Data (map[string]interface{})
    Contains data of image created
    Map corresponding to image data

    ImageId (string) -
        Database created ID associated with the image. Use to retrieve.

    UserID (string) -
        User the created image belongs to

    OHPostID (string) -
        OHPost the created image belongs to

    Encoding (string)
        Image encoding that represents the image
        Backend stores in base64
        Converts to frontend encoding (currently responds with what was sent)


**Get Image:**

**URI** - GET hostname/images/get/:imageid
    Ex. http.message(GET, localhost:8000/images/get/123456)

**Body** - No body for GET request.
    Ex. nil or null

**Response** - 
Sample response
{
    "status": 200,
    "message": "success",
    "data": {
        "data": {
            "imageid": "String",
            "userid": "String",
            "ohpostid": "String",
            "encoding": "String"
        }
    }
}

Status (int):
    StatusOK - 200
        Successfully retrieved image.
    StatusBadRequest - 400
        Could not parse request URI. Missing required fields.
        Did not retrieve image.

Message (string):
    success
        Successfully retrieved image.
    error
        Did not retrieve image

Data (map[string]interface{})
    Contains data of image created
    Map corresponding to image data

    ImageId (string) -
        Database created ID associated with the image. Just used to retrieve this image.

    UserID (string) -
        User the created image belongs to.

    OHPostID (string) -
        OHPost the created image belongs to.

    Encoding (string) - 
        Image encoding that represents the image.
        Backend stores in base64.
        Converts to frontend encoding (currently responds with what was sent).
