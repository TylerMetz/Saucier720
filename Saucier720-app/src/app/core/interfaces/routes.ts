const MealDealzRoutes: Routes = {
  loginUrl: 'http://localhost:8080/Login',
  logoutUrl: 'http://localhost:8080/Logout',
  getPantryUrl: 'http://localhost:8080/Pantry',
  updatePantryUrl:  'http://localhost:8080/UpdatePantry',
  signupUrl:  'http://localhost:8080/Signup',
  getListUrl: 'http://localhost:8080/List',
  postListUrl: 'http://localhost:8080/NewListIngredient',
  // Add more routes as needed
};


// routes.interface.ts
export interface Routes {
    loginUrl: string;
    logoutUrl: string;
    getPantryUrl: string;
    updatePantryUrl: string; 
    signupUrl: string;
    getListUrl: string; 
    postListUrl: string; 
    // Add more routes as needed
}
  
  
export default MealDealzRoutes;