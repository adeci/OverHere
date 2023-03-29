describe('template spec', () => {
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