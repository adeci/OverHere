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