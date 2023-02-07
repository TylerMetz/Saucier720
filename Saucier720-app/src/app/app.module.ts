import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { PantryComponent } from './pantry/pantry.component';
import { ListComponent } from './list/list.component';

@NgModule({
  declarations: [
    AppComponent,
    PantryComponent,
    ListComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
