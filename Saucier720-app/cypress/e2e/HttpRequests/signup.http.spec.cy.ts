/// <reference types="cypress" />
import faker from 'faker';

context('Pantry Requests', () => {
  let signupApiUrl = 'http://localhost:8081/api/Signup';
  let signupPageUrl = 'http://localhost:4200/Signup';

  beforeEach(() => {

  })

  afterEach(() => {
    cy.wait(5000);
  })

  it('signup new users', () => {
    cy.visit(signupPageUrl);

    const firstName = faker.name.firstName();
    const lastName = faker.name.lastName();
    const email = faker.internet.email();
    const userName = faker.internet.userName();
    const password = faker.internet.password();

    cy.get('#firstName').type(firstName);
    cy.get('#lastName').type(lastName);
    cy.get('#email').type(email);
    cy.get('#username').type(userName);
    cy.get('#password').type(password);

    cy.get('.signup-button button').click();

    cy.request({
      method: 'POST',
      url: signupApiUrl,
      headers: {
        'Content-Type': 'application/json',
      },
      body: { 
        user: {
          FirstName: firstName,
          LastName: lastName,
          Email: email,
          UserName: userName,
          Password: password
        }
      }
    });

    // cy.visit('http://localhost:4200/Login');
    // cy.get('#username').should('be.visible').clear().type(userName, {delay: 150});
    // cy.get('#password').should('be.visible').clear().type(password, {delay: 150});
    
    // cy.setCookie('sessionID', 'ri720');
    // cy.get('button[type="submit"]').click();
  })
})
