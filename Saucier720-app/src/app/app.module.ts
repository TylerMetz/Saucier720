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
import { ListComponent } from './list/list.component';

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
    ListComponent
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
