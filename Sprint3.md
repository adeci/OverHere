Backend:

Priority   
- GetOHPostsWithinCoordinates(4 corners)
- This function will take in 4 corners of the visible map and use it as borders to sort through
posts within this border. Return OHPosts, return 1 image for thumbnail.

    - GetImagesWithinOHPost(OHPostID)
        - Return images associated with OHPost.

    - Generalize Get, Post, Put, Delete for Image, OHPost, User
    - Generate unique UserIDs
    - Generate unique OHPostIDs

Low Priority
- Support JPEG

Notes
- An OHPost has 1 coordinate (average location of all images within ohpost)
- An image has 1 coordinate (where it taken)

HTTP Method Documention:
    
    Post:
    
    PostUser(username string) UserObject:
        Creates and stores User. UserID Generated.
            Returns User created.

    PostOHPost(userid string, description string, xcoord float32, ycoord float32) OHPostObject:
        Creates and stores OHPost. OHPostID generated (it's userid + "/" + generated id).
            Returns OHPost created.

    PostImage(base64encode string, userid string, ohpostid string, xcoord float32, ycoord float32) ImageObject:
        Creates and stores Image. ImageID generated (it's ohpostid + "/" + generated id).
            Returns Image created.

    Put:

    PutUser_Username(userid string, username string):
        Updates username for User. userid is search key, username is new updated username to put.

    PutOHPost_Description(ohpostid string, description string):
        Updates description for OHPost. ohpostid is search key, description is new updated description to put.

    PutOHPost_XCoord(ohpostid string, xcoord float32):
        Updates xcoord for OHPost. ohpostid is search key, xcoord is new updated xcoord to put.
    
    PutOHPost_YCoord(ohpostid string, ycoord float32):
        Updates ycoord for OHPost. ohpostid is search key, ycoord is new updated ycoord to put.
    
    PutImage_OHPostID(imageid string, ohpostid string):
        Updates ohpostid for Image. imageid is search key, ohpostid is new updated ohpostid to put.
    
    PutImage_XCoord(imageid string, xcoord float32):
        Updates xcoord for Image. imageid is search key, xcoord is new updated xcoord to put.

    PutImage_YCoord(imageid string, ycoord float32):
        Updates ycoord for Image. imageid is search key, ycoord is new updated xcoord to put.

    Get:

    GetUser_UserID(userid string) UserObject:
        Gets User. userid is search key.
            Returns User.

    GetUser_Username(username string) UserObject:
        Gets User. username is search key.
            Returns User.

    GetOHPost_OHPostID(ohpostid string) OHPostObject:
        Gets OHPost. ohpostid is search key.
            Returns OHPost.

    GetOHPost_UserID(userid string) []OHPostObject:
        Gets OHPost/s. userid is search key.
            Returns array of OHPosts.

    GetImage_ImageID(imageid string) ImageObject:
        Gets Image. imageid is search key.
            Returns Image.

    GetImage_UserID(userid string) []ImageObject:
        Gets Image/s. userid is search key.
            Returns array of Images.

    GetImage_OHPostID(ohpostid string) []ImageObject:
        Gets Image/s. ohpostid is search key.
            Returns array of Images.

    Delete:

    DeleteUser_UserID(userid string):
        Deletes User. userid is search key.

    DeleteUser_Username(username string):
        Deletes User. username is search key.

    DeleteOHPost_OHPostID(ohpostid string):
        Deletes OHPost. ohpostid is search key.

    DeleteOHPost_UserID(userid string):
        Deletes OHPost/s. userid is search key.

    DeleteImage_ImageID(imageid string):
        Deletes Image. imageid is search key.

    DeleteImage_UserID(userid string):
        Deletes Image/s. userid is search key.

    DeleteImage_OHPostID(ohpostid string):
        Deletes Image/s. ohpostid is search key.
