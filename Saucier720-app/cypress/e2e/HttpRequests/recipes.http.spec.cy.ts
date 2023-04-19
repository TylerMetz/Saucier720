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

  it('receiving recipes items', () => {
    cy.visit(recipePageUrl)
    cy.request(recipeApiUrl).then((response) => {
      expect(response.status).to.eq(200); // Verify that the HTTP request was successful
      expect(response.body).to.have.length.above(0); // Verify that the response data contains at least one recipe
    });
    cy.visit(recipePageUrl)
    cy.wait(5000); // Wait for 5 seconds for the recipe card to load
    cy.get('app-recipe-card').should('be.visible');
    
  })
})
