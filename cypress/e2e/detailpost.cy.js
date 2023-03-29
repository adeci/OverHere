describe('template spec', () => {
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
  })