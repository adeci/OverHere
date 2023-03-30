describe('template spec', () => {
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