const MealDealzRoutes: Routes = {
  loginUrl: 'http://localhost:8080/Login',
  logoutUrl: 'http://localhost:8080/Logout',
  getPantryUrl: 'http://localhost:8080/Pantry',
  postPantryUrl: 'http://localhost:8080/NewPantryItem',
  updatePantryUrl:  'http://localhost:8080/UpdatePantry',
  deletePantryUrl: 'http://localhost:8080/DeletePantryIngredient',
  signupUrl:  'http://localhost:8080/Signup',
  // Add more routes as needed
};


// routes.interface.ts
export interface Routes {
    loginUrl: string;
    logoutUrl: string;
    getPantryUrl: string;
    updatePantryUrl: string; 
    postPantryUrl: string;
    deletePantryUrl: string;
    signupUrl: string;
    // Add more routes as needed
}
  
  
export default MealDealzRoutes;