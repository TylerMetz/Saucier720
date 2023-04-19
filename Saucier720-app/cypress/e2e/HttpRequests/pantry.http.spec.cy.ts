/// <reference types="cypress" />
import { Ingredient } from "src/app/core/interfaces/ingredient";

context('Network Requests', () => {
  let pantryPageUrl = 'http://localhost:4200/Pantry';
  let pantryGETUrl = 'http://localhost:8080/api/Pantry';
  let pantryPostUrl = 'http://localhost:8083/api/NewPantryItem';
  beforeEach(() => {
    // cy.visit(pantryPageUrl)
    cy.login();
    cy.wait(3000);
  })

  afterEach(() => {
    cy.wait(5000);
  })

  // Manage HTTP requests in your app

  it('make a GET request when loading the pantry', () => {
   
    cy.request(pantryGETUrl)
      .should((response) => {
        expect(response.status).to.eq(200)
        // the server sometimes gets an extra comment posted from another machine
        // which gets returned as 1 extra object
        const ingredients: Ingredient[] = response.body as Ingredient[];

       expect(ingredients).to.be.an('array').that.is.not.empty;
       expect(ingredients[0]).to.have.property('Name');
       expect(ingredients[0]).to.have.property('Quantity');
        
        expect(response).to.have.property('headers')
        expect(response).to.have.property('duration')
      })
  })

  it('posting new Pantry Item from pantry form', () => {
    
    cy.visit(pantryPageUrl);
    cy.get('app-new-pantry-item-button');
    const name = 'Pear';

    cy.get('#name').type(name, {delay: 150});
    cy.contains('Post').click();
    cy.request(pantryGETUrl)
    
    cy.get('app-pantry-table').contains('Pear');
  })

  it('deleting pantry items', () => {
    cy.visit(pantryPageUrl);
    let cookie = cy.getCookie('sessionID')
    cy.request({
      url: pantryGETUrl,
      headers: {
        'Content-Type': 'application/json',
        'Cookie': cookie
      },
    })
    const name = 'starfruit';

    cy.get('app-new-pantry-item-button');

    cy.get('#name').type(name, {delay: 150});
    cy.contains('Post').click();
    
    cy.request({
      url: pantryGETUrl,
      headers: {
        'Content-Type': 'application/json',
        'Cookie': cookie
      },
    })
    
    

    cy.get('tr')
    .contains('starfruit')
    .parent()
    .find('button')
    .contains('Delete')
    .click();
    cy.get('button')
    .contains('UpdatePantry')
    .click();

    
    cy.request({
      url: pantryGETUrl,
      headers: {
        'Content-Type': 'application/json',
        'Cookie': cookie
      },
    })
    cy.get('tr')
    .should('not.contain', 'starfruit')

  })

  it('increasing quantity pantry items', () => {
    cy.visit(pantryPageUrl);
    let cookie = cy.getCookie('sessionID')
    cy.request({
      url: pantryGETUrl,
      headers: {
        'Content-Type': 'application/json',
        'Cookie': cookie
      },
    })
    const name = 'starfruit';

    cy.get('app-new-pantry-item-button');

    cy.get('#name').type(name, {delay: 150});
    cy.contains('Post').click();
    
    cy.reload();
    cy.request({
      url: pantryGETUrl,
      headers: {
        'Content-Type': 'application/json',
        'Cookie': cookie
      },
    })
  
    cy.get('tr')
    .contains('starfruit')
    .parent()
    .find('button')
    .contains('+')
    .click();

    cy.get('button')
    .contains('UpdatePantry')
    .click();


    cy.visit(pantryPageUrl);
    cy.request({
      url: pantryGETUrl,
      headers: {
        'Content-Type': 'application/json',
        'Cookie': cookie
      },
    })
    cy.reload();
    cy.get('tr')
    .should('contain', 'starfruit', '2')

  })

})
