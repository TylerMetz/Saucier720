/// <reference types="cypress" />

context('Deals Requests', () => {
  let dealsPageUrl: string = 'http://localhost:4200/Deals';
  let dealsGETUrl: string = 'http://localhost:8081/api/Deals';

  // Manage HTTP requests in your app
  beforeEach(() => {
    cy.login();
    cy.wait(3000);
    cy.visit(dealsPageUrl)
  })

  it('cy.request() - make a GET request when loading the page', () => {
    cy.request(dealsGETUrl)
      .should((response) => {
        expect(response.status).to.eq(200)
        // the server sometimes gets an extra comment posted from another machine
        // which gets returned as 1 extra object
        expect(response.body).to.have.property('length')
        expect(response).to.have.property('headers')
        expect(response).to.have.property('duration')
      })
  })
});
