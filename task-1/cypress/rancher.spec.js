// cypress/rancher.spec.js

describe('Rancher UI Tests', () => {
  it('logs in to Rancher', () => {
    cy.visit('https://rancher.example.com')
    cy.get('input[name="username"]').type('myusername')
    cy.get('input[name="password"]').type('mypassword')
    cy.get('button[type="submit"]').click()
    cy.url().should('include', '/dashboard')
  })

  it('checks if the main page opens up', () => {
    cy.visit('https://rancher.example.com')
    cy.get('input[name="username"]').type('myusername')
    cy.get('input[name="password"]').type('mypassword')
    cy.get('button[type="submit"]').click()
    cy.get('.navbar-brand').should('be.visible')
    cy.get('.nav-link').contains('Clusters').should('be.visible')
    cy.get('.nav-link').contains('Projects').should('be.visible')
    cy.get('.nav-link').contains('Catalog').should('be.visible')
  })

  it('checks if the main page title is correct', () => {
    cy.visit('https://rancher.example.com')
    cy.get('input[name="username"]').type('myusername')
    cy.get('input[name="password"]').type('mypassword')
    cy.get('button[type="submit"]').click()
    cy.title().should('eq', 'Rancher - Dashboard')
  })
})

