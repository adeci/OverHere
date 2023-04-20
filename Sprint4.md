Backend:

Priority   
- Optimize runtime for http functions
  - The http functions interacting directly with our online database were really slow. This is because each http 
  function contained the process of establishing a connection to the database and establishing a connection to a
  collection within the database. Successive connections lead to really slow testing and function of the web app. I 
  optimized it so that on runtime the app connects to the database and its collections once. This improved test runtime
  by over 300%.
  
- Change id standard
  - The old id standards were unreadable and didn't help in the debugging process. Changing the standard for userids 
  to "USER-" + blah blah blah and the same with ohpostids and imageids helps with readability.

- Implemented tags
  - Tags have been implemented for OHPosts. This introduces a new search key that can be used to display OHPosts on
  the public map.

Low Priority
- Host web app on a domain (overhere.tech)

Public Database API Documention:

    200 - success
    201 - created
    400 - bad request
    500 - internal server error

    Connect to Database:

    ConnectMongoDBAtlas() *mongo.Client:
      Connects app to online database. Runs once at runtime or once at the beginning of each test.
        Returns *mongo.Client (you don't need this).

    ID Validation:

    ValidateUserID(userid string) error:
      Validates userid.
        Returns error if userid is not valid.

    ValidateOHPostID(ohpostid string) error:
      Validates ohpostid.
        Returns error if ohpostid is not valid.

    ValidateImageID(imageid string) error:
      Validates imageid.
        Returns error if imageid is not valid.
    
    Post:
    
    PostUser(username string) (UserObject, error):
        Creates and stores User. UserID Generated.
            Returns User created.

        URI: Post [Hostname]/users/post
      
        Json Body:
            "username": String

    PostOHPost(userid string, description string, xcoord float64, ycoord float64, tag string) (OHPostObject, error):
        Creates and stores OHPost. OHPostID generated ("OHPOST-" + generated ohpostid).
            Returns OHPost created.

    PostImage(base64encode string, userid string, ohpostid string, xcoord float64, ycoord float64) (ImageObject, error):
        Creates and stores Image. ImageID generated ("IMAGE-" + generated imageid).
            Returns Image created.

        URI: Post [Hostname]/images/post

        Json Body:
            "userid": String,
            "encoding": String,
            "xcoord": float64,
            "ycoord": float64,
            "caption": String,
            "tag": String

    Put:

    PutUser(userid string, username string) error:
        Updates username for User. userid is search key, username is new updated username to put.
        
        URI: Put [Hostname]/users/put

        Json Body:
            "username": String

    PutOHPost(object OHPostObject):
        Updates OHPost.

    PutOHPost_XCoord(ohpostid string, xcoord float64) error:
        Updates xcoord for OHPost. ohpostid is search key, xcoord is new updated xcoord to put.
    
    PutOHPost_YCoord(ohpostid string, ycoord float64) error:
        Updates ycoord for OHPost. ohpostid is search key, ycoord is new updated ycoord to put.

    PutImage(object ImageObject):
        Updates Image.

        URI: Put [Hostname]/images/put/[imageid]

        Json Body:
            "imageid": String,
            "userid": String,
            "ohpostid": String,
            "encoding": String,
            "xcoord": Float64,
            "ycoord": Float64,
            "tag": String,
            "caption": String

    PutImage_OHPostID(imageid string, ohpostid string):
        Updates ohpostid for Image. imageid is search key, ohpostid is new updated ohpostid to put.
    
    PutImage_XCoord(imageid string, xcoord float32):
        Updates xcoord for Image. imageid is search key, xcoord is new updated xcoord to put.

    PutImage_YCoord(imageid string, ycoord float32):
        Updates ycoord for Image. imageid is search key, ycoord is new updated xcoord to put.

    Get:

    GetUser_UserID(userid string) (UserObject, error):
        Gets User. userid is search key.
            Returns User.

        URI: Get [Hostname]/users/get/[userid]

        Json Body:
            "userid": String

    GetUser_Username(username string) (UserObject, error):
        Gets User. username is search key.
            Returns User.
      
        URI: Get [Hostname]/users/get/byusername/[username]

        Json Body:
            "username": String

    GetUser_All() ([]UserObject, error):
      Gets User/s. Every user is returned.
            Returns array of Users.

    GetOHPost_OHPostID(ohpostid string) (OHPostObject, error):
        Gets OHPost. ohpostid is search key.
            Returns OHPost.

    GetOHPost_UserID(userid string) ([]OHPostObject, error):
        Gets OHPost/s. userid is search key.
            Returns array of OHPosts.

    GetOHPost_Tag(tag string) ([]OHPostObject, error):
        Gets OHPost/s. tag is search key.
            Returns array of OHPosts.

    GetImage_ImageID(imageid string) (ImageObject, error):
        Gets Image. imageid is search key.
            Returns Image.
        
        URI: Get [Hostname]/images/get/[imageid]

        Json Body:
            "imageid": String

    GetImage_UserID(userid string) ([]ImageObject, error):
        Gets Image/s. userid is search key.
            Returns array of Images.

        URI: Get [Hostname]/images/get/byuserid/[userid]

        Json Body:
            "userid": String

    GetImage_OHPostID(ohpostid string) ([]ImageObject, error):
        Gets Image/s. ohpostid is search key.
            Returns array of Images.

    GetImage_All() ([]ImageObject, error):
        Gets Image/s. Every image is returned.
            Returns array of Images.

    Delete:

    DeleteUser_UserID(userid string) error:
        Deletes User. userid is search key.

        URI: Delete [Hostname]/users/delete/[userid]

        Json Body:
            "userid": String

    DeleteUser_Username(username string) error:
        Deletes User. username is search key.

        URI: Delete [Hostname]/users/delete/byusername[username]

        Json Body:
            "username": String

    DeleteOHPost_OHPostID(ohpostid string) error:
        Deletes OHPost. ohpostid is search key.

    DeleteOHPost_UserID(userid string) error:
        Deletes OHPost/s. userid is search key.

    DeleteImage_ImageID(imageid string) error:
        Deletes Image. imageid is search key.
      
        URI: Delete [Hostname]/images/delete/[imageid]

        Json Body:
            "imageid": String

    DeleteImage_UserID(userid string) error:
        Deletes Image/s. userid is search key.

        URI: Delete [Hostname]/images/delete/byuserid/[userid]

        Json Body:
            "userid": String

    DeleteImage_OHPostID(ohpostid string) error:
        Deletes Image/s. ohpostid is search key.



Database Function Tests:

    TestPostUser
    TestPostOHPost
    TestPostImage
    TestPutUser_Username
    TestGetUser_Username
    TestGetImage_ImageID
    TestDeleteUser_UserID
    TestDeleteOHPost_OHPostID
    TestDeleteOHPost_UserID
    TestDeleteImage_ImageID
    TestDeleteImage_UserID
    TestDeleteImage_OHPostID
    TestGetOHPost_UserID