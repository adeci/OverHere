describe('template spec', () => {
    it('runs', () => {
      cy.visit('localhost:4200')
    })
  
    it('user doesnt exist', () => {
        cy.visit('localhost:4200/returning-user')
    
        cy.get('input').type('badusername')
    
        cy.contains('Confirm').click()
        
      })
  })