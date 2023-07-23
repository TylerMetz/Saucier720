import { StringNullableChain } from "cypress/types/lodash";
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

export interface RecipePost {
    R: Recipe;
    ItemsInPantry: Ingredient[];
    ItemsOnSale: Ingredient[];
}
