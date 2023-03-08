import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DealsComponent } from './deals/deals.component';
import { ListComponent } from './list/list.component';
import { PantryComponent } from './pantry/pantry.component';
import { RecipesComponent } from './recipes/recipes.component';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';

const routes: Routes = [
  { path: 'Pantry', component: PantryComponent},
  { path: 'List', component: ListComponent},
  { path: 'Deals', component: DealsComponent},
  { path: 'Recipes', component: RecipesComponent},
  { path: 'login', component: LoginComponent},
  { path: 'signup', component: SignupComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
export const routingComponents = [AppComponent, PantryComponent, ListComponent, DealsComponent, RecipesComponent]