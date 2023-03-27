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