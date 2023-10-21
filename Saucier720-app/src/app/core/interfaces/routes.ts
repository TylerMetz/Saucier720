interface Routes {
    Login: {
      path: 'localhost:8080/Login';
      method: 'GET';
    };
    Pantry: {
      path: 'localhost:8080/Pantry';
      method: 'GET';
    };
    Recipes: {
      path: 'localhost:8080/Recipes';
      method: 'GET';
    };
    FavoriteRecipes: {
      path: 'localhost:8080/Recipes/Favorite';
      method: 'GET';
    };
    Deals: {
      path: 'localhost:8080/Deals';
      method: 'GET';
    };
    DealsByStore: {
      path: 'localhost:8080/Deals/Store';
      method: 'GET';
    };
    List: {
      path: 'localhost:8080/List';
      method: 'GET';
    };
    Signup: {
      path: 'localhost:8080/Signup';
      method: 'PUT';
    };
    NewPantryItem: {
      path: 'localhost:8080/NewPantryItem';
      method: 'PUT';
    };
    NewRecipe: {
      path: 'localhost:8080/NewRecipe';
      method: 'PUT';
    };
    NewListIngredient: {
      path: 'localhost:8080/NewListIngredient';
      method: 'PUT';
    };
    NewFavoriteRecipe: {
      path: 'localhost:8080/NewFavoriteRecipe';
      method: 'PUT';
    };
    Logout: {
      path: 'localhost:8080/Logout';
      method: 'DELETE';
    };
    DeletePantryIngredient: {
      path: 'localhost:8080/DeletePantryIngredient';
      method: 'DELETE';
    };
    DeleteListIngredient: {
      path: 'localhost:8080/DeleteListIngredient';
      method: 'DELETE';
    };
    DeleteFavoriteRecipe: {
      path: 'localhost:8080/DeleteFavoriteRecipe';
      method: 'DELETE';
    };
    DeleteUserRecipe: {
      path: 'localhost:8080/DeleteUserRecipe';
      method: 'DELETE';
    };
    UpdatePantry: {
      path: 'localhost:8080/UpdatePantry';
      method: 'PUT';
    };
    UpdateList: {
      path: 'localhost:8080/UpdateList';
      method: 'PUT';
    };
    UpdateRecipe: {
      path: 'localhost:8080/UpdateRecipe';
      method: 'PUT';
    };
  }
  
  export default Routes;
  