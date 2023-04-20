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