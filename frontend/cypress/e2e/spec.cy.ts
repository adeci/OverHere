describe('Basic Cypress Test', () => {
  it('Visits the initial project page', () => {
    cy.visit('/')
    
    cy.contains('Sign Up').click()

    cy.url().should('include', '/new-user')
  })
})
