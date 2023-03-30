describe('template spec', () => {
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