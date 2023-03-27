describe('Basic Cypress Test', () => {
  it('Reroutes on button press', () => {
    cy.visit('/')
    
    cy.contains('Sign Up').click()

    cy.url().should('include', '/new-user')
  })
})
