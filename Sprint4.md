Backend:

Nico:
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


    
Alex:
Priority
Add useful controllers for Frontend to practically use.
Added X and Y coordinates to images.
Updated controllers to have tag functionality.
Created Image coordinate averaging for 

New Controllers:

  -PostOHPostWithImageIds (Biggest addition)
  -PostOHPost 
  -GetOHPostsByCoordinateBoundary (Unused for now)
  -GetOHPostsByUserId 
  -GetImagesByOHPost
  -DeleteOHPostsByUserId


  -GetImagesByUserId
  -PutAddImageToOHPost - Unused since slightly broken
  -DeleteImagesByOHPost
  -DeleteImagesByUserId


  -GetUserByUsername
  -DeleteUserByUsername
  -GetUserByUsername

Frontend:

For this sprint, we wanted to ensure that our backend database was fully implemented into our program. We were able to do this by posting and getting images during runtime. When the user loads their map, they can see their own posts on the exact coordinates they marked them on. A user can also view all of the images they have posted, in a library that can be traversed. Lastly, username sign in functionality was improved upon. Users can now only sign in if their username already exists in the database. If it is not found, they are prompted to sign up. Lastly, some visual changes, most notably to the photo library page, were added.

Tests:
Unit tests (Jasmine)
LoginpagesComponent
should link to login on signupbutton click
should link to login on loginbutton click
should display buttons
should disp title
should create
AfterloginPagesComponent
should create afterlogin
should link to lib on library click
should link to map on mapbutton click
should link to home on signout click
should disp title
PhotoupPagesComponent
should link to homepage on back click
should create photoup
PhotoLibPagesComponent
should create photolib
should disp photolib title
NewuserPagesComponent
should link to afterlogin on confirm click
should display submit and back buttons
should link to login on back click
should create newuser
should disp new title
ReturninguserPagesComponent
should create returninguser
should display submit and back buttons
should link to afterlogin on confirm click
should disp returning title
should link to login on back click

This test checks an attempt to login with an incorrect preexisting username.
It tries to log in and shows the unavailable username and confirms returning to main page.

incorrectlogin.cy.s:
    
        describe('template spec', () => {
            it('runs', () => {
              cy.visit('localhost:4200')
            })

            it('user doesnt exist', () => {
                cy.visit('localhost:4200/returning-user')

                cy.get('input').type('badusername')

                cy.contains('Confirm').click()

              })
          })

This test creates a new user and opens the library to display the new library page.

 openlibrary.cy.js:
   
        describe('template spec', () => {
            it('runs', () => {
              cy.visit('localhost:4200')
            })

            it('opens new user library', () => {
                cy.visit('http://localhost:4200/new-user')

                cy.get('input').type('brandnewuser')

                cy.contains('Confirm').click()

                cy.contains('View Your Photo Library').click()
                cy.url().should('include', '/photo-library')
              })
          })

This test creates a new user, opens the map, and displays the tag filter page and selects all tags for viewing, then submits.

  opentagfilter.cy.js:
  
        describe('template spec', () => {
            it('runs', () => {
              cy.visit('localhost:4200/new-user')
            })

            it('checks you can open tag filter and apply all tags', () => {
                cy.visit('localhost:4200/new-user')

                cy.get('input').type('fakename')

                cy.contains('Confirm').click()

                cy.url().should('include', "/home")

                cy.contains('fakename')

                cy.contains('View Map/Post to OverHere').click()

                cy.url().should('include', "/map")

                cy.contains('Pin Color Key').click()

                cy.get('button[id="restyes"]').click()
                cy.get('button[id="hangyes"]').click()
                cy.get('button[id="studyyes"]').click()
                cy.get('button[id="socialyes"]').click()

                cy.get('button[id="tagsubmit"]').click()

            })


          })

This test demonstrates some real time functionality. It creates a new user, makes a new post, filters to only show that post type to demonstrate filter functionality, and then logs out.

   newexperience.cy.js:
    
        describe('template spec', () => {
            it('runs', () => {
              cy.visit('localhost:4200/new-user')
            })

            it('Reroutes on back button press', () => {
              cy.visit('localhost:4200/new-user')

              cy.contains('Back').click()

              cy.url().should('include', '/login')
            })

            it('makes new acc and posts then filters only study spot', () => {
                cy.visit('localhost:4200/new-user')

                cy.get('input').type('fakename')

                cy.contains('Confirm').click()

                cy.url().should('include', "/home")

                cy.contains('fakename')

                cy.contains('View Map/Post to OverHere').click()

                cy.url().should('include', "/map")

                cy.contains('New Post').click()

                cy.get('input[type=text]').type('cypress test post caption!')

                cy.get('select').select('Hangout Spot')

                cy.contains('Submit').click()

                cy.get('#map').click('center')

                cy.contains('Pin Color Key').click()

                cy.get('button[id="restno"]').click()
                cy.get('button[id="hangyes"]').click()
                cy.get('button[id="studyno"]').click()
                cy.get('button[id="socialno"]').click()

                cy.get('button[id="tagsubmit"]').click()

                cy.contains('Back to Homepage').click()

                cy.contains('Sign Out').click()
            })


          })

Old cypress tests that are no longer functional are due to the implementation of the backend which won't allow for users to just log in without creating an account and storing on the backend.
