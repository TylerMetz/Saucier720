import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { AppRoutingModule, routingComponents } from './app-routing.module';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { DealsTableComponent } from './deals/FEC/deals-table/deals-table.component';
import { DealsStoreButtonComponent } from './deals/FEC/deals-store-button/deals-store-button.component';
import { PantryTableComponent } from './pantry/FEC/pantry-table/pantry-table.component';
import { NewPantryItemButtonComponent } from './pantry/FEC/new-pantry-item-button/new-pantry-item-button.component';
import { SignupFormComponent } from './signup/FEC/signup-form/signup-form.component';
import { FormsModule } from '@angular/forms';
import { RecipeCardComponent } from './recipes/FEC/recipe-card/recipe-card.component';
import { LoginFormComponent } from './login/FEC/login-form/login-form.component';
import { SubRecipeComponent } from './recipes/FEC/sub-recipe/sub-recipe.component';
import { RecipeNavBarComponent } from './recipes/FEC/recipe-nav-bar/recipe-nav-bar.component';
import { RecipesComponent } from './recipes/recipes.component'; 
import { NewRecipeComponent } from './recipes/FEC/new-recipe/new-recipe.component';
import { ListComponent } from './list/list.component';
import { FilterRecipeComponent } from './recipes/FEC/filter-menu/filter-menu.component';
import { HomePageComponent } from './homepage/homepage.component';
import { LandingPageComponent } from './homepage/FEC/landing-page/landing-page.component';
import { UserDashboardComponent } from './homepage/FEC/user-dashboard/user-dashboard.component';
import { AnimatedLogoComponent } from './homepage/FEC/landing-page/FEC/animated-logo/animated-logo.component';
import { GeneralButtonsComponent } from './homepage/FEC/landing-page/FEC/general-buttons/general-buttons.component';

@NgModule({
  declarations: [
    AppComponent,
    routingComponents,
    LoginComponent,
    SignupComponent,
    DealsTableComponent,
    DealsStoreButtonComponent,
    PantryTableComponent,
    NewPantryItemButtonComponent,
    SignupFormComponent,
    RecipeCardComponent,
    LoginFormComponent,
    SubRecipeComponent,
    RecipesComponent,
    RecipeNavBarComponent,
    NewRecipeComponent,
    ListComponent,
    FilterRecipeComponent,
    HomePageComponent,
    LandingPageComponent,
    UserDashboardComponent,
    AnimatedLogoComponent,
    GeneralButtonsComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    HttpClientModule,
    FormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
