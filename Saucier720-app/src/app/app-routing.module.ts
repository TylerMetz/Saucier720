import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DealsComponent } from './deals/deals.component';
import { ListComponent } from './list/list.component';
import { PantryComponent } from './pantry/pantry.component';
import { RecipesComponent } from './recipes/recipes.component';

const routes: Routes = [
  { path: 'Pantry', component: PantryComponent},
  { path: 'List', component: ListComponent},
  { path: 'Deals', component: DealsComponent},
  { path: 'Recipes', component: RecipesComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
export const routingComponents = [PantryComponent, ListComponent, DealsComponent, RecipesComponent]