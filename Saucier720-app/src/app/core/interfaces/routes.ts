const MealDealzRoutes: Routes = {
  //user stuff
  loginUrl: 'http://localhost:8080/Login',
  logoutUrl: 'http://localhost:8080/Logout',
  signupUrl:  'http://localhost:8080/Signup',
  //pantry
  getPantryUrl: 'http://localhost:8080/Pantry',
  postPantryUrl: 'http://localhost:8080/NewPantryItem',
  updatePantryUrl:  'http://localhost:8080/UpdatePantry',
  deletePantryUrl: 'http://localhost:8080/DeletePantryIngredient',
  //list
  getListUrl: 'http://localhost:8080/List',
  postListUrl: 'http://localhost:8080/NewListIngredient',
  updateListUrl: 'http://localhost:8080/UpdateList',
  //recipes
  getRecipesUrl: 'http://localhost:8080/Recipes'
};


// routes.interface.ts
export interface Routes {
  // user stuff
    loginUrl: string;
    logoutUrl: string;
    signupUrl: string;
  // pantry
    getPantryUrl: string;
    updatePantryUrl: string; 
    postPantryUrl: string;
    deletePantryUrl: string;
  // list
    getListUrl: string; 
    postListUrl: string; 
    updateListUrl: string;
  // recipes
    getRecipesUrl: string;
}
  
  
export default MealDealzRoutes;