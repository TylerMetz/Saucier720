export interface Recipe {
    instructions: string;
    ingredients: string[];
    title: string;
    pictureLink: string;
    [key: string]: any; // index signature
}

export interface RecipePost {
    R: Recipe;
    ItemsInPantry: string[];
    ItemsOnSale: string[];
}
