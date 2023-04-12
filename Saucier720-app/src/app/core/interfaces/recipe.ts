export interface Recipe {
    instructions: string;
    ingredients: string[];
    title: string;
    pictureLink: string;
}

export interface RecipePost {
    R: Recipe;
    ItemsInPantry: string[];
    ItemsOnSale: string[];
}
