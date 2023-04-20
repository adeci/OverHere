describe('template spec', () => {
    it('runs', () => {
      cy.visit('localhost:4200/new-user')
    })
  
    it('checks you can open tag filter and apply all tags', () => {
        cy.visit('localhost:4200/new-user')

        cy.get('input').type('fakename')

        cy.contains('Confirm').click()

        cy.url().should('include', "/home")

        cy.contains('fakename')

        cy.contains('View Map/Post to OverHere').click()

        cy.url().should('include', "/map")
    
        cy.contains('Pin Color Key').click()

        cy.get('button[id="restyes"]').click()
        cy.get('button[id="hangyes"]').click()
        cy.get('button[id="studyyes"]').click()
        cy.get('button[id="socialyes"]').click()

        cy.get('button[id="tagsubmit"]').click()

    })
  

  })