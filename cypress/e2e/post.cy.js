describe('template spec', () => {
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