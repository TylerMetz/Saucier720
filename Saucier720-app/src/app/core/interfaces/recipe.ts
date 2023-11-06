import { Ingredient } from "./ingredient";

export interface Recipe {
    instructions: string;
    ingredients: string[];
    title: string;
    // pictureLink: string | null; we should do this later but im commenting out for now
    recipeID: number;
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

