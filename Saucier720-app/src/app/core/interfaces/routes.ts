const MealDealzRoutes: Routes = {
  loginUrl: 'http://localhost:8080/Login',
  logoutUrl: 'http://localhost:8080/Logout',
  getPantryUrl: 'http://localhost:8080/Pantry',
  // Add more routes as needed
};


// routes.interface.ts
export interface Routes {
    loginUrl: string;
    logoutUrl: string;
    getPantryUrl: string;
    // Add more routes as needed
}
  
  
export default MealDealzRoutes;