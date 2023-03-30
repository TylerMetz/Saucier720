import { Ingredient } from "../core/interfaces/ingredient";

    var mockIngredient1: Ingredient = {
        "Name":"riley butter",
        "StoreCost":369.99,
        "OnSale":true,
        "SalePrice":0,
        "SaleDetails":"BOGO",
        "Quantity":10
    };

    var mockIngredient2: Ingredient = {
        "Name":"peanut butter",
        "StoreCost":369.99,
        "OnSale":true,
        "SalePrice":0,
        "SaleDetails":"BOGO",
        "Quantity":10
    };

    var mockIngredient3: Ingredient = {
        "Name":"jelly",
        "StoreCost":1,
        "OnSale":false,
        "SalePrice":0,
        "SaleDetails":"N/A",
        "Quantity":30
    };
    
    var mockIngredient4: Ingredient = {
        "Name":"bread",
        "StoreCost":10.69,
        "OnSale":true,
        "SalePrice":0,
        "SaleDetails":"$2 for 2",
        "Quantity":2
    };

export const PANTRY: Ingredient[] = [
    mockIngredient1,
    mockIngredient2,
    mockIngredient3,
    mockIngredient4,
];
