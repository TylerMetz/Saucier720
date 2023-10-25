export interface Ingredient {
    Name: string;
    FoodType: string;
    SaleDetails: string;
    Quantity: number;
}

export interface Pantry {
    Ingredients: Ingredient[]
}

export interface List {
    Ingredients: Ingredient[]
}

export interface Deals {
    Ingredients: Ingredient[]
}
