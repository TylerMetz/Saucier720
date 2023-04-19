/// <reference types="cypress" />

context('Pantry Requests', () => {
  let recipeApiUrl = 'http://localhost:8082/api/Recipes';
  let recipePageUrl = 'http://localhost:4200/Recipes';

  beforeEach(() => {
    cy.login();
    cy.wait(3000);
    cy.visit(recipePageUrl);
  })

  afterEach(() => {
    cy.wait(5000);
  })

  it('signup new users', () => {
    
    
  })
})
