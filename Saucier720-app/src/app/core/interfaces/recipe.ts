import { Ingredient } from "./ingredient";

export interface Recipe {
    instructions: string;
    ingredients: string[];
    title: string;
    pictureLink: string | null;
    recipeID: string;
    userFavorite: boolean;
    recipeAuthor: string;
    [key: string]: any; // index signature
}


export interface Recommendation {
    R: Recipe;
    ItemsInPantry: Ingredient[];
    ItemsOnSale: Ingredient[];
}

export interface RecommendedRecipes {
    Recommendations: Recommendation[];
}

