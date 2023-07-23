import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DealsComponent } from './deals/deals.component';
import { ListComponent } from './list/list.component';
import { HomePageComponent } from './homepage/homepage.component';
import { PantryComponent } from './pantry/pantry.component';
import { RecipesComponent } from './recipes/recipes.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { HttpClientModule } from '@angular/common/http';
import { AuthGuard } from './core/services/Auth/auth.guard.service';
import { AuthService } from './core/services/Auth/auth.service';

const routes: Routes = [
  { path: 'Home', component: HomePageComponent},
  { path: 'Pantry', component: PantryComponent, canActivate: [AuthGuard]},
  { path: 'List', component: ListComponent, canActivate: [AuthGuard]},
  { path: 'Deals', component: DealsComponent, canActivate: [AuthGuard]},
  { path: 'Recipes', component: RecipesComponent, canActivate: [AuthGuard]},
  { path: 'Login', component: LoginComponent},
  { path: 'Signup', component: SignupComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes), HttpClientModule],
  exports: [RouterModule],
  providers: [AuthGuard, AuthService]
})
export class AppRoutingModule { }
export const routingComponents = [PantryComponent, ListComponent, DealsComponent, RecipesComponent, LoginComponent, SignupComponent, HomePageComponent];
