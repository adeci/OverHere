Frontend:
We had two main focuses for this sprint, and we were able to target them both. The most important was creating a post functionality, in which a user can upload an image from their device, add a caption, and add a tag and have it displayed on the map at the precise location they choose. In addition to adding this core funcionality, we added changes to the overall flow of the app. More visual effects and images were added to all pages and the app's visuals were made more consistent.
In addition, we created more extensive tests to ensure our routing and input functions work. We also tested the map and it's post functionality. Listed below are the new cypress tests we implemented. Note: due to the large amount of components used in our app, we opted to do our unit and end to end tests in e2e format. Video documentation of these tests can also be found in the cypress folder.

    detailpost.cy ------
    it('runs', () => {
          cy.visit('localhost:4200')
    })

        it('Full existing user post 2x', () => {
            cy.visit('localhost:4200/returning-user')

            cy.get('input').type('fakename')

            cy.contains('Confirm').click()

            cy.url().should('include', "/home")

            cy.contains('View Map/Post to OverHere').click()

            cy.url().should('include', "/map")

            cy.contains('New Post').click()

            cy.get('input[type=text]').type('cypress test post caption!')

            cy.get('select').select('Hangout Spot')

            cy.contains('Submit').click()

            cy.get('#map').click('center')

            cy.get('.leaflet-marker-icon.leaflet-interactive').click()

            cy.get('.leaflet-touch .leaflet-control-zoom-out').click()
            cy.get('.leaflet-touch .leaflet-control-zoom-out').click()

            cy.contains('New Post').click()

            cy.get('input[type=text]').type('new second post test!')

            cy.get('select').select('Study Spot')

            cy.contains('Submit').click()

            cy.get('#map').click(240, 460)

            cy.get('#map').click(240, 460)


          })
          
First test loads the page, the second test creates a thorough test in which a user signs in, creates two posts with captions, and posts them on seperate locations on the map

    homepage.cy.js------
     it('runs', () => {
        cy.visit('localhost:4200/home')
      })

      it('All button components properly route', () => {
        cy.visit('localhost:4200/home')

        cy.contains('View Map').click()

        cy.url().should('include', '/map')

        cy.contains('Back').click()

        cy.url().should('include', '/home')

        cy.contains('View Your').click()

        cy.url().should('include', '/photo-library')

        cy.contains('Back').click()

        cy.url().should('include', '/home')

        cy.contains('Sign Out').click()

        cy.url().should('include', '/login')
      })
    })
    
The first test loads the page, and the second test manually checks that each routing function works for the button component it is tied to.

    loginexistinguser---
    it('runs', () => {
        cy.visit('localhost:4200/returning-user')
      })

      it('Reroutes on back button press', () => {
        cy.visit('localhost:4200/returning-user')

        cy.contains('Back').click()

        cy.url().should('include', '/login')
      })

      it('Recieves inputs and moves on to the next page- e2e', () => {
        cy.visit('localhost:4200/returning-user')

        cy.get('input').type('fakename')

        cy.contains('Confirm').click()

        cy.url().should('include', "/home")

        cy.contains('fakename')
      })
    })
    
First test loads the page, the second test inputs a username, clicks the button to move on, and checks that the username was successfully saved and displayed between routes.

    loginpage.cy.js
    it('runs', () => {
        cy.visit('localhost:4200')
      })

      it('Reroutes on SignUp button press', () => {
        cy.visit('localhost:4200')

        cy.contains('Sign Up').click()

        cy.url().should('include', '/new-user')
      })

      it('Reroutes on LogIn button press', () => {
        cy.visit('localhost:4200')

        cy.contains('Log In').click()

        cy.url().should('include', '/returning-user')
      })
    })
First test loads the page, the second ensures the routing functions tied to both buttons execute correctly

    post.cy.js-----
    it('runs', () => {
          cy.visit('localhost:4200')
        })

        it('Log in, open map, create blank post pin', () => {
            cy.visit('localhost:4200/returning-user')

            cy.get('input').type('fakename')

            cy.contains('Confirm').click()

            cy.url().should('include', "/home")

            cy.contains('View Map/Post to OverHere').click()

            cy.url().should('include', "/map")

            cy.contains('New Post').click()

            cy.contains('Submit').click()

            cy.get('#map').click('center')

            cy.get('.leaflet-marker-icon.leaflet-interactive').click()

          })
      })
      
The first test loads the page, the second logs in with a test user and creates a basic post on the map with a default image and no caption or tag

    signup.cy.js----
    it('runs', () => {
        cy.visit('localhost:4200/new-user')
      })

      it('Reroutes on back button press', () => {
        cy.visit('localhost:4200/new-user')

        cy.contains('Back').click()

        cy.url().should('include', '/login')
      })

      it('Recieves inputs and moves on to the next page- e2e', () => {
        cy.visit('localhost:4200/new-user')

        cy.get('input').type('fakename')

        cy.contains('Confirm').click()

        cy.url().should('include', "/home")

        cy.contains('fakename')
      })
    })
    
First test loads the page, the second creates a new user with a test username and ensures that it is saved across routes and displays properly


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

    PutUser(userid string, username string):
        Updates username for User. userid is search key, username is new updated username to put.

    PutOHPost(object OHPostObject):
        Updates OHPost.

    PutImage(object ImageObject):
        Updates Image.

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


Alex- 
For this sprint, I focused on creating API controllers for images, ohposts, and users.
We started with Image and user CREATE and GET controllers, for a total of 4.
Now we have 12, a set of 4 POST, GET, PUT, and DELETE for each one. All can be stored in the database.
They function similarly, but there were many setbacks encountered when making them.

When the objects are created, an id is generated and the whole object is returned.
Unfortunately, there's not a great way of faking those ids right now, so unit tests for GET, PUT, and DELETE weren't feasible, often becoming 
monoliths of code.
I opted to focus on making the Postman tests

Added error checking 
Added ID code, turns out / and # for "Username-code" cause issues for Gin routing.
Matched frontend model so our team can pass all the data to the backend.

There's a lot more to do. 
Particularly, our team wanted to have the creation and deletion of Posts combined with the creation of Images.
Additionally, an API to return posts within a square of coordinates is important.
This will be my next goal, and will allow the frontend to display all images.