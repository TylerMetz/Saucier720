import { Ingredient } from "./ingredient";

export interface Recipe {
    instructions: string;
    ingredients: string[];
    title: string;
    pictureLink: string | null;
    recipeID: string;
    [key: string]: any; // index signature
}

export interface RecipePost {
    R: Recipe;
    ItemsInPantry: Ingredient[];
    ItemsOnSale: Ingredient[];
}
